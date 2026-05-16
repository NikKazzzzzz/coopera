package dao

import (
	"context"
	"fmt"

	repoErr "github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/errors"
	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/model/activity_model"
	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/postgres"
	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
)

type ActivityDAO struct {
	db *postgres.DB
}

func NewActivityDAO(db *postgres.DB) *ActivityDAO {
	return &ActivityDAO{db: db}
}

func (d *ActivityDAO) Create(ctx context.Context, a activity_model.Activity) (entity.ActivityEntry, error) {
	const query = `
		INSERT INTO coopera.user_activity (user_id, team_id, team_emoji, team_color, type, title, detail, is_read, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, FALSE, NOW())
		RETURNING id, user_id, team_id, team_emoji, team_color, type, title, detail, is_read, created_at
	`
	tx, ok := ctx.Value(postgres.TransactionKey{}).(postgres.Transaction)
	if !ok {
		return entity.ActivityEntry{}, repoErr.ErrTransactionNotFound
	}
	var m activity_model.Activity
	if err := tx.QueryRow(ctx, query, a.UserID, a.TeamID, a.TeamEmoji, a.TeamColor, a.Type, a.Title, a.Detail).Scan(
		&m.ID, &m.UserID, &m.TeamID, &m.TeamEmoji, &m.TeamColor, &m.Type, &m.Title, &m.Detail, &m.IsRead, &m.CreatedAt,
	); err != nil {
		return entity.ActivityEntry{}, fmt.Errorf("%w: %v", repoErr.ErrFailCreate, err)
	}
	return entity.ActivityEntry{
		ID: m.ID, UserID: m.UserID, TeamID: m.TeamID, TeamEmoji: m.TeamEmoji, TeamColor: m.TeamColor,
		Type: m.Type, Title: m.Title, Detail: m.Detail, IsRead: m.IsRead, CreatedAt: m.CreatedAt,
	}, nil
}

func (d *ActivityDAO) GetByUserID(ctx context.Context, userID int32) ([]entity.ActivityEntry, error) {
	const query = `
		SELECT a.id, a.user_id, a.team_id,
		       COALESCE(NULLIF(a.team_emoji, ''), t.emoji) AS team_emoji,
		       COALESCE(NULLIF(a.team_color, ''), t.color) AS team_color,
		       a.type, a.title, a.detail, a.is_read, a.created_at
		FROM coopera.user_activity a
		LEFT JOIN coopera.teams t ON t.id = a.team_id
		WHERE a.user_id = $1
		ORDER BY a.created_at DESC
		LIMIT 50
	`
	rows, err := d.db.Pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", repoErr.ErrFailGet, err)
	}
	defer rows.Close()

	var result []entity.ActivityEntry
	for rows.Next() {
		var m activity_model.Activity
		if err := rows.Scan(&m.ID, &m.UserID, &m.TeamID, &m.TeamEmoji, &m.TeamColor, &m.Type, &m.Title, &m.Detail, &m.IsRead, &m.CreatedAt); err != nil {
			return nil, fmt.Errorf("%w: %v", repoErr.ErrFailGet, err)
		}
		result = append(result, entity.ActivityEntry{
			ID: m.ID, UserID: m.UserID, TeamID: m.TeamID, TeamEmoji: m.TeamEmoji, TeamColor: m.TeamColor,
			Type: m.Type, Title: m.Title, Detail: m.Detail, IsRead: m.IsRead, CreatedAt: m.CreatedAt,
		})
	}
	return result, nil
}

func (d *ActivityDAO) MarkAllRead(ctx context.Context, userID int32) error {
	const query = `UPDATE coopera.user_activity SET is_read = TRUE WHERE user_id = $1`
	tx, ok := ctx.Value(postgres.TransactionKey{}).(postgres.Transaction)
	if !ok {
		return repoErr.ErrTransactionNotFound
	}
	if _, err := tx.Exec(ctx, query, userID); err != nil {
		return fmt.Errorf("%w: %v", repoErr.ErrFailUpdate, err)
	}
	return nil
}

func (d *ActivityDAO) MarkSingleRead(ctx context.Context, id int32, userID int32) error {
	const query = `UPDATE coopera.user_activity SET is_read = TRUE WHERE id = $1 AND user_id = $2`
	tx, ok := ctx.Value(postgres.TransactionKey{}).(postgres.Transaction)
	if !ok {
		return repoErr.ErrTransactionNotFound
	}
	if _, err := tx.Exec(ctx, query, id, userID); err != nil {
		return fmt.Errorf("%w: %v", repoErr.ErrFailUpdate, err)
	}
	return nil
}

func (d *ActivityDAO) DeleteAll(ctx context.Context, userID int32) error {
	const query = `DELETE FROM coopera.user_activity WHERE user_id = $1`
	tx, ok := ctx.Value(postgres.TransactionKey{}).(postgres.Transaction)
	if !ok {
		return repoErr.ErrTransactionNotFound
	}
	if _, err := tx.Exec(ctx, query, userID); err != nil {
		return fmt.Errorf("%w: %v", repoErr.ErrFailUpdate, err)
	}
	return nil
}
