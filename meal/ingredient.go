package meal

import (
	"github.com/b-sea/meal-planner/food"
	"github.com/bcicen/go-units"
)

// Ingredient is a component of a meal with a quantity.
type Ingredient struct {
	id       ID
	quantity units.Value
	item     food.Food
}

// NewIngredient creates a new ingredient.
func NewIngredient(id ID, quantity units.Value, item food.Food) Ingredient {
	return Ingredient{
		id:       id,
		quantity: quantity,
		item:     item,
	}
}

// ID returns the ingredient id.
func (i Ingredient) ID() ID {
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
