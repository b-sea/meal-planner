// Package meal defines and implements meal planning.
package meal

import (
	"github.com/b-sea/meal-planner/food"
)

// ID is a unique identifier for meal planning related things.
type ID string

// Meal is a planned meal.
type Meal struct {
	id          ID
	name        string
	ingredients []Ingredient
}

// New creates a new meal.
func New(id ID, name string, options ...Option) Meal {
	meal := Meal{
		id:          id,
		name:        "",
		ingredients: make([]Ingredient, 0),
	}

	options = append([]Option{WithName(name)}, options...)

	for _, option := range options {
		option(&meal)
	}

	return meal
}

// Update an existing meal.
func (m Meal) Update(options ...Option) {
	for _, option := range options {
		option(&m)
	}
}

// ID returns the meal id.
func (m Meal) ID() ID {
	return m.id
}

// Name returns the meal name.
func (m Meal) Name() string {
	return m.name
}

// Ingredients returns the meal ingredients.
func (m Meal) Ingredients() []Ingredient {
	return m.ingredients
}

// NutritionFacts calculates the total nutritional values for a meal.
func (m Meal) NutritionFacts() (food.Nutrition, error) {
	total := food.NewNutrition(0, 0, 0, 0)

	for _, ingredient := range m.ingredients {
		facts, err := ingredient.item.NutritionFacts(ingredient.quantity)
		if err != nil {
			return food.NewNutrition(0, 0, 0, 0), err
		}

		if facts == nil {
			continue
		}

		total.Add(*facts)
	}

	return total, nil
}
