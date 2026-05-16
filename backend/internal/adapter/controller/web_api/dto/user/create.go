package user

import (
	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
	"time"
)

type CreateUserRequest struct {
	TelegramID int64   `json:"telegram_id" validate:"required"`
	Username   string  `json:"username" validate:"required,max=32"`
	PhotoURL   *string `json:"photo_url" validate:"omitempty,url"`
}

type CreateUserResponse struct {
	ID         int32     `json:"id"`
	TelegramID int64     `json:"telegram_id"`
	Username   string    `json:"username"`
	PhotoURL   *string   `json:"photo_url,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
}

func FromCreateUserRequest(req *CreateUserRequest) *entity.UserEntity {
	return &entity.UserEntity{
		TelegramID: &req.TelegramID,
		Username:   &req.Username,
		PhotoURL:   req.PhotoURL,
	}
}

func ToCreateUserResponse(user *entity.UserEntity) *CreateUserResponse {
	return &CreateUserResponse{
		ID:         *user.ID,
		TelegramID: *user.TelegramID,
		Username:   *user.Username,
		PhotoURL:   user.PhotoURL,
		CreatedAt:  user.CreatedAt.Truncate(time.Second),
	}
}
