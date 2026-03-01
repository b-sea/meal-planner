// Package meal defines and implements meal planning.
package meal

import (
	"github.com/b-sea/meal-planner/internal/food"
	"github.com/google/uuid"
)

// Meal is a planned meal.
type Meal struct {
	id          uuid.UUID
	name        string
	ingredients map[uuid.UUID]Ingredient
}

// New creates a new meal.
func New(id uuid.UUID, name string, options ...Option) Meal {
	meal := Meal{
		id:          id,
		name:        "",
		ingredients: make(map[uuid.UUID]Ingredient),
	}

	options = append([]Option{WithName(name)}, options...)

	for _, option := range options {
		option(&meal)
	}

	return meal
}

// Update an existing meal.
func (m *Meal) Update(options ...Option) {
	for _, option := range options {
		option(m)
	}
}

// ID returns the meal id.
func (m *Meal) ID() uuid.UUID {
	return m.id
}

// Name returns the meal name.
func (m *Meal) Name() string {
	return m.name
}

// Ingredients returns the meal ingredients.
func (m *Meal) Ingredients() []Ingredient {
	ingredients := make([]Ingredient, 0)

	for _, ingredient := range m.ingredients {
		ingredients = append(ingredients, ingredient)
	}

	return ingredients
}

// NutritionFacts calculates the total nutritional values for a meal.
func (m *Meal) NutritionFacts() (food.Nutrition, error) {
	total := food.NewNutrition(0, 0, 0, 0)

	for _, ingredient := range m.ingredients {
		facts, err := ingredient.item.NutritionFacts(ingredient.quantity)
		if err != nil {
			return food.NewNutrition(0, 0, 0, 0), err
		}

		if facts == nil {
			continue
		}

		total = total.Add(*facts)
	}

	return total, nil
}
