package task_comment

import (
	"context"

	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
	"github.com/NikKazzzzzz/coopera-backend/internal/usecase"
)

type TaskCommentUsecase struct {
	txManager usecase.TransactionManageRepository
	repo      usecase.TaskCommentRepository
}

func NewTaskCommentUsecase(repo usecase.TaskCommentRepository, txManager usecase.TransactionManageRepository) *TaskCommentUsecase {
	return &TaskCommentUsecase{txManager: txManager, repo: repo}
}

func (uc *TaskCommentUsecase) GetByTaskIDUsecase(ctx context.Context, taskID int32) ([]entity.TaskComment, error) {
	return uc.repo.GetByTaskIDRepo(ctx, taskID)
}

func (uc *TaskCommentUsecase) CreateUsecase(ctx context.Context, c entity.TaskComment) (entity.TaskComment, error) {
	var created entity.TaskComment
	err := uc.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		var err error
		created, err = uc.repo.CreateRepo(txCtx, c)
		return err
	})
	return created, err
}

func (uc *TaskCommentUsecase) DeleteUsecase(ctx context.Context, id int32, userID int32) error {
	return uc.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		return uc.repo.DeleteRepo(txCtx, id, userID)
	})
}
