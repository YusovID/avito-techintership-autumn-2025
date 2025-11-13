package repository

import (
	"context"

	"github.com/YusovID/pr-reviewer-service/internal/domain"
	"github.com/YusovID/pr-reviewer-service/pkg/api"
)

type TeamRepository interface {
	CreateTeamWithUsers(ctx context.Context, team api.Team) (*domain.TeamWithMembers, error)
	GetTeamByName(ctx context.Context, name string) (*domain.TeamWithMembers, error)
}

type UserRepository interface {
	SetIsActive(ctx context.Context, userID string, isActive bool) (*api.User, error)
}
