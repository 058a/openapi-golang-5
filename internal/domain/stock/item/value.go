package item

import (
	"fmt"
)

type itemName struct {
	value string
}

func NewItemName(value string) (itemName, error) {
	if value == "" {
		return itemName{}, fmt.Errorf("empty item name")
	}
	return itemName{value}, nil
}

func (v itemName) String() string {
	return v.value
}
