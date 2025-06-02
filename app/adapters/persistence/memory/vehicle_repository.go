package memory

import (
	"context"
	"errors"
	"sync"

	"github.com/juanbautista0/go-hexagonal-archetype/app/domain/entity"
	"github.com/juanbautista0/go-hexagonal-archetype/app/domain/ports"
)

type InMemoryVehicleRepository struct {
	mu       sync.RWMutex
	vehicles map[string]*entity.Vehicle
}

func NewInMemoryVehicleRepository() ports.VehicleRepository {
	return &InMemoryVehicleRepository{
		vehicles: make(map[string]*entity.Vehicle),
	}
}

func (r *InMemoryVehicleRepository) Save(ctx context.Context, vehicle *entity.Vehicle) (*entity.Vehicle, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.vehicles[vehicle.Id.String()] = vehicle
	return vehicle, nil
}

func (r *InMemoryVehicleRepository) GetByID(ctx context.Context, id string) (*entity.Vehicle, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	vehicle, ok := r.vehicles[id]
	if !ok {
		return nil, errors.New("vehicle not found")
	}
	return vehicle, nil
}

func (r *InMemoryVehicleRepository) UpdateByID(ctx context.Context, id string, updated *entity.Vehicle) (*entity.Vehicle, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.vehicles[id]
	if !ok {
		return nil, errors.New("vehicle not found")
	}

	r.vehicles[id] = updated
	return updated, nil
}
