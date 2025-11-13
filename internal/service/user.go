package service

import (
	"context"
	"fmt"

	"github.com/YusovID/pr-reviewer-service/pkg/api"
)

func (s *UserServiceImpl) SetIsActive(ctx context.Context, userID string, isActive bool) (*api.User, error) {
	user, err := s.repo.SetIsActive(ctx, userID, isActive)
	if err != nil {
		return nil, fmt.Errorf("repo.SetIsActive failed: %w", err)
	}
	return user, nil
}
