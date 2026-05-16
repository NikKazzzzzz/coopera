package entity

import "time"

type TeamEntity struct {
	ID         *int32
	Name       string
	CreatedAt  *time.Time
	CreatedBy  int32
	Emoji      string
	Color      string
	Autoassign bool
	PhotoURL   *string
}

type UpdateTeamMeta struct {
	TeamID int32
	Emoji  string
	Color  string
}
