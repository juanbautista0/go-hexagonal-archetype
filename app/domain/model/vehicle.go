package model

import "github.com/google/uuid"

type Vehicle struct {
	Id    uuid.UUID `json:"id"`
	Brand string    `json:"brand"`
	Model string    `json:"model"`
	Year  int       `json:"year"`
}
