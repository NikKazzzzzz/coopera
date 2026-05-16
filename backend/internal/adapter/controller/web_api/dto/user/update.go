package user

type UpdateUserSettingsRequest struct {
	UserID             int32  `json:"user_id" validate:"required"`
	Wallpaper          string `json:"wallpaper"`
	WallpaperCustomURL string `json:"wallpaper_custom_url"`
	Theme              string `json:"theme"`
}
