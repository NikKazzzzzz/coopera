package web_api

import (
	"io"
	"net/http"
	"strconv"

	teamdto "github.com/NikKazzzzzz/coopera-backend/internal/adapter/controller/web_api/dto/team"
	"github.com/NikKazzzzzz/coopera-backend/internal/usecase"
	"github.com/NikKazzzzzz/coopera-backend/pkg/errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type TeamController struct {
	teamUseCase usecase.TeamUseCase
}

func NewTeamController(teamUseCase usecase.TeamUseCase) *TeamController {
	return &TeamController{
		teamUseCase: teamUseCase,
	}
}

func (tc *TeamController) Create(w http.ResponseWriter, r *http.Request) error {
	var req teamdto.CreateTeamRequest
	if err := BindRequest(r, &req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return errors.WrapValidationError(ve)
		}
		return errors.ErrInvalidInput
	}

	team, err := tc.teamUseCase.CreateUsecase(r.Context(), *teamdto.ToEntityCreateTeamRequest(&req))
	if err != nil {
		return err
	}

	writeJSON(w, http.StatusCreated, teamdto.ToCreateTeamResponse(&team))
	return nil
}

func (tc *TeamController) Get(w http.ResponseWriter, r *http.Request) error {
	var req teamdto.GetTeamRequest
	if err := BindRequest(r, &req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return errors.WrapValidationError(ve)
		}
		return errors.ErrInvalidInput
	}

	team, membership, username, err := tc.teamUseCase.GetByIDUsecase(r.Context(), req.TeamID)
	if err != nil {
		return err
	}

	writeJSON(w, http.StatusOK, teamdto.ToGetTeamResponse(team, membership, username))
	return nil
}

func (tc *TeamController) UpdateMeta(w http.ResponseWriter, r *http.Request) error {
	var req teamdto.UpdateTeamMetaRequest
	if err := BindRequest(r, &req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return errors.WrapValidationError(ve)
		}
		return errors.ErrInvalidInput
	}

	if err := tc.teamUseCase.UpdateMetaUsecase(r.Context(), req.TeamID, req.CurrentUserID, req.Emoji, req.Color); err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (tc *TeamController) UpdateAutoassign(w http.ResponseWriter, r *http.Request) error {
	var req teamdto.UpdateTeamAutoassignRequest
	if err := BindRequest(r, &req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return errors.WrapValidationError(ve)
		}
		return errors.ErrInvalidInput
	}
	if err := tc.teamUseCase.UpdateAutoassignUsecase(r.Context(), req.TeamID, req.CurrentUserID, req.Autoassign); err != nil {
		return err
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (tc *TeamController) GetPhoto(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return errors.ErrInvalidInput
	}

	team, _, _, err := tc.teamUseCase.GetByIDUsecase(r.Context(), int32(id))
	if err != nil {
		return err
	}

	if team.PhotoURL == nil || *team.PhotoURL == "" {
		http.NotFound(w, r)
		return nil
	}

	resp, err := photoClient.Get(*team.PhotoURL)
	if err != nil || resp.StatusCode != http.StatusOK {
		http.NotFound(w, r)
		return nil
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

func (tc *TeamController) UpdatePhoto(w http.ResponseWriter, r *http.Request) error {
	var req teamdto.UpdateTeamPhotoRequest
	if err := BindRequest(r, &req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return errors.WrapValidationError(ve)
		}
		return errors.ErrInvalidInput
	}

	if err := tc.teamUseCase.UpdatePhotoURLUsecase(r.Context(), req.TeamID, req.CurrentUserID, req.PhotoURL); err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (tc *TeamController) Delete(w http.ResponseWriter, r *http.Request) error {
	var req teamdto.DeleteTeamRequest
	if err := BindRequest(r, &req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return errors.WrapValidationError(ve)
		}
		return errors.ErrInvalidInput
	}

	if err := tc.teamUseCase.DeleteUsecase(r.Context(), req.TeamID, req.CurrentUserID); err != nil {
		return err
	}

	writeJSON(w, http.StatusNoContent, nil)
	return nil
}
