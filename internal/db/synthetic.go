package db

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Synthetic struct {
	gorm.Model
	Name string
}

func NewSyntheticsRepository(db *gorm.DB) *SyntheticsRepository {
	return &SyntheticsRepository{db: db}
}

type SyntheticsRepository struct {
	db *gorm.DB
}

func (r *SyntheticsRepository) List(ctx context.Context) ([]*Synthetic, error) {
	var results []*Synthetic
	out := r.db.WithContext(ctx).Find(&results)
	if out.Error != nil {
		return nil, fmt.Errorf("failed to list synthetics from db: %w", out.Error)
	}

	return results, nil
}
