package db

import "gorm.io/gorm"

type Synthetic struct {
	gorm.Model
	Name string
}
