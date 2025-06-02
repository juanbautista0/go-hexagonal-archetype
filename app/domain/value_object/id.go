package value_object

import (
	"errors"

	"github.com/juanbautista0/go-hexagonal-archetype/app/libraries"
)

type Id struct {
	value libraries.Uuid
}

func NewId(stringId string) (Id, error) {
	id, err := libraries.ParseUuid(stringId)
	if err != nil {
		return Id{}, errors.New("invalid Id")
	}
	return Id{value: id}, nil
}

func (i Id) String() string {
	return i.value.String()
}

func (i Id) Equals(other Id) bool {
	return i.value == other.value
}
