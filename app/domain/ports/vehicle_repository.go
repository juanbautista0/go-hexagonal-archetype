package ports

import (
	"context"

	"github.com/juanbautista0/go-hexagonal-archetype/app/domain/entity"
)

type VehicleRepository interface {
	Save(ctx context.Context, vehicle *entity.Vehicle) (*entity.Vehicle, error)
	GetByID(ctx context.Context, id string) (*entity.Vehicle, error)
	UpdateByID(ctx context.Context, id string, updated *entity.Vehicle) (*entity.Vehicle, error)
}
