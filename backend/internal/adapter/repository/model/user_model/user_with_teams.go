package user_model

import "time"

type UserWithTeams struct {
	ID                 int32
	TelegramID         int64
	Username           string
	PhotoURL           *string
	CreatedAt          time.Time
	Teams              []TeamWithRole
	Wallpaper          string
	WallpaperCustomURL string
	Theme              string
}

type TeamWithRole struct {
	ID    int32
	Name  string
	Role  string
	Emoji string
	Color string
}
