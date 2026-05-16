package task_comment_repo

import (
	"context"

	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/model/task_comment_model"
	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/postgres/dao"
	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
)

type TaskCommentRepository struct {
	dao *dao.TaskCommentDAO
}

func NewTaskCommentRepository(d *dao.TaskCommentDAO) *TaskCommentRepository {
	return &TaskCommentRepository{dao: d}
}

func (r *TaskCommentRepository) GetByTaskIDRepo(ctx context.Context, taskID int32) ([]entity.TaskComment, error) {
	return r.dao.GetByTaskID(ctx, taskID)
}

func (r *TaskCommentRepository) CreateRepo(ctx context.Context, c entity.TaskComment) (entity.TaskComment, error) {
	return r.dao.Create(ctx, task_comment_model.TaskComment{
		TaskID:   c.TaskID,
		UserID:   c.UserID,
		Username: c.Username,
		Text:     c.Text,
	})
}

func (r *TaskCommentRepository) DeleteRepo(ctx context.Context, id int32, userID int32) error {
	return r.dao.Delete(ctx, id, userID)
}
