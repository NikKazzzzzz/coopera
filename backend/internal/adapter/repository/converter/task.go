package converter

import (
	taskModel "github.com/NikKazzzzzz/coopera-backend/internal/adapter/repository/model/task_model"
	"github.com/NikKazzzzzz/coopera-backend/internal/entity"
)

func FromEntityToModelTask(task entity.Task) taskModel.Task {
	mtask := taskModel.Task{
		ID:              task.ID,
		TeamID:          task.TeamID,
		Title:           task.Title,
		Description:     task.Description,
		Points:          task.Points,
		AssignedTo:      task.AssignedToMember,
		CreatedByUser:   task.CreatedByUserID,
		CreatedByMember: task.CreatedByMemberID,
		CreatedAt:       task.CreatedAt,
		UpdatedAt:       task.UpdatedAt,
		Tags:            task.Tags,
		Priority:        task.Priority,
	}

	if task.AssignedToMember != nil {
		mtask.Status = entity.StatusAssigned.String()
	} else {
		mtask.Status = entity.StatusOpen.String()
	}

	if mtask.Tags == nil {
		mtask.Tags = []string{}
	}
	if mtask.Priority == "" {
		mtask.Priority = "low"
	}

	return mtask
}

func FromModelToEntityTask(m taskModel.Task) entity.Task {
	status := entity.Status(m.Status)
	tags := m.Tags
	if tags == nil {
		tags = []string{}
	}
	return entity.Task{
		ID:                m.ID,
		TeamID:            m.TeamID,
		Title:             m.Title,
		Description:       m.Description,
		Points:            m.Points,
		Status:            &status,
		AssignedToMember:  m.AssignedTo,
		CreatedByUserID:   m.CreatedByUser,
		CreatedByMemberID: m.CreatedByMember,
		CreatedAt:         m.CreatedAt,
		UpdatedAt:         m.UpdatedAt,
		Tags:              tags,
		Priority:          m.Priority,
		CommentCount:      m.CommentCount,
	}
}

func FromEntityToModelTaskStatus(status entity.TaskStatus) taskModel.TaskStatus {
	return taskModel.TaskStatus{
		TaskID: status.TaskID,
		Status: status.Status,
	}
}

func FromEntityToModelUpdateTask(task entity.UpdateTask) taskModel.UpdateTask {
	return taskModel.UpdateTask{
		ID:          task.TaskID,
		Title:       task.Title,
		Description: task.Description,
		Points:      task.Points,
		AssignedTo:  task.AssignedToMember,
		Tags:        task.Tags,
		Priority:    task.Priority,
	}
}
