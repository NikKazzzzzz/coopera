package team_repo

import (
	"context"

	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/converter"
	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/postgres/dao"
	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
)

type TeamRepository struct {
	dao dao.TeamDAO
}

func NewTeamRepository(dao dao.TeamDAO) *TeamRepository {
	return &TeamRepository{dao: dao}
}

func (r *TeamRepository) CreateRepo(ctx context.Context, e entity.TeamEntity) (entity.TeamEntity, error) {
	return r.dao.Create(ctx, converter.FromEntityToModelTeam(e))
}

func (r *TeamRepository) DeleteRepo(ctx context.Context, teamID int32) error {
	return r.dao.Delete(ctx, teamID)
}

func (r *TeamRepository) GetByIDRepo(ctx context.Context, teamID int32) (entity.TeamEntity, error) {
	return r.dao.GetByID(ctx, teamID)
}

func (r *TeamRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	return r.dao.ExistsByName(ctx, name)
}

func (r *TeamRepository) ExistsByID(ctx context.Context, teamID int32) (bool, error) {
	return r.dao.ExistsByID(ctx, teamID)
}

func (r *TeamRepository) UpdateMetaRepo(ctx context.Context, teamID int32, emoji, color string) error {
	return r.dao.UpdateMeta(ctx, teamID, emoji, color)
}

func (r *TeamRepository) UpdateAutoassignRepo(ctx context.Context, teamID int32, autoassign bool) error {
	return r.dao.UpdateAutoassign(ctx, teamID, autoassign)
}

func (r *TeamRepository) UpdatePhotoURLRepo(ctx context.Context, teamID int32, photoURL string) error {
	return r.dao.UpdatePhotoURL(ctx, teamID, photoURL)
}
