package entity

import "time"

type TaskComment struct {
	ID        int32
	TaskID    int32
	UserID    int32
	Username  string
	Text      string
	CreatedAt time.Time
}
