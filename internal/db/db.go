package db

import (
	"fmt"
	"github.com/easyCZ/qfy/internal/db/internal"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnectionParams struct {
	Host         string
	Port         uint
	User         string
	Password     string
	DatabaseName string
}

func New(params ConnectionParams) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", params.Host, params.Port, params.User, params.Password, params.DatabaseName)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&internal.Synthetic{})
}
