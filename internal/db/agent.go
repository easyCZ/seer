package db

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Agent struct {
	ID        string
	Region    string
	IP        string
	Connected bool
}

type AgentsRepository struct {
	db *gorm.DB
}

func NewAgentsRepository(db *gorm.DB) *AgentsRepository {
	return &AgentsRepository{db: db}
}

func (r *AgentsRepository) Upsert(ctx context.Context, a *Agent) error {
	tx := r.db.WithContext(ctx).Clauses(
		clause.OnConflict{
			Columns: []clause.Column{
				{Name: "id"},
			},
			DoUpdates: clause.AssignmentColumns([]string{
				"region",
				"ip",
				"connected",
			}),
			UpdateAll: false,
		}).Create(a)
	if tx.Error != nil {
		return fmt.Errorf("failed to create agent recorc: %v", tx.Error)
	}

	return nil
}
