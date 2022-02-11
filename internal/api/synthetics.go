package api

import (
	"context"
	"fmt"

	"github.com/easyCZ/qfy/internal/db"
)

func NewSyntheticsService(repo *db.SyntheticsRepository) *SyntheticsService {
	return &SyntheticsService{repo: repo}
}

type SyntheticsService struct {
	repo *db.SyntheticsRepository
}

func (s *SyntheticsService) List(ctx context.Context) ([]*db.Synthetic, error) {
	results, err := s.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list synthetics: %w", err)
	}

	return results, nil
}

func (s *SyntheticsService) Get(ctx context.Context, id uint) (*db.Synthetic, error) {
	result, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get synthetic: %w", err)
	}

	return result, nil
}
