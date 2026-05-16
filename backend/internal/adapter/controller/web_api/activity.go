package web_api

import (
	"net/http"
	"strconv"

	actdto "github.com/NikKazzzzzz/coopera-backend/internal/adapter/controller/web_api/dto/activity"
	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
	"github.com/NikKazzzzzz/coopera-backend/internal/usecase"
	"github.com/NikKazzzzzz/coopera-backend/pkg/errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type ActivityController struct {
	activityUseCase usecase.ActivityUseCase
}

func NewActivityController(activityUseCase usecase.ActivityUseCase) *ActivityController {
	return &ActivityController{activityUseCase: activityUseCase}
}

func (ac *ActivityController) Create(w http.ResponseWriter, r *http.Request) error {
	var req actdto.CreateActivityRequest
	if err := BindRequest(r, &req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return errors.WrapValidationError(ve)
		}
		return errors.ErrInvalidInput
	}
	entry, err := ac.activityUseCase.CreateUsecase(r.Context(), entity.ActivityEntry{
		UserID:    req.UserID,
		TeamID:    req.TeamID,
		TeamEmoji: req.TeamEmoji,
		TeamColor: req.TeamColor,
		Type:      req.Type,
		Title:     req.Title,
		Detail:    req.Detail,
	})
	if err != nil {
		return err
	}
	writeJSON(w, http.StatusCreated, actdto.ActivityResponse{
		ID: entry.ID, UserID: entry.UserID, TeamID: entry.TeamID, Type: entry.Type,
		Title: entry.Title, Detail: entry.Detail, IsRead: entry.IsRead, CreatedAt: entry.CreatedAt,
	})
	return nil
}

func (ac *ActivityController) GetByUser(w http.ResponseWriter, r *http.Request) error {
	var req actdto.GetActivityRequest
	if err := BindRequest(r, &req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return errors.WrapValidationError(ve)
		}
		return errors.ErrInvalidInput
	}
	entries, err := ac.activityUseCase.GetByUserIDUsecase(r.Context(), req.UserID)
	if err != nil {
		return err
	}
	resp := make([]actdto.ActivityResponse, len(entries))
	for i, e := range entries {
		resp[i] = actdto.ActivityResponse{
			ID: e.ID, UserID: e.UserID, TeamID: e.TeamID, TeamEmoji: e.TeamEmoji, TeamColor: e.TeamColor,
			Type: e.Type, Title: e.Title, Detail: e.Detail, IsRead: e.IsRead, CreatedAt: e.CreatedAt,
		}
	}
	writeJSON(w, http.StatusOK, resp)
	return nil
}

func (ac *ActivityController) MarkAllRead(w http.ResponseWriter, r *http.Request) error {
	var req actdto.MarkReadRequest
	if err := BindRequest(r, &req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return errors.WrapValidationError(ve)
		}
		return errors.ErrInvalidInput
	}
	if err := ac.activityUseCase.MarkAllReadUsecase(r.Context(), req.UserID); err != nil {
		return err
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (ac *ActivityController) MarkSingleRead(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return errors.ErrInvalidInput
	}
	var req actdto.MarkSingleReadRequest
	if err := BindRequest(r, &req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return errors.WrapValidationError(ve)
		}
		return errors.ErrInvalidInput
	}
	if err := ac.activityUseCase.MarkSingleReadUsecase(r.Context(), int32(id), req.UserID); err != nil {
		return err
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (ac *ActivityController) DeleteAll(w http.ResponseWriter, r *http.Request) error {
	var req actdto.DeleteRequest
	if err := BindRequest(r, &req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return errors.WrapValidationError(ve)
		}
		return errors.ErrInvalidInput
	}
	if err := ac.activityUseCase.DeleteAllUsecase(r.Context(), req.UserID); err != nil {
		return err
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}
