package dao

import (
	"context"
	"fmt"

	repoErr "github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/errors"
	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/model/task_comment_model"
	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/postgres"
	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
)

type TaskCommentDAO struct {
	db *postgres.DB
}

func NewTaskCommentDAO(db *postgres.DB) *TaskCommentDAO {
	return &TaskCommentDAO{db: db}
}

func (d *TaskCommentDAO) GetByTaskID(ctx context.Context, taskID int32) ([]entity.TaskComment, error) {
	const query = `
		SELECT id, task_id, user_id, username, text, created_at
		FROM coopera.task_comments
		WHERE task_id = $1
		ORDER BY created_at ASC
	`
	rows, err := d.db.Pool.Query(ctx, query, taskID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", repoErr.ErrFailGet, err)
	}
	defer rows.Close()

	var result []entity.TaskComment
	for rows.Next() {
		var m task_comment_model.TaskComment
		if err := rows.Scan(&m.ID, &m.TaskID, &m.UserID, &m.Username, &m.Text, &m.CreatedAt); err != nil {
			return nil, fmt.Errorf("%w: %v", repoErr.ErrFailGet, err)
		}
		result = append(result, entity.TaskComment{
			ID: m.ID, TaskID: m.TaskID, UserID: m.UserID,
			Username: m.Username, Text: m.Text, CreatedAt: m.CreatedAt,
		})
	}
	return result, nil
}

func (d *TaskCommentDAO) Create(ctx context.Context, c task_comment_model.TaskComment) (entity.TaskComment, error) {
	const query = `
		INSERT INTO coopera.task_comments (task_id, user_id, username, text, created_at)
		VALUES ($1, $2, $3, $4, NOW())
		RETURNING id, task_id, user_id, username, text, created_at
	`
	tx, ok := ctx.Value(postgres.TransactionKey{}).(postgres.Transaction)
	if !ok {
		return entity.TaskComment{}, repoErr.ErrTransactionNotFound
	}
	var m task_comment_model.TaskComment
	if err := tx.QueryRow(ctx, query, c.TaskID, c.UserID, c.Username, c.Text).Scan(
		&m.ID, &m.TaskID, &m.UserID, &m.Username, &m.Text, &m.CreatedAt,
	); err != nil {
		return entity.TaskComment{}, fmt.Errorf("%w: %v", repoErr.ErrFailCreate, err)
	}
	return entity.TaskComment{
		ID: m.ID, TaskID: m.TaskID, UserID: m.UserID,
		Username: m.Username, Text: m.Text, CreatedAt: m.CreatedAt,
	}, nil
}

func (d *TaskCommentDAO) Delete(ctx context.Context, id int32, userID int32) error {
	const query = `DELETE FROM coopera.task_comments WHERE id = $1 AND user_id = $2`
	tx, ok := ctx.Value(postgres.TransactionKey{}).(postgres.Transaction)
	if !ok {
		return repoErr.ErrTransactionNotFound
	}
	if _, err := tx.Exec(ctx, query, id, userID); err != nil {
		return fmt.Errorf("%w: %v", repoErr.ErrFailDelete, err)
	}
	return nil
}
