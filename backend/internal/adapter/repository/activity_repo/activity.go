package activity_repo

import (
	"context"

	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/model/activity_model"
	"github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/postgres/dao"
	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
)

type ActivityRepository struct {
	dao *dao.ActivityDAO
}

func NewActivityRepository(activityDAO *dao.ActivityDAO) *ActivityRepository {
	return &ActivityRepository{dao: activityDAO}
}

func (r *ActivityRepository) CreateRepo(ctx context.Context, a entity.ActivityEntry) (entity.ActivityEntry, error) {
	return r.dao.Create(ctx, activity_model.Activity{
		UserID: a.UserID,
		Type:   a.Type,
		Title:  a.Title,
		Detail: a.Detail,
	})
}

func (r *ActivityRepository) GetByUserIDRepo(ctx context.Context, userID int32) ([]entity.ActivityEntry, error) {
	return r.dao.GetByUserID(ctx, userID)
}

func (r *ActivityRepository) MarkAllReadRepo(ctx context.Context, userID int32) error {
	return r.dao.MarkAllRead(ctx, userID)
}

func (r *ActivityRepository) MarkSingleReadRepo(ctx context.Context, id int32, userID int32) error {
	return r.dao.MarkSingleRead(ctx, id, userID)
}

func (r *ActivityRepository) DeleteAllRepo(ctx context.Context, userID int32) error {
	return r.dao.DeleteAll(ctx, userID)
}
