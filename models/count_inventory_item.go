package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

type CountInventoryItem struct {
	ID               uuid.UUID     `json:"id" db:"id"`
	CreatedAt        time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time     `json:"updated_at" db:"updated_at"`
	Count            float64       `json:"count" db:"count"`
	InventoryID      uuid.UUID     `json:"inventory_id" db:"inventory_id"`
	SelectedVendorID uuid.NullUUID `json:"selected_vendor_id" db:"selected_vendor_id"`
	InventoryItemID  uuid.UUID     `json:"inventory_item_id" db:"inventory_item_id"`
	Inventory        Inventory     `belongs_to:"inventories" db:"-"`
	SelectedVendor   Vendor        `belongs_to:"vendors" db:"-"`
	InventoryItem    InventoryItem `belongs_to:"inventory_items" db:"-"`
}

// CountInventoryItems is not required by pop and may be deleted
type CountInventoryItems []CountInventoryItem

// String is not required by pop and may be deleted
func (c CountInventoryItems) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *CountInventoryItem) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *CountInventoryItem) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *CountInventoryItem) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
