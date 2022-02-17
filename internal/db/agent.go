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

func (r *AgentsRepository) Upsert(ctx context.Context, a *Agent) (*Agent, error) {
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
		return nil, fmt.Errorf("failed to create agent recorc: %v", tx.Error)
	}

	return a, nil
}

func (r *AgentsRepository) SetConnected(ctx context.Context, agentID string, connected bool) error {
	tx := r.db.Model(&Agent{}).Where(&Agent{
		ID: agentID,
	}).UpdateColumn("connected", false)
	if tx.Error != nil {
		return fmt.Errorf("failed to set agent connected field: %v", tx.Error)
	}
	return nil
}

func (r *AgentsRepository) Get(ctx context.Context, id string) (*Agent, error) {
	var agent Agent
	tx := r.db.First(&agent, id)
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to lookup agent with ID %s: %w", id, tx.Error)
	}

	return &agent, nil
}
