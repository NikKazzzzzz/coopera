package web_api

import (
	"net/http"
	"strconv"

	tcdto "github.com/NikKazzzzzz/coopera-backend/internal/adapter/controller/web_api/dto/task_comment"
	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
	"github.com/NikKazzzzzz/coopera-backend/internal/usecase"
	"github.com/NikKazzzzzz/coopera-backend/pkg/errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type TaskCommentController struct {
	uc usecase.TaskCommentUseCase
}

func NewTaskCommentController(uc usecase.TaskCommentUseCase) *TaskCommentController {
	return &TaskCommentController{uc: uc}
}

func (tc *TaskCommentController) GetComments(w http.ResponseWriter, r *http.Request) error {
	taskIDStr := chi.URLParam(r, "taskId")
	taskID, err := strconv.ParseInt(taskIDStr, 10, 32)
	if err != nil {
		return errors.ErrInvalidInput
	}
	comments, err := tc.uc.GetByTaskIDUsecase(r.Context(), int32(taskID))
	if err != nil {
		return err
	}
	resp := make([]tcdto.CommentResponse, len(comments))
	for i, c := range comments {
		resp[i] = tcdto.CommentResponse{
			ID: c.ID, TaskID: c.TaskID, UserID: c.UserID,
			Username: c.Username, Text: c.Text, CreatedAt: c.CreatedAt,
		}
	}
	writeJSON(w, http.StatusOK, resp)
	return nil
}

func (tc *TaskCommentController) CreateComment(w http.ResponseWriter, r *http.Request) error {
	taskIDStr := chi.URLParam(r, "taskId")
	taskID, err := strconv.ParseInt(taskIDStr, 10, 32)
	if err != nil {
		return errors.ErrInvalidInput
	}
	var req tcdto.CreateCommentRequest
	if err := BindRequest(r, &req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return errors.WrapValidationError(ve)
		}
		return errors.ErrInvalidInput
	}
	created, err := tc.uc.CreateUsecase(r.Context(), entity.TaskComment{
		TaskID:   int32(taskID),
		UserID:   req.UserID,
		Username: req.Username,
		Text:     req.Text,
	})
	if err != nil {
		return err
	}
	writeJSON(w, http.StatusCreated, tcdto.CommentResponse{
		ID: created.ID, TaskID: created.TaskID, UserID: created.UserID,
		Username: created.Username, Text: created.Text, CreatedAt: created.CreatedAt,
	})
	return nil
}

func (tc *TaskCommentController) DeleteComment(w http.ResponseWriter, r *http.Request) error {
	commentIDStr := chi.URLParam(r, "commentId")
	commentID, err := strconv.ParseInt(commentIDStr, 10, 32)
	if err != nil {
		return errors.ErrInvalidInput
	}
	var req tcdto.DeleteCommentRequest
	if err := BindRequest(r, &req); err != nil {
		if ve, ok := err.(validator.ValidationErrors); ok {
			return errors.WrapValidationError(ve)
		}
		return errors.ErrInvalidInput
	}
	if err := tc.uc.DeleteUsecase(r.Context(), int32(commentID), req.UserID); err != nil {
		return err
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}
