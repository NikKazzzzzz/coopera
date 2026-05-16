package user

import (
	"fmt"
	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
)

type GetUserRequest struct {
	ID         int32  `form:"id" validate:"omitempty"`
	TelegramID int64  `form:"telegram_id" validate:"omitempty"`
	UserName   string `form:"username" validate:"omitempty,max=32"`
}

type TeamInfo struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Emoji string `json:"emoji"`
	Color string `json:"color"`
}

type GetUserResponse struct {
	ID                 int32      `json:"id"`
	TelegramID         int64      `json:"telegram_id"`
	Username           string     `json:"username"`
	PhotoURL           *string    `json:"photo_url,omitempty"`
	CreatedAt          string     `json:"created_at"`
	Teams              []TeamInfo `json:"teams"`
	Wallpaper          string     `json:"wallpaper"`
	WallpaperCustomURL string     `json:"wallpaper_custom_url"`
	Theme              string     `json:"theme"`
}

func ToGetUserResponse(user *entity.UserEntity) *GetUserResponse {
	teams := make([]TeamInfo, len(user.Teams))
	for i, t := range user.Teams {
		teams[i] = TeamInfo{
			ID:    t.TeamID,
			Name:  t.TeamName,
			Role:  string(t.Role),
			Emoji: t.TeamEmoji,
			Color: t.TeamColor,
		}
	}

	return &GetUserResponse{
		ID:                 *user.ID,
		TelegramID:         *user.TelegramID,
		Username:           *user.Username,
		PhotoURL:           user.PhotoURL,
		CreatedAt:          user.CreatedAt.Format("2006-01-02T15:04:05Z"),
		Teams:              teams,
		Wallpaper:          user.Wallpaper,
		WallpaperCustomURL: user.WallpaperCustomURL,
		Theme:              user.Theme,
	}
}

func (r *GetUserRequest) Validate() error {
	if r.TelegramID == 0 && r.UserName == "" && r.ID == 0 {
		return fmt.Errorf("either telegram_id or username must or id be provided")
	}
	return nil
}
