package team

import (
	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
	"time"
)

type GetTeamRequest struct {
	TeamID int32 `form:"team_id" validate:"required"`
}

type GetTeamResponse struct {
	ID            int32            `json:"id"`
	Name          string           `json:"name"`
	CreatedAt     string           `json:"created_at"`
	CreatedByUser int32            `json:"created_by_user"`
	Members       []TeamMemberInfo `json:"members"`
	Emoji         string           `json:"emoji"`
	Color         string           `json:"color"`
	Autoassign    bool             `json:"autoassign"`
}

type TeamMemberInfo struct {
	MemberID int32       `json:"member_id"`
	UserID   int32       `json:"user_id"`
	Username string      `json:"username"`
	Role     entity.Role `json:"role"`
}

func ToGetTeamResponse(team entity.TeamEntity, members []entity.MembershipEntity, usernames map[int32]string) GetTeamResponse {

	membersInfo := make([]TeamMemberInfo, 0, len(members))

	for _, m := range members {
		username, ok := usernames[m.UserID]
		if !ok {
			username = ""
		}

		membersInfo = append(membersInfo, TeamMemberInfo{
			MemberID: m.ID,
			UserID:   m.UserID,
			Username: username,
			Role:     m.Role,
		})
	}

	return GetTeamResponse{
		ID:            *team.ID,
		Name:          team.Name,
		CreatedAt:     team.CreatedAt.Format(time.RFC3339),
		CreatedByUser: team.CreatedBy,
		Members:       membersInfo,
		Emoji:         team.Emoji,
		Color:         team.Color,
		Autoassign:    team.Autoassign,
	}
}
