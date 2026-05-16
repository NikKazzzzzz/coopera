package dao

import (
	"context"
	"fmt"

	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/converter"
	repoErr "github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/errors"
	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/model/user_model"
	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/postgres"
	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
)

type UserRepository struct {
	db *postgres.DB
}

func NewUserDAO(db *postgres.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) Create(ctx context.Context, muser user_model.User) (entity.UserEntity, error) {
	// ON CONFLICT обновляет photo_url только если он ещё не был установлен
	const query = `
		INSERT INTO coopera.users (telegram_id, username, photo_url, created_at)
		VALUES ($1, $2, $3, NOW())
		ON CONFLICT (telegram_id) DO UPDATE
			SET photo_url = COALESCE(EXCLUDED.photo_url, coopera.users.photo_url)
		RETURNING id, telegram_id, username, photo_url, created_at
	`

	tx, ok := ctx.Value(postgres.TransactionKey{}).(postgres.Transaction)
	if !ok {
		return entity.UserEntity{}, repoErr.ErrTransactionNotFound
	}

	var userModel user_model.User
	err := tx.QueryRow(ctx, query, muser.TelegramID, muser.Username, muser.PhotoURL).Scan(
		&userModel.ID,
		&userModel.TelegramID,
		&userModel.Username,
		&userModel.PhotoURL,
		&userModel.CreatedAt,
	)
	if err != nil {
		return entity.UserEntity{}, fmt.Errorf("%w: %v", repoErr.ErrFailCreate, err)
	}

	return converter.FromModelToEntity(userModel), nil
}

func (ur *UserRepository) Delete(ctx context.Context, userID int32) error {
	const query = `
		DELETE FROM coopera.users
		WHERE id = $1
	`

	tx, ok := ctx.Value(postgres.TransactionKey{}).(postgres.Transaction)
	if !ok {
		return repoErr.ErrTransactionNotFound
	}

	cmdTag, err := tx.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("%w: %v", repoErr.ErrFailDelete, err)
	}

	if cmdTag.RowsAffected() == 0 {
		return repoErr.ErrNotFound
	}

	return nil
}

func (ur *UserRepository) GetByTelegramID(ctx context.Context, telegramID int64) (entity.UserEntity, error) {
	const query = `
		SELECT
			u.id, u.telegram_id, u.username, u.photo_url, u.created_at,
			u.wallpaper, u.wallpaper_custom_url, u.theme,
			t.id AS team_id, t.name AS team_name, m.role, t.emoji, t.color
		FROM coopera.users u
		LEFT JOIN coopera.memberships m ON m.user_id = u.id
		LEFT JOIN coopera.teams t ON t.id = m.team_id
		WHERE u.telegram_id = $1
`

	tx, ok := ctx.Value(postgres.TransactionKey{}).(postgres.Transaction)
	if !ok {
		return entity.UserEntity{}, repoErr.ErrTransactionNotFound
	}

	rows, err := tx.Query(ctx, query, telegramID)
	if err != nil {
		return entity.UserEntity{}, fmt.Errorf("%w: %v", repoErr.ErrFailGet, err)
	}
	defer rows.Close()

	var user user_model.UserWithTeams
	for rows.Next() {
		var (
			teamID   *int32
			teamName *string
			role     *string
			emoji    *string
			color    *string
		)

		if err := rows.Scan(
			&user.ID,
			&user.TelegramID,
			&user.Username,
			&user.PhotoURL,
			&user.CreatedAt,
			&user.Wallpaper,
			&user.WallpaperCustomURL,
			&user.Theme,
			&teamID,
			&teamName,
			&role,
			&emoji,
			&color,
		); err != nil {
			return entity.UserEntity{}, fmt.Errorf("%w: %v", repoErr.ErrFailGet, err)
		}

		if teamID != nil && teamName != nil && role != nil {
			t := user_model.TeamWithRole{
				ID:   *teamID,
				Name: *teamName,
				Role: *role,
			}
			if emoji != nil {
				t.Emoji = *emoji
			}
			if color != nil {
				t.Color = *color
			}
			user.Teams = append(user.Teams, t)
		}
	}

	if user.ID == 0 {
		return entity.UserEntity{}, repoErr.ErrNotFound
	}

	return converter.FromModelToEntityWithTeams(user), nil
}

func (ur *UserRepository) GetByUsername(ctx context.Context, username string) (entity.UserEntity, error) {
	const query = `
		SELECT
			u.id, u.telegram_id, u.username, u.photo_url, u.created_at,
			u.wallpaper, u.wallpaper_custom_url, u.theme,
			t.id AS team_id, t.name AS team_name, m.role, t.emoji, t.color
		FROM coopera.users u
		LEFT JOIN coopera.memberships m ON m.user_id = u.id
		LEFT JOIN coopera.teams t ON t.id = m.team_id
		WHERE u.username = $1
	`

	tx, ok := ctx.Value(postgres.TransactionKey{}).(postgres.Transaction)
	if !ok {
		return entity.UserEntity{}, repoErr.ErrTransactionNotFound
	}

	rows, err := tx.Query(ctx, query, username)
	if err != nil {
		return entity.UserEntity{}, fmt.Errorf("%w: %v", repoErr.ErrFailGet, err)
	}
	defer rows.Close()

	var user user_model.UserWithTeams
	for rows.Next() {
		var (
			teamID   *int32
			teamName *string
			role     *string
			emoji    *string
			color    *string
		)

		if err := rows.Scan(
			&user.ID,
			&user.TelegramID,
			&user.Username,
			&user.PhotoURL,
			&user.CreatedAt,
			&user.Wallpaper,
			&user.WallpaperCustomURL,
			&user.Theme,
			&teamID,
			&teamName,
			&role,
			&emoji,
			&color,
		); err != nil {
			return entity.UserEntity{}, fmt.Errorf("%w: %v", repoErr.ErrFailGet, err)
		}

		if teamID != nil && teamName != nil && role != nil {
			t := user_model.TeamWithRole{
				ID:   *teamID,
				Name: *teamName,
				Role: *role,
			}
			if emoji != nil {
				t.Emoji = *emoji
			}
			if color != nil {
				t.Color = *color
			}
			user.Teams = append(user.Teams, t)
		}
	}

	if user.ID == 0 {
		return entity.UserEntity{}, repoErr.ErrNotFound
	}

	return converter.FromModelToEntityWithTeams(user), nil
}

func (ur *UserRepository) GetByUserID(ctx context.Context, userID int32) (entity.UserEntity, error) {
	const query = `
		SELECT
			u.id, u.telegram_id, u.username, u.photo_url, u.created_at,
			u.wallpaper, u.wallpaper_custom_url, u.theme,
			t.id AS team_id, t.name AS team_name, m.role, t.emoji, t.color
		FROM coopera.users u
		LEFT JOIN coopera.memberships m ON m.user_id = u.id
		LEFT JOIN coopera.teams t ON t.id = m.team_id
		WHERE u.id = $1
	`

	tx, ok := ctx.Value(postgres.TransactionKey{}).(postgres.Transaction)
	if !ok {
		return entity.UserEntity{}, repoErr.ErrTransactionNotFound
	}

	rows, err := tx.Query(ctx, query, userID)
	if err != nil {
		return entity.UserEntity{}, fmt.Errorf("%w: %v", repoErr.ErrFailGet, err)
	}
	defer rows.Close()

	var user user_model.UserWithTeams

	for rows.Next() {
		var (
			teamID   *int32
			teamName *string
			role     *string
			emoji    *string
			color    *string
		)

		if err := rows.Scan(
			&user.ID,
			&user.TelegramID,
			&user.Username,
			&user.PhotoURL,
			&user.CreatedAt,
			&user.Wallpaper,
			&user.WallpaperCustomURL,
			&user.Theme,
			&teamID,
			&teamName,
			&role,
			&emoji,
			&color,
		); err != nil {
			return entity.UserEntity{}, fmt.Errorf("%w: %v", repoErr.ErrFailGet, err)
		}

		if teamID != nil && teamName != nil && role != nil {
			t := user_model.TeamWithRole{
				ID:   *teamID,
				Name: *teamName,
				Role: *role,
			}
			if emoji != nil {
				t.Emoji = *emoji
			}
			if color != nil {
				t.Color = *color
			}
			user.Teams = append(user.Teams, t)
		}
	}

	if user.ID == 0 {
		return entity.UserEntity{}, repoErr.ErrNotFound
	}

	return converter.FromModelToEntityWithTeams(user), nil
}

func (ur *UserRepository) UpdateSettings(ctx context.Context, userID int32, wallpaper, customURL, theme string) error {
	const query = `
		UPDATE coopera.users
		SET wallpaper = $1, wallpaper_custom_url = $2, theme = $3
		WHERE id = $4
	`
	tx, ok := ctx.Value(postgres.TransactionKey{}).(postgres.Transaction)
	if !ok {
		return repoErr.ErrTransactionNotFound
	}
	if _, err := tx.Exec(ctx, query, wallpaper, customURL, theme, userID); err != nil {
		return fmt.Errorf("%w: %v", repoErr.ErrFailUpdate, err)
	}
	return nil
}
