package task

import (
	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
)

type CreateTaskRequest struct {
	TeamID           int32    `json:"team_id" validate:"required"`
	Description      string   `json:"description" validate:"max=1000"`
	Points           int32    `json:"points" validate:"omitempty,gte=1"`
	CurrentUserID    int32    `json:"current_user_id" validate:"required"`
	Title            string   `json:"title" validate:"required,min=1,max=255"`
	AssignedToMember int32    `json:"assigned_to_member,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	Priority         string   `json:"priority,omitempty" validate:"omitempty,oneof=low medium high"`
}

type CreateTaskResponse struct {
	ID               int32    `json:"id"`
	TeamID           int32    `json:"team_id"`
	Title            string   `json:"title"`
	Description      *string  `json:"description,omitempty"`
	Points           *int32   `json:"points,omitempty"`
	Status           string   `json:"status,omitempty"`
	CreatedByUser    int32    `json:"created_by_user"`
	CreatedByMember  int32    `json:"created_by_member,omitempty"`
	AssignedToMember *int32   `json:"assigned_to_member,omitempty"`
	CreatedAt        string   `json:"created_at"`
	UpdatedAt        *string  `json:"updated_at"`
	Tags             []string `json:"tags"`
	Priority         string   `json:"priority"`
}

func ToEntityCreateTaskRequest(req *CreateTaskRequest) *entity.Task {
	task := &entity.Task{
		TeamID:          req.TeamID,
		CreatedByUserID: req.CurrentUserID,
		Title:           req.Title,
		Tags:            req.Tags,
		Priority:        req.Priority,
	}

	if task.Tags == nil {
		task.Tags = []string{}
	}
	if task.Priority == "" {
		task.Priority = "low"
	}

	if req.AssignedToMember != 0 {
		task.AssignedToMember = &req.AssignedToMember
	}

	if req.Description != "" {
		task.Description = &req.Description
	}

	if req.Points != 0 {
		task.Points = &req.Points
	}

	return task
}

func ToCreateTaskResponse(task *entity.Task) *CreateTaskResponse {
	taskResponse := &CreateTaskResponse{
		ID:              task.ID,
		TeamID:          task.TeamID,
		Title:           task.Title,
		Status:          task.Status.String(),
		CreatedByMember: task.CreatedByMemberID,
		CreatedByUser:   task.CreatedByUserID,
		CreatedAt:       task.CreatedAt.Format("2006-01-02T15:04:05Z"),
		Tags:            task.Tags,
		Priority:        task.Priority,
	}

	if taskResponse.Tags == nil {
		taskResponse.Tags = []string{}
	}

	if task.Points != nil {
		taskResponse.Points = task.Points
	}

	if task.Description != nil {
		taskResponse.Description = task.Description
	}

	if task.AssignedToMember != nil {
		taskResponse.AssignedToMember = task.AssignedToMember
	}

	if task.UpdatedAt != nil {
		updatedAt := task.UpdatedAt.Format("2006-01-02T15:04:05Z")
		taskResponse.UpdatedAt = &updatedAt
	}

	return taskResponse
}
