package api

import (
	"context"
	"fmt"

	"github.com/easyCZ/qfy/internal/db"
)

type Synthetic struct {
	ID   uint
	Name string
}

func NewSyntheticsService(repo *db.SyntheticsRepository) *SyntheticsService {
	return &SyntheticsService{repo: repo}
}

type SyntheticsService struct {
	repo *db.SyntheticsRepository
}

func (s *SyntheticsService) List(ctx context.Context) ([]*Synthetic, error) {
	results, err := s.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list synthetics: %w", err)
	}

	var synthetics []*Synthetic
	for _, s := range results {
		synthetics = append(synthetics, &Synthetic{
			ID:   s.ID,
			Name: s.Name,
		})
	}

	return synthetics, nil
}
