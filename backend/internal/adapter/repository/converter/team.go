package converter

import (
	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/model/team_model"
	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
)

func FromEntityToModelTeam(team entity.TeamEntity) team_model.Team {
	return team_model.Team{
		ID:         0,
		Name:       team.Name,
		CreatedBy:  team.CreatedBy,
		Emoji:      team.Emoji,
		Color:      team.Color,
		Autoassign: team.Autoassign,
		PhotoURL:   team.PhotoURL,
	}
}

func FromModelToEntityTeam(team team_model.Team) entity.TeamEntity {
	return entity.TeamEntity{
		ID:         &team.ID,
		Name:       team.Name,
		CreatedBy:  team.CreatedBy,
		CreatedAt:  &team.CreatedAt,
		Emoji:      team.Emoji,
		Color:      team.Color,
		Autoassign: team.Autoassign,
		PhotoURL:   team.PhotoURL,
	}
}
