package team

type UpdateTeamPhotoRequest struct {
	TeamID        int32  `json:"team_id" validate:"required"`
	CurrentUserID int32  `json:"current_user_id" validate:"required"`
	PhotoURL      string `json:"photo_url" validate:"required"`
}
