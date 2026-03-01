package meal

import (
	"github.com/b-sea/meal-planner/internal/food"
	"github.com/bcicen/go-units"
	"github.com/google/uuid"
)

// Ingredient is a component of a meal with a quantity.
type Ingredient struct {
	id       uuid.UUID
	quantity units.Value
	item     food.Food
	order    int
}

// NewIngredient creates a new ingredient.
func NewIngredient(id uuid.UUID, quantity units.Value, item food.Food, order int) Ingredient {
	return Ingredient{
		id:       id,
		quantity: quantity,
		item:     item,
		order:    order,
	}
}

// ID returns the ingredient id.
func (i Ingredient) ID() uuid.UUID {
	return i.id
}

// Item returns the ingredient item.
func (i Ingredient) Item() food.Food {
	return i.item
}

// Quantity returns the ingredient quantity.
func (i Ingredient) Quantity() units.Value {
	return i.quantity
}

// Order returns the ingredient order.
func (i Ingredient) Order() int {
	return i.order
}
