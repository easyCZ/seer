package db

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Agent struct {
	gorm.Model
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

func (r *AgentsRepository) Create(ctx context.Context, a *Agent) (*Agent, error) {
	tx := r.db.WithContext(ctx).Create(a)
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to create agent recorc: %v", tx.Error)
	}

	return a, nil
}

func (r *AgentsRepository) SetConnected(ctx context.Context, agentID uint, connected bool) error {
	tx := r.db.Model(&Agent{}).Where(&Agent{
		Model: gorm.Model{
			ID: agentID,
		},
	}).UpdateColumn("connected", false)
	if tx.Error != nil {
		return fmt.Errorf("failed to set agent connected field: %v", tx.Error)
	}
	return nil
}

func (r *AgentsRepository) Get(ctx context.Context, id uint) (*Agent, error) {
	var agent Agent
	tx := r.db.First(&agent, id)
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to lookup agent with ID %d: %w", id, tx.Error)
	}

	return &agent, nil
}
