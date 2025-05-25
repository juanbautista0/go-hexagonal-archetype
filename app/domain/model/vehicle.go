package model

import (
	"github.com/juanbautista0/go-hexagonal-archetype/app/libraries"
)

type Vehicle struct {
	Id    libraries.Uuid `json:"id"`
	Brand string         `json:"brand"`
	Model string         `json:"model"`
	Year  int            `json:"year"`
}
