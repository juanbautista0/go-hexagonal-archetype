package command

import "github.com/juanbautista0/go-hexagonal-archetype/app/domain/model"

type CreateVehicleCommand struct {
	model.Vehicle
}
