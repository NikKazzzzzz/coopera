package team

type UpdateTeamMetaRequest struct {
	TeamID        int32  `json:"team_id" validate:"required"`
	CurrentUserID int32  `json:"current_user_id" validate:"required"`
	Emoji         string `json:"emoji"`
	Color         string `json:"color"`
}
