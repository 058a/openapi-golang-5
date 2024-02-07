package item_test

import (
	"openapi/internal/domain/stock/item"
	"testing"

	"github.com/google/uuid"
)

func TestNewAggregate(t *testing.T) {
	// When
	itemId, err := item.NewItemId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}
	itemName, err := item.NewItemName("test")
	if err != nil {
		t.Fatal(err)
	}
	a, err := item.NewAggregate(itemId, itemName)
	if err != nil {
		t.Fatal(err)
	}

	// Then
	if a.GetId() != itemId {
		t.Errorf("expected %v, got %v", itemId, a.GetId())
	}
	if a.Name != itemName {
		t.Errorf("expected %s, got %s", itemName, a.Name)
	}
	if a.IsDeleted() != false {
		t.Errorf("expected %t, got %t", false, a.IsDeleted())
	}
}

func TestChangName(t *testing.T) {
	// When
	itemId, err := item.NewItemId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}
	itemName, err := item.NewItemName("test")
	if err != nil {
		t.Fatal(err)
	}
	a, err := item.NewAggregate(itemId, itemName)
	if err != nil {
		t.Fatal(err)
	}

	// When
	newName, err := item.NewItemName("test2")
	if err != nil {
		t.Fatal(err)
	}
	a.Name = newName

	// Then
	if a.Name != newName {
		t.Errorf("expected %s, got %s", newName, a.Name)
	}
}

func TestDelete(t *testing.T) {
	// When
	itemId, err := item.NewItemId(uuid.New())
	if err != nil {
		t.Fatal(err)
	}

	itemName, err := item.NewItemName("test")
	if err != nil {
		t.Fatal(err)
	}

	a, err := item.NewAggregate(itemId, itemName)
	if err != nil {
		t.Fatal(err)
	}

	// When
	a.Delete()

	// Then
	if !a.IsDeleted() {
		t.Errorf("expected %t, got %t", true, a.IsDeleted())
	}
}
