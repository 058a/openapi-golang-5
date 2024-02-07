package item_test

import (
	"testing"

	"openapi/internal/domain/stock/item"

	"github.com/google/uuid"
)

func TestNewItemId(t *testing.T) {
	// When
	value := uuid.New()
	itemId, err := item.NewItemId(value)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if itemId.UUID() != value {
		t.Errorf("expected %s, got %s", value, itemId.UUID())
	}
}

func TestNewItemIdFail(t *testing.T) {
	// When
	_, err := item.NewItemId(uuid.Nil)

	// Then
	if err == nil {
		t.Fatal(err)
	}
}
