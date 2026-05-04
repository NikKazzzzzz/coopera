package http

import (
	"context"

	"github.com/NikKazzzzzz/coopera-bot/internal/domain"
	"github.com/NikKazzzzzz/coopera-bot/internal/domain/memory"
	"github.com/NikKazzzzzz/coopera-bot/internal/domain/transport"
)

type cachedTeam struct {
	id     int64
	name   string
	client transport.Client
}

func (c cachedTeam) Details(_ context.Context) (domain.TeamDetails, error) {
	return memory.TeamDetails(c.id, c.name), nil
}

func (c cachedTeam) AddMember(ctx context.Context, user domain.User) (domain.Member, error) {
	// TODO implement me
	panic("implement me")
}

func (c cachedTeam) Members(ctx context.Context) ([]domain.Member, error) {
	// TODO implement me
	panic("implement me")
}
