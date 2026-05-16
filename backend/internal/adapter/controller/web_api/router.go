package web_api

import (
	"github.com/NikKazzzzzz/coopera-backend/config"
	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/controller/web_api/middleware"
	"github.com/NikKazzzzzz/coopera-backend/pkg/logger"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"time"

	"github.com/NikKazzzzzz/coopera-backend/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type Router struct {
	userController        *UserController
	teamController        *TeamController
	taskController        *TaskController
	membershipController  *MembershipController
	activityController    *ActivityController
	taskCommentController *TaskCommentController
	logger                *logger.Logger
	config                *config.Config
}

func NewRouter(
	userUseCase usecase.UserUseCase,
	teamUseCase usecase.TeamUseCase,
	taskUseCase usecase.TaskUseCase,
	membershipUseCase usecase.MembershipUseCase,
	activityUseCase usecase.ActivityUseCase,
	taskCommentUseCase usecase.TaskCommentUseCase,
	logger *logger.Logger,
	config *config.Config,
) *Router {
	return &Router{
		userController:        NewUserController(userUseCase),
		teamController:        NewTeamController(teamUseCase),
		taskController:        NewTaskController(taskUseCase),
		membershipController:  NewMembershipController(membershipUseCase),
		activityController:    NewActivityController(activityUseCase),
		taskCommentController: NewTaskCommentController(taskCommentUseCase),
		logger:                logger,
		config:                config,
	}
}

func (r *Router) SetupRoutes() http.Handler {
	router := chi.NewRouter()

	corsMiddleware := cors.Handler(cors.Options{
		AllowedOrigins:   []string{r.config.FrontendURL},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "ngrok-skip-browser-warning"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	router.Use(corsMiddleware)

	// Photo proxies без 5с таймаута — external CDN может отвечать дольше
	router.With(chimw.Recoverer).Get("/api/v1/users/{id}/photo", middleware.ErrorHandler(r.userController.GetPhoto))
	router.With(chimw.Recoverer).Get("/api/v1/teams/{id}/photo", middleware.ErrorHandler(r.teamController.GetPhoto))

	router.Group(func(api chi.Router) {
		api.Use(middleware.TimeoutMiddleware(5*time.Second), chimw.Recoverer)

		api.Route("/api/v1", func(api chi.Router) {
			api.Route("/users", func(users chi.Router) {
				users.Post("/", middleware.ErrorHandler(r.userController.Create))
				users.Get("/", middleware.ErrorHandler(r.userController.Get))
				users.Patch("/settings", middleware.ErrorHandler(r.userController.UpdateSettings))
			})

			api.Route("/teams", func(teams chi.Router) {
				teams.Post("/", middleware.ErrorHandler(r.teamController.Create))
				teams.Get("/", middleware.ErrorHandler(r.teamController.Get))
				teams.Delete("/", middleware.ErrorHandler(r.teamController.Delete))
				teams.Patch("/", middleware.ErrorHandler(r.teamController.UpdateMeta))
				teams.Patch("/autoassign", middleware.ErrorHandler(r.teamController.UpdateAutoassign))
				teams.Patch("/photo", middleware.ErrorHandler(r.teamController.UpdatePhoto))
			})

			api.Route("/memberships", func(members chi.Router) {
				members.Post("/", middleware.ErrorHandler(r.membershipController.AddMember))
				members.Delete("/", middleware.ErrorHandler(r.membershipController.DeleteMember))
			})

			api.Route("/tasks", func(tasks chi.Router) {
				tasks.Post("/", middleware.ErrorHandler(r.taskController.Create))
				tasks.Get("/", middleware.ErrorHandler(r.taskController.Get))
				tasks.Patch("/status", middleware.ErrorHandler(r.taskController.UpdateStatus))
				tasks.Delete("/", middleware.ErrorHandler(r.taskController.Delete))
				tasks.Patch("/", middleware.ErrorHandler(r.taskController.Update))

				tasks.Route("/{taskId}/comments", func(c chi.Router) {
					c.Get("/", middleware.ErrorHandler(r.taskCommentController.GetComments))
					c.Post("/", middleware.ErrorHandler(r.taskCommentController.CreateComment))
					c.Delete("/{commentId}", middleware.ErrorHandler(r.taskCommentController.DeleteComment))
				})
			})

			api.Route("/activity", func(act chi.Router) {
				act.Post("/", middleware.ErrorHandler(r.activityController.Create))
				act.Get("/", middleware.ErrorHandler(r.activityController.GetByUser))
				act.Patch("/read", middleware.ErrorHandler(r.activityController.MarkAllRead))
				act.Patch("/read/{id}", middleware.ErrorHandler(r.activityController.MarkSingleRead))
				act.Delete("/", middleware.ErrorHandler(r.activityController.DeleteAll))
			})
		}) // /api/v1
	}) // group with timeout

	return router
}
