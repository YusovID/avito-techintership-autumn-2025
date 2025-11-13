package service

import (
	"context"
	"errors"
	"testing"

	"github.com/YusovID/pr-reviewer-service/internal/domain"
	"github.com/YusovID/pr-reviewer-service/pkg/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTeamServiceImpl_CreateTeam(t *testing.T) {
	ctx := context.Background()

	inputTeam := api.Team{
		TeamName: "test-team",
		Members: []api.TeamMember{
			{UserId: "u1", Username: "Test User", IsActive: true},
		},
	}

	domainTeamWithMembers := &domain.TeamWithMembers{
		ID:   1,
		Name: "test-team",
		Members: []domain.User{
			{ID: "u1", Username: "Test User", TeamID: 1, IsActive: true},
		},
	}

	testCases := []struct {
		name          string
		setupMock     func(repoMock *TeamRepositoryMock)
		inputTeam     api.Team
		expectedTeam  *api.Team
		expectedError bool
	}{
		{
			name: "Success: Team and users are created",
			setupMock: func(repoMock *TeamRepositoryMock) {
				repoMock.On("CreateTeamWithUsers", mock.Anything, inputTeam).Return(domainTeamWithMembers, nil)
			},
			inputTeam: inputTeam,
			expectedTeam: &api.Team{
				TeamName: "test-team",
				Members: []api.TeamMember{
					{UserId: "u1", Username: "Test User", IsActive: true},
				},
			},
			expectedError: false,
		},
		{
			name: "Failure: Repository returns error on CreateTeamWithUsers",
			setupMock: func(repoMock *TeamRepositoryMock) {
				repoMock.On("CreateTeamWithUsers", mock.Anything, inputTeam).Return(nil, errors.New("database connection lost"))
			},
			inputTeam:     inputTeam,
			expectedTeam:  nil,
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repoMock := new(TeamRepositoryMock)
			tc.setupMock(repoMock)

			service := NewTeamService(repoMock)

			resultTeam, err := service.CreateTeam(ctx, tc.inputTeam)

			assert.Equal(t, tc.expectedTeam, resultTeam)
			if tc.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			repoMock.AssertExpectations(t)
		})
	}
}
