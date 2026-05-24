package user

import (
	"context"
	"fmt"
	"log"

	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
	"github.com/NikKazzzzzz/coopera-backend/internal/usecase"
)

type PhotoFetcher interface {
	GetPhotoURL(userID int64) (string, error)
}

type UserUsecase struct {
	txManager      usecase.TransactionManageRepository
	userRepository usecase.UserRepository
	photoFetcher   PhotoFetcher
}

func NewUserUsecase(userRepository usecase.UserRepository, txManager usecase.TransactionManageRepository, photoFetcher PhotoFetcher) *UserUsecase {
	return &UserUsecase{
		txManager:      txManager,
		userRepository: userRepository,
		photoFetcher:   photoFetcher,
	}
}

func (uc *UserUsecase) CreateUsecase(ctx context.Context, euser entity.UserEntity) (entity.UserEntity, error) {
	if uc.photoFetcher != nil && euser.TelegramID != nil {
		url, err := uc.photoFetcher.GetPhotoURL(*euser.TelegramID)
		if err != nil {
			log.Printf("tgphoto fetch warning (user %d): %v", *euser.TelegramID, err)
		} else if url != "" {
			euser.PhotoURL = &url
		}
	}

	var createdUser entity.UserEntity

	err := uc.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		var err error
		createdUser, err = uc.userRepository.CreateRepo(txCtx, euser)
		if err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
		return nil
	})
	if err != nil {
		return entity.UserEntity{}, err
	}

	return createdUser, nil
}

func (uc *UserUsecase) GetUsecase(ctx context.Context, opts ...any) (entity.UserEntity, error) {
	var user entity.UserEntity

	err := uc.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		var err error
		user, err = uc.userRepository.GetRepo(txCtx, opts...)
		if err != nil {
			return fmt.Errorf("failed to get user: %w", err)
		}
		return nil
	})
	if err != nil {
		return entity.UserEntity{}, err
	}

	return user, nil
}

func (uc *UserUsecase) DeleteUsecase(ctx context.Context, userID int32) error {
	return uc.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		if err := uc.userRepository.DeleteRepo(txCtx, userID); err != nil {
			return fmt.Errorf("failed to delete user: %w", err)
		}
		return nil
	})
}

func (uc *UserUsecase) UpdateSettingsUsecase(ctx context.Context, userID int32, wallpaper, customURL, theme string) error {
	return uc.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		return uc.userRepository.UpdateSettingsRepo(txCtx, userID, wallpaper, customURL, theme)
	})
}

func (uc *UserUsecase) RefreshPhotoUsecase(ctx context.Context, userID int32) (*string, error) {
	user, err := uc.GetUsecase(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user.TelegramID == nil {
		return nil, nil
	}

	if uc.photoFetcher == nil {
		return nil, nil
	}

	newURL, err := uc.photoFetcher.GetPhotoURL(*user.TelegramID)
	if err != nil {
		log.Printf("tgphoto refresh warning (user %d): %v", *user.TelegramID, err)
		return nil, nil
	}

	if newURL == "" {
		return nil, nil
	}

	updateErr := uc.txManager.WithinTransaction(ctx, func(txCtx context.Context) error {
		return uc.userRepository.UpdatePhotoURLRepo(txCtx, userID, newURL)
	})
	if updateErr != nil {
		log.Printf("photo_url DB update warning (user %d): %v", userID, updateErr)
	}

	return &newURL, nil
}
