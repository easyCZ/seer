package internal

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Synthetic struct {
	gorm.Model
	Name string

	Spec datatypes.JSON
}
