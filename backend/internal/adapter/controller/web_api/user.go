package web_api

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/NikKazzzzzz/coopera-backend/pkg/errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"

	userdto "github.com/NikKazzzzzz/coopera-backend/internal/adapter/controller/web_api/dto/user"
	"github.com/NikKazzzzzz/coopera-backend/internal/usecase"
)

type UserController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(userUseCase usecase.UserUseCase) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) error {
	var req userdto.CreateUserRequest
	if err := BindRequest(r, &req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return errors.WrapValidationError(ve)
		}
		return errors.ErrInvalidInput
	}

	user, err := uc.userUseCase.CreateUsecase(r.Context(), *userdto.FromCreateUserRequest(&req))
	if err != nil {
		return err
	}

	writeJSON(w, http.StatusCreated, userdto.ToCreateUserResponse(&user))
	return nil
}

func (uc *UserController) Get(w http.ResponseWriter, r *http.Request) error {
	var req userdto.GetUserRequest
	if err := BindRequest(r, &req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return errors.WrapValidationError(ve)
		}
		return errors.ErrInvalidInput
	}

	if err := req.Validate(); err != nil {
		return fmt.Errorf("%w: %v", errors.ErrInvalidInput, err)
	}

	user, err := uc.userUseCase.GetUsecase(r.Context(), req.TelegramID, req.UserName, req.ID)
	if err != nil {
		return err
	}

	writeJSON(w, http.StatusOK, userdto.ToGetUserResponse(&user))
	return nil
}

var photoClient = &http.Client{Timeout: 15 * time.Second}

func (uc *UserController) GetPhoto(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return errors.ErrInvalidInput
	}

	user, err := uc.userUseCase.GetUsecase(r.Context(), int32(id))
	if err != nil {
		return err
	}

	var photoURL string
	if user.PhotoURL != nil && *user.PhotoURL != "" {
		photoURL = *user.PhotoURL
	}

	var resp *http.Response
	if photoURL != "" {
		resp, err = photoClient.Get(photoURL)
		if err != nil || resp.StatusCode != http.StatusOK {
			if resp != nil {
				resp.Body.Close()
			}
			resp = nil
		}
	}

	if resp == nil {
		freshURL, refreshErr := uc.userUseCase.RefreshPhotoUsecase(r.Context(), int32(id))
		if refreshErr != nil || freshURL == nil || *freshURL == "" {
			http.NotFound(w, r)
			return nil
		}
		resp, err = photoClient.Get(*freshURL)
		if err != nil || resp.StatusCode != http.StatusOK {
			if resp != nil {
				resp.Body.Close()
			}
			http.NotFound(w, r)
			return nil
		}
	}
	defer resp.Body.Close()

	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "image/jpeg"
	}
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Cache-Control", "public, max-age=86400")
	io.Copy(w, resp.Body)
	return nil
}

func (uc *UserController) Delete(w http.ResponseWriter, r *http.Request) error {
	var req userdto.DeleteUserRequest
	if err := BindRequest(r, &req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return errors.WrapValidationError(ve)
		}
		return errors.ErrInvalidInput
	}

	if err := uc.userUseCase.DeleteUsecase(r.Context(), req.ID); err != nil {
		return err
	}

	writeJSON(w, http.StatusNoContent, nil)
	return nil
}

func (uc *UserController) UpdateSettings(w http.ResponseWriter, r *http.Request) error {
	var req userdto.UpdateUserSettingsRequest
	if err := BindRequest(r, &req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return errors.WrapValidationError(ve)
		}
		return errors.ErrInvalidInput
	}
	if err := uc.userUseCase.UpdateSettingsUsecase(r.Context(), req.UserID, req.Wallpaper, req.WallpaperCustomURL, req.Theme); err != nil {
		return err
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}
