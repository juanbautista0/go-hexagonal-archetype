package libraries

import "github.com/google/uuid"

type Uuid = uuid.UUID

func NewUuid() Uuid {
	return uuid.New()
}
