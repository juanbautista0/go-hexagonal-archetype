package ports

import (
	"context"

	"github.com/juanbautista0/go-hexagonal-archetype/app/domain/model"
)

type VehicleRepository interface {
	Save(ctx context.Context, vehicle *model.Vehicle) (*model.Vehicle, error)
}
