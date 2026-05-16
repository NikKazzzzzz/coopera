package task_comment_dto

import "time"

type CreateCommentRequest struct {
	UserID   int32  `json:"user_id"  validate:"required"`
	Username string `json:"username" validate:"required"`
	Text     string `json:"text"     validate:"required"`
}

type DeleteCommentRequest struct {
	UserID int32 `form:"user_id" validate:"required"`
}

type CommentResponse struct {
	ID        int32     `json:"id"`
	TaskID    int32     `json:"task_id"`
	UserID    int32     `json:"user_id"`
	Username  string    `json:"username"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
