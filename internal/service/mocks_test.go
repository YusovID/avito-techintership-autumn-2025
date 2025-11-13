package service

import (
	"context"

	"github.com/YusovID/pr-reviewer-service/internal/domain"
	"github.com/YusovID/pr-reviewer-service/internal/repository"
	"github.com/YusovID/pr-reviewer-service/pkg/api"
	"github.com/stretchr/testify/mock"
)

type TeamRepositoryMock struct {
	mock.Mock
}

var _ repository.TeamRepository = (*TeamRepositoryMock)(nil)

func (m *TeamRepositoryMock) CreateTeamWithUsers(ctx context.Context, team api.Team) (*domain.TeamWithMembers, error) {
	args := m.Called(ctx, team)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.TeamWithMembers), args.Error(1)
}

func (m *TeamRepositoryMock) GetTeamByName(ctx context.Context, name string) (*domain.TeamWithMembers, error) {
	args := m.Called(ctx, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.TeamWithMembers), args.Error(1)
}

type UserRepositoryMock struct {
	mock.Mock
}

var _ repository.UserRepository = (*UserRepositoryMock)(nil)

func (m *UserRepositoryMock) SetIsActive(ctx context.Context, userID string, isActive bool) (*api.User, error) {
	args := m.Called(ctx, userID, isActive)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*api.User), args.Error(1)
}
