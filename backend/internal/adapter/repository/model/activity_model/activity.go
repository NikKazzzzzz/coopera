package activity_model

import "time"

type Activity struct {
	ID        int32
	UserID    int32
	TeamID    *int32
	TeamEmoji *string
	TeamColor *string
	Type      string
	Title     string
	Detail    string
	IsRead    bool
	CreatedAt time.Time
}
