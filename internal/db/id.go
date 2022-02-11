package db

import (
	"k8s.io/apimachinery/pkg/util/rand"
)

func generateID(n int) string {
	return rand.String(8)
}

func NewSyntheticID() string {
	return generateID(8)
}
