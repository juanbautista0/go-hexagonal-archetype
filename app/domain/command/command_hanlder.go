package command

import (
	"context"

	"github.com/juanbautista0/go-hexagonal-archetype/app/domain/ports"
)

type CreateVehicleCommandHandler struct {
	repository ports.VehicleRepository
}

func NewCreateVehicleCommandHandler(r ports.VehicleRepository) *CreateVehicleCommandHandler {
	return &CreateVehicleCommandHandler{repository: r}
}

func (h *CreateVehicleCommandHandler) Execute(ctx context.Context, command CreateVehicleCommand) (bool, error) {
	if _, err := h.repository.Save(ctx, &command.Vehicle); err != nil {
		return false, err
	}
	return true, nil
}

func (h *CreateVehicleCommandHandler) Handle(ctx context.Context, event interface{}) (interface{}, error) {
	cmd := event.(CreateVehicleCommand)
	return h.Execute(ctx, cmd)
}
