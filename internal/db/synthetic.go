package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/easyCZ/qfy/internal/db/internal"
	"time"

	"gorm.io/gorm"
)

// Synthetic is the internal representation
type Synthetic struct {
	ID uint `json:"id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`

	Name string        `json:"name"`
	Spec SyntheticSpec `json:"spec"`
}

type SyntheticSpec struct {
	Variables []*Variable `json:"variables,omitempty"`
	Steps     []*StepSpec `json:"steps,omitempty"`
}

func (s *SyntheticSpec) ToJSON() ([]byte, error) {
	return json.Marshal(s)
}

type StepSpec struct {
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers,omitempty"`
	Body    *string           `json:"body,omitempty"`
}

type Variable struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

func (m *Synthetic) TableName() string {
	return "synthetics"
}

func NewSyntheticsRepository(db *gorm.DB) *SyntheticsRepository {
	return &SyntheticsRepository{db: db}
}

func NewSyntheticsRepositoryFromDBParams(params ConnectionParams) (*SyntheticsRepository, error) {
	conn, err := New(params)
	if err != nil {
		return nil, err
	}

	return NewSyntheticsRepository(conn), nil
}

type SyntheticsRepository struct {
	db *gorm.DB
}

func (r *SyntheticsRepository) List(ctx context.Context) ([]*Synthetic, error) {
	var results []*internal.Synthetic
	out := r.db.WithContext(ctx).Find(&results)
	if out.Error != nil {
		return nil, fmt.Errorf("failed to list synthetics from db: %w", out.Error)
	}

	var synthetics []*Synthetic
	for _, res := range results {
		syn, err := syntheticFromModel(res)
		if err != nil {
			return nil, err
		}

		synthetics = append(synthetics, syn)
	}

	return synthetics, nil
}

func (r *SyntheticsRepository) Get(ctx context.Context, id uint) (*Synthetic, error) {
	var model internal.Synthetic
	tx := r.db.WithContext(ctx).First(&model, id)
	if tx.Error != nil {
		return nil, fmt.Errorf("synthetic ID: %d does not exist", id)
	}

	return syntheticFromModel(&model)
}

func (r *SyntheticsRepository) Create(ctx context.Context, s *Synthetic) (*Synthetic, error) {
	spec, err := s.Spec.ToJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to convert spec to JSON: %w", err)
	}

	record := &internal.Synthetic{
		Name: s.Name,
		Spec: spec,
	}

	tx := r.db.WithContext(ctx).Create(record)
	if tx.Error != nil {
		return nil, fmt.Errorf("failed to create a new synthetic record: %w", tx.Error)
	}

	return syntheticFromModel(record)
}

func syntheticFromModel(s *internal.Synthetic) (*Synthetic, error) {
	var spec SyntheticSpec
	if err := json.Unmarshal(s.Spec, &spec); err != nil {
		return nil, fmt.Errorf("failed to convert spec to object: %w", err)
	}

	return &Synthetic{
		ID:        s.ID,
		DeletedAt: s.DeletedAt.Time,
		UpdatedAt: s.UpdatedAt,
		CreatedAt: s.CreatedAt,

		Name: s.Name,
		Spec: spec,
	}, nil
}
