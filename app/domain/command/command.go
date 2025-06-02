package command

import (
	"github.com/juanbautista0/go-hexagonal-archetype/app/domain/entity"
)

type CreateVehicleCommand struct {
	Vehicle entity.Vehicle `json:"vehicle"`
}
