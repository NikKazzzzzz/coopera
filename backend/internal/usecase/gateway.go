package usecase

import (
	"context"

	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
	"github.com/jackc/pgx/v4"
)

// Здесь будут заданы интерфейсы для взаимодействия с имплементациями сдоя репозитория и управления транзакциями

type TransactionManageRepository interface {
	WithinTransaction(ctx context.Context, fn func(txCtx context.Context) error) error
	WithinTransactionWithIsolation(ctx context.Context, level pgx.TxIsoLevel, fn func(txCtx context.Context) error) error
}

type UserRepository interface {
	CreateRepo(ctx context.Context, euser entity.UserEntity) (entity.UserEntity, error)
	DeleteRepo(ctx context.Context, userID int32) error
	GetRepo(ctx context.Context, opts ...any) (entity.UserEntity, error)
	UpdateSettingsRepo(ctx context.Context, userID int32, wallpaper, customURL, theme string) error
}

type TeamRepository interface {
	CreateRepo(ctx context.Context, team entity.TeamEntity) (entity.TeamEntity, error)
	DeleteRepo(ctx context.Context, teamID int32) error
	GetByIDRepo(ctx context.Context, teamID int32) (entity.TeamEntity, error)
	ExistsByName(ctx context.Context, name string) (bool, error)
	ExistsByID(ctx context.Context, teamID int32) (bool, error)
	UpdateMetaRepo(ctx context.Context, teamID int32, emoji, color string) error
	UpdateAutoassignRepo(ctx context.Context, teamID int32, autoassign bool) error
	UpdatePhotoURLRepo(ctx context.Context, teamID int32, photoURL string) error
}

type MembershipRepository interface {
	AddMemberRepo(ctx context.Context, membership entity.MembershipEntity) (int32, error)
	DeleteMemberRepo(ctx context.Context, memberID int32) error
	GetMembersRepo(ctx context.Context, teamID int32) ([]entity.MembershipEntity, error)
	MemberExistsRepo(ctx context.Context, memberID int32) (bool, error)
	GetMemberRepo(ctx context.Context, teamID, memberID int32) (entity.MembershipEntity, error)
}

type TaskRepository interface {
	CreateRepo(ctx context.Context, task entity.Task) (entity.Task, error)
	GetByTaskID(ctx context.Context, id int32) (entity.Task, error)
	GetByAssignedTo(ctx context.Context, memberID int32) ([]entity.Task, error)
	GetByTeamID(ctx context.Context, teamID int32) ([]entity.Task, error)
	UpdateStatus(ctx context.Context, status entity.TaskStatus) error
	DeleteRepo(ctx context.Context, taskID int32) error
	UpdateRepo(ctx context.Context, task entity.UpdateTask) error
	GetAllTasks(ctx context.Context) ([]entity.Task, error)
}

type TaskCommentRepository interface {
	GetByTaskIDRepo(ctx context.Context, taskID int32) ([]entity.TaskComment, error)
	CreateRepo(ctx context.Context, c entity.TaskComment) (entity.TaskComment, error)
	DeleteRepo(ctx context.Context, id int32, userID int32) error
}

type ActivityRepository interface {
	CreateRepo(ctx context.Context, a entity.ActivityEntry) (entity.ActivityEntry, error)
	GetByUserIDRepo(ctx context.Context, userID int32) ([]entity.ActivityEntry, error)
	MarkAllReadRepo(ctx context.Context, userID int32) error
	MarkSingleReadRepo(ctx context.Context, id int32, userID int32) error
	DeleteAllRepo(ctx context.Context, userID int32) error
}
