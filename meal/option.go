package meal

import (
	"time"

	"github.com/b-sea/meal-planner/food"
	"github.com/bcicen/go-units"
)

type Option func(m *Meal)

func WithName(name string) Option {
	return func(m *Meal) {
		m.name = name
	}
}

func WithIngredient(quantity units.Value, item food.Food) Option {
	return func(m *Meal) {
		m.ingredients = append(m.ingredients, Ingredient{quantity: quantity, item: item})
	}
}

type PlanOption func(p Plan)

func WithDay(date time.Time, meals []Meal) PlanOption {
	return func(p Plan) {
		p.days[date] = meals
	}
}
