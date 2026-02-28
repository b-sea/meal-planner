package meal

import (
	"time"
)

// Option is a create or update option for meals.
type Option func(m *Meal)

// WithName sets the meal name.
func WithName(name string) Option {
	return func(m *Meal) {
		m.name = name
	}
}

// WithIngredients sets the meal ingredients.
func WithIngredients(ingredients []Ingredient) Option {
	return func(m *Meal) {
		m.ingredients = ingredients
	}
}

// PlanOption is a create or update option for meal plans.
type PlanOption func(p *Plan)

// WithMeals sets the meals for a given day in the meal plan.
func WithMeals(date time.Time, meals []Meal) PlanOption {
	return func(p *Plan) {
		p.days[date] = meals
	}
}

// WithCalorieTarget sets the calorie target for the meal plan.
func WithCalorieTarget(minimum float64, maximum float64) PlanOption {
	return func(p *Plan) {
		p.kcalTarget = CalorieTarget{
			min: minimum,
			max: maximum,
		}
	}
}
