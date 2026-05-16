package usecase

import (
	"context"
	"time"

	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
)

type UserUseCase interface {
	CreateUsecase(ctx context.Context, euser entity.UserEntity) (entity.UserEntity, error)
	GetUsecase(ctx context.Context, opts ...any) (entity.UserEntity, error)
	DeleteUsecase(ctx context.Context, userID int32) error
	UpdateSettingsUsecase(ctx context.Context, userID int32, wallpaper, customURL, theme string) error
}

type TeamUseCase interface {
	CreateUsecase(ctx context.Context, team entity.TeamEntity) (entity.TeamEntity, error)
	DeleteUsecase(ctx context.Context, teamID, currentUserID int32) error
	GetByIDUsecase(ctx context.Context, teamID int32) (entity.TeamEntity, []entity.MembershipEntity, map[int32]string, error)
	ExistTeamByIDUsecase(ctx context.Context, teamID int32) (bool, error)
	UpdateMetaUsecase(ctx context.Context, teamID, currentUserID int32, emoji, color string) error
	UpdateAutoassignUsecase(ctx context.Context, teamID, currentUserID int32, autoassign bool) error
	UpdatePhotoURLUsecase(ctx context.Context, teamID, currentUserID int32, photoURL string) error
}

type MembershipUseCase interface {
	AddMemberUsecase(ctx context.Context, membership entity.MembershipEntity) (int32, error)
	DeleteMemberUsecase(ctx context.Context, memberID, teamID, currentUserID int32) error
	GetMembersUsecase(ctx context.Context, teamID int32) ([]entity.MembershipEntity, error)
	ExistsMemberUsecase(ctx context.Context, memberID int32) (bool, error)
	GetMemberUsecase(ctx context.Context, teamID, memberID int32) (entity.MembershipEntity, error)
}

type TaskUseCase interface {
	CreateUsecase(ctx context.Context, task entity.Task) (entity.Task, error)
	GetUsecase(ctx context.Context, taskFilter entity.TaskFilter) ([]entity.Task, error)
	UpdateStatus(ctx context.Context, status entity.TaskStatus) error
	DeleteUsecase(ctx context.Context, taskID, currentUserID int32) error
	UpdateUsecase(ctx context.Context, task entity.UpdateTask, currentUserID int32) error
	UpdateStatusForEngine(ctx context.Context, taskStatus entity.TaskStatus) error
	GetAllTasks(ctx context.Context) ([]entity.Task, error)
	UpdateForEngine(ctx context.Context, task entity.UpdateTask) error
}

type TaskAssignmentUsecase interface {
	AssignTasks(ctx context.Context, taskMinAge time.Duration) error
}

type TaskCommentUseCase interface {
	GetByTaskIDUsecase(ctx context.Context, taskID int32) ([]entity.TaskComment, error)
	CreateUsecase(ctx context.Context, c entity.TaskComment) (entity.TaskComment, error)
	DeleteUsecase(ctx context.Context, id int32, userID int32) error
}

type ActivityUseCase interface {
	CreateUsecase(ctx context.Context, a entity.ActivityEntry) (entity.ActivityEntry, error)
	GetByUserIDUsecase(ctx context.Context, userID int32) ([]entity.ActivityEntry, error)
	MarkAllReadUsecase(ctx context.Context, userID int32) error
	MarkSingleReadUsecase(ctx context.Context, id int32, userID int32) error
	DeleteAllUsecase(ctx context.Context, userID int32) error
}
