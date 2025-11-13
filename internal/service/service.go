package service

import (
	"context"

	"github.com/YusovID/pr-reviewer-service/internal/repository"
	"github.com/YusovID/pr-reviewer-service/pkg/api"
)

type TeamService interface {
	CreateTeam(ctx context.Context, team api.Team) (*api.Team, error)
	GetTeam(ctx context.Context, name string) (*api.Team, error)
}

type UserService interface {
	SetIsActive(ctx context.Context, userID string, isActive bool) (*api.User, error)
}

type TeamServiceImpl struct {
	repo repository.TeamRepository
}

func NewTeamService(repo repository.TeamRepository) *TeamServiceImpl {
	return &TeamServiceImpl{repo: repo}
}

type UserServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}
