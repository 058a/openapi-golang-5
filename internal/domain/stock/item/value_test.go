package item_test

import (
	"testing"

	"openapi/internal/domain/stock/item"
)

func TestNewItemName(t *testing.T) {
	// Given
	value := "test"

	// When
	a, err := item.NewItemName(value)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if a.String() != value {
		t.Errorf("expected %s, got %s", value, a.String())
	}
}

func TestNewItemNameEmpty(t *testing.T) {
	// Given
	value := ""

	// When
	_, err := item.NewItemName(value)

	// Then
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
