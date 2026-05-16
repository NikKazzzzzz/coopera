package activity

import (
	"context"

	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
	"github.com/NikKazzzzzz/coopera-backend/internal/usecase"
)

type ActivityUsecase struct {
	txManager          usecase.TransactionManageRepository
	activityRepository usecase.ActivityRepository
}

func NewActivityUsecase(activityRepo usecase.ActivityRepository, txManager usecase.TransactionManageRepository) *ActivityUsecase {
	return &ActivityUsecase{txManager: txManager, activityRepository: activityRepo}
}

func (uc *ActivityUsecase) CreateUsecase(ctx context.Context, a entity.ActivityEntry) (entity.ActivityEntry, error) {
	var created entity.ActivityEntry
	err := uc.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		var err error
		created, err = uc.activityRepository.CreateRepo(txCtx, a)
		return err
	})
	return created, err
}

func (uc *ActivityUsecase) GetByUserIDUsecase(ctx context.Context, userID int32) ([]entity.ActivityEntry, error) {
	// NOTE: GetByUserID doesn't need a transaction, uses direct pool query
	return uc.activityRepository.GetByUserIDRepo(ctx, userID)
}

func (uc *ActivityUsecase) MarkAllReadUsecase(ctx context.Context, userID int32) error {
	return uc.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		return uc.activityRepository.MarkAllReadRepo(txCtx, userID)
	})
}

func (uc *ActivityUsecase) MarkSingleReadUsecase(ctx context.Context, id int32, userID int32) error {
	return uc.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		return uc.activityRepository.MarkSingleReadRepo(txCtx, id, userID)
	})
}

func (uc *ActivityUsecase) DeleteAllUsecase(ctx context.Context, userID int32) error {
	return uc.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		return uc.activityRepository.DeleteAllRepo(txCtx, userID)
	})
}
