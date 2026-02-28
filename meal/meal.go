package meal

import (
	"github.com/b-sea/meal-planner/food"
	"github.com/bcicen/go-units"
)

type ID string

type Ingredient struct {
	item     food.Food
	quantity units.Value
}

func (i Ingredient) Item() food.Food {
	return i.item
}

func (i Ingredient) Quantity() units.Value {
	return i.quantity
}

type Meal struct {
	id          ID
	name        string
	ingredients []Ingredient
}

func NewMeal(id ID, name string, options ...Option) Meal {
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

func (m Meal) Update(options ...Option) {
	for _, option := range options {
		option(&m)
	}
}

func (m Meal) ID() ID {
	return m.id
}

func (m Meal) Name() string {
	return m.name
}

func (m Meal) Ingredients() []Ingredient {
	return m.ingredients
}

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
