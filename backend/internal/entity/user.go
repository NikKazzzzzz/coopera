package entity

import "time"

type UserEntity struct {
	ID                 *int32
	TelegramID         *int64
	Username           *string
	PhotoURL           *string
	CreatedAt          *time.Time
	Teams              []TeamWithRole
	Wallpaper          string
	WallpaperCustomURL string
	Theme              string
}

type TeamWithRole struct {
	TeamID    int32
	TeamName  string
	Role      Role
	TeamEmoji string
	TeamColor string
}

type UpdateUserSettings struct {
	UserID             int32
	Wallpaper          string
	WallpaperCustomURL string
	Theme              string
}
