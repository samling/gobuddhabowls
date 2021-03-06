package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type InventoryItem struct {
	ID                   uuid.UUID             `json:"id" db:"id"`
	CreatedAt            time.Time             `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time             `json:"updated_at" db:"updated_at"`
	Name                 string                `json:"name" db:"name"`
	Category             InventoryItemCategory `belongs_to:"inventory_item_categories" db:"-"`
	CategoryID           uuid.UUID             `json:"inventory_item_category_id" db:"inventory_item_category_id"`
	CountUnit            string                `json:"count_unit" db:"count_unit"`
	RecipeUnit           string                `json:"recipe_unit" db:"recipe_unit"`
	RecipeUnitConversion float64               `json:"recipe_unit_conversion" db:"recipe_unit_conversion"`
	Yield                float64               `json:"yield" db:"yield"`
	Index                int                   `json:"index" db:"index"`
	IsActive             bool                  `json:"is_active" db:"is_active"`
}

// String is not required by pop and may be deleted
func (i InventoryItem) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// InventoryItems is not required by pop and may be deleted
type InventoryItems []InventoryItem

// String is not required by pop and may be deleted
func (i InventoryItems) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (i *InventoryItem) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: i.Name, Name: "Name"},
		&validators.StringIsPresent{Field: i.CountUnit, Name: "CountUnit"},
		&validators.StringIsPresent{Field: i.RecipeUnit, Name: "RecipeUnit"},
		&validators.IntIsPresent{Field: i.Index, Name: "Index"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (i *InventoryItem) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (i *InventoryItem) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (i InventoryItem) GetID() uuid.UUID {
	return i.ID
}

func (i InventoryItem) GetName() string {
	return i.Name
}

func (i InventoryItem) GetCategory() InventoryItemCategory {
	return i.Category
}

func (i InventoryItem) GetIndex() int {
	return i.Index
}

// GetSortValue returns a value for sorting where Category is highest prcedence
// and item index is second
func (i InventoryItem) GetSortValue() int {
	return i.Category.Index*1000 + i.Index
}
