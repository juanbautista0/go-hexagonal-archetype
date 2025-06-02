package libraries

import "github.com/google/uuid"

type Uuid = uuid.UUID

func ValidateUuid(stringId string) error {
	_, err := uuid.Parse(stringId)
	return err
}

func ParseUuid(stringId string) (Uuid, error) {
	return uuid.Parse(stringId)
}
func NewUuid() Uuid {
	return uuid.New()
}
