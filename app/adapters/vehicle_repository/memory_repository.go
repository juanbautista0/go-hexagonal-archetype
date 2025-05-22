package vehicle_repository

import (
	"context"

	"github.com/juanbautista0/go-hexagonal-archetype/app/domain/model"
	"github.com/juanbautista0/go-hexagonal-archetype/app/domain/ports"
)

type MemoryRepository struct {
	vehicles map[string]*model.Vehicle
}

func NewVehicleRepositoryImpl() ports.VehicleRepository {
	return &MemoryRepository{
		vehicles: make(map[string]*model.Vehicle),
	}
}

func (r *MemoryRepository) Save(ctx context.Context, vehicle *model.Vehicle) (*model.Vehicle, error) {
	r.vehicles[vehicle.Id.String()] = vehicle
	return vehicle, nil
}
