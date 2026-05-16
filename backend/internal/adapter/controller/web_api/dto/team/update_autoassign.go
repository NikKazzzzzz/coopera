package team

type UpdateTeamAutoassignRequest struct {
	TeamID        int32 `json:"team_id" validate:"required"`
	CurrentUserID int32 `json:"current_user_id" validate:"required"`
	Autoassign    bool  `json:"autoassign"`
}
