package entity

import (
	"github.com/juanbautista0/go-hexagonal-archetype/app/domain/value_object"
)

type Vehicle struct {
	Id    value_object.Id
	Brand value_object.Brand
	Model value_object.Model
	Year  value_object.Year
}

func NewVehicleFromPrimitives(id string, brand string, model string, year string) (Vehicle, error) {
	voId, err := value_object.NewId(id)
	if err != nil {
		return Vehicle{}, err
	}

	voBrand, err := value_object.NewBrand(brand)
	if err != nil {
		return Vehicle{}, err
	}

	voModel, err := value_object.NewModel(model)
	if err != nil {
		return Vehicle{}, err
	}

	voYear, err := value_object.NewYear(year)
	if err != nil {
		return Vehicle{}, err
	}

	return Vehicle{
		Id:    voId,
		Brand: voBrand,
		Model: voModel,
		Year:  voYear,
	}, nil
}
