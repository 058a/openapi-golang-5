package item

import (
	"fmt"

	"github.com/google/uuid"
)

type itemId struct {
	value uuid.UUID
}

func NewItemId(value uuid.UUID) (itemId, error) {
	if value == uuid.Nil {
		return itemId{}, fmt.Errorf("nil item id")
	}
	return itemId{value}, nil
}

func (e itemId) UUID() uuid.UUID {
	return e.value
}
