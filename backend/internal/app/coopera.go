package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/NikKazzzzzz/coopera-backend/config"
	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/controller/task_controller"
	"github.com/NikKazzzzzz/coopera-backend/internal/usecase/task"
	"github.com/NikKazzzzzz/coopera-backend/internal/usecase/taskassigner"
	"github.com/NikKazzzzzz/coopera-backend/internal/usecase/user"
	"github.com/NikKazzzzzz/coopera-backend/pkg/tgphoto"

	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/controller/web_api"
	repoactivity "github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/activity_repo"
	repomembership "github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/membership_repo"
	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/postgres"
	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/postgres/dao"
	repotask "github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/task_repo"
	repocomment "github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/task_comment_repo"
	repoteams "github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/team_repo"
	repouser "github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/user_repo"
	usecaseactivity "github.com/NikKazzzzzz/coopera-backend/internal/usecase/activity"
	usecasecomment "github.com/NikKazzzzzz/coopera-backend/internal/usecase/task_comment"
	"github.com/NikKazzzzzz/coopera-backend/internal/usecase/memberships"
	"github.com/NikKazzzzzz/coopera-backend/internal/usecase/team"
	"github.com/NikKazzzzzz/coopera-backend/pkg/logger"
	"github.com/NikKazzzzzz/coopera-backend/pkg/migrator"
	"github.com/go-playground/validator/v10"
)

func Start() error {
	localSetupEnvPath := "config/dev/.env"
	cfg := config.LoadConfig(localSetupEnvPath)

	connectionString := postgres.BuildPath(cfg)

	if err := migrator.Migrate(cfg.MigrationsPath, connectionString, cfg.DBSchema); err != nil {
		return fmt.Errorf("migration error: %w", err)
	}

	db, err := postgres.NewDB(connectionString)
	if err != nil {
		return err
	}

	validate := validator.New()
	web_api.InitValidator(validate)

	logService := logger.NewLogger(cfg.LogLevel)

	userRepo := repouser.NewUserRepository(*dao.NewUserDAO(db))
	teamRepo := repoteams.NewTeamRepository(*dao.NewTeamDAO(db))
	taskRepo := repotask.NewTaskRepository(*dao.NewTaskDAO(db))
	memberRepo := repomembership.NewMembershipRepository(*dao.NewMembershipDAO(db))
	activityRepo := repoactivity.NewActivityRepository(dao.NewActivityDAO(db))
	commentRepo  := repocomment.NewTaskCommentRepository(dao.NewTaskCommentDAO(db))

	photoFetcher := tgphoto.New(cfg.TelegramBotToken)
	userUC := user.NewUserUsecase(userRepo, db, photoFetcher)
	memberUC := memberships.NewMembershipsUsecase(memberRepo, db)
	teamUC := team.NewTeamUsecase(teamRepo, memberUC, userUC, db)
	taskUC := task.NewTaskUsecase(taskRepo, memberUC, db, teamUC)
	activityUC := usecaseactivity.NewActivityUsecase(activityRepo, db)
	commentUC  := usecasecomment.NewTaskCommentUsecase(commentRepo, db)

	taskCtx, taskCancel := context.WithCancel(context.Background())
	defer taskCancel()

	taskAssignerUsecase := taskassigner.NewTaskAssignmentUsecase(db, taskUC, memberUC)
	taskAssigner := task_controller.NewTaskAssignmentController(taskAssignerUsecase)
	go taskAssigner.StartAssignmentLoop(taskCtx, cfg.AssignmentsWorkerInterval, cfg.TaskMinAge)

	router := web_api.NewRouter(userUC, teamUC, taskUC, memberUC, activityUC, commentUC, logService, cfg).SetupRoutes()

	srv := &http.Server{
		Addr:        ":" + cfg.BackendPort,
		Handler:     router,
		IdleTimeout: 60 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()
	log.Println("HTTP server started on port", cfg.BackendPort)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down...")

	taskCancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	return srv.Shutdown(shutdownCtx)
}
