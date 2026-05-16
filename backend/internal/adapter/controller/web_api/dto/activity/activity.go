package activity

import "time"

type CreateActivityRequest struct {
	UserID    int32   `json:"user_id" validate:"required"`
	TeamID    *int32  `json:"team_id,omitempty"`
	TeamEmoji *string `json:"team_emoji,omitempty"`
	TeamColor *string `json:"team_color,omitempty"`
	Type      string  `json:"type" validate:"required"`
	Title     string  `json:"title" validate:"required"`
	Detail    string  `json:"detail"`
}

type GetActivityRequest struct {
	UserID int32 `form:"user_id" validate:"required"`
}

type MarkReadRequest struct {
	UserID int32 `json:"user_id" validate:"required"`
}

type MarkSingleReadRequest struct {
	UserID int32 `json:"user_id" validate:"required"`
}

type DeleteRequest struct {
	UserID int32 `form:"user_id" validate:"required"`
}

type ActivityResponse struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	TeamID    *int32    `json:"team_id,omitempty"`
	TeamEmoji *string   `json:"team_emoji,omitempty"`
	TeamColor *string   `json:"team_color,omitempty"`
	Type      string    `json:"type"`
	Title     string    `json:"title"`
	Detail    string    `json:"detail"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}
