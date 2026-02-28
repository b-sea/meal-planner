package food

import (
	"github.com/b-sea/meal-planner/dash"
	"github.com/bcicen/go-units"
)

type Option func(f *Food)

func WithName(name string) Option {
	return func(f *Food) {
		f.name = name
	}
}

func WithGroup(group dash.Group) Option {
	return func(f *Food) {
		f.group = group
	}
}

func WithServingSize(value units.Value) Option {
	return func(f *Food) {
		f.servingSize = value
	}
}

func WithNutritionFacts(facts *Nutrition) Option {
	return func(f *Food) {
		f.facts = facts
	}
}

type NutritionOption func(n *Nutrition)

func WithCalories(kcal float64) NutritionOption {
	return func(n *Nutrition) {
		n.kcal = kcal
	}
}

func WithTotalFat(grams float64) NutritionOption {
	return func(n *Nutrition) {
		n.totalFat = units.NewValue(grams, units.Gram)
	}
}

func WithSaturatedFat(grams float64) NutritionOption {
	return func(n *Nutrition) {
		n.saturatedFat = units.NewValue(grams, units.Gram)
	}
}

func WithTransFat(grams float64) NutritionOption {
	return func(n *Nutrition) {
		n.totalFat = units.NewValue(grams, units.Gram)
	}
}

func WithSodium(milligrams float64) NutritionOption {
	return func(n *Nutrition) {
		n.sodium = units.NewValue(milligrams, units.MilliGram)
	}
}

func WithTotalCarbohydrates(grams float64) NutritionOption {
	return func(n *Nutrition) {
		n.totalCarbs = units.NewValue(grams, units.Gram)
	}
}

func WithFiber(grams float64) NutritionOption {
	return func(n *Nutrition) {
		n.fiber = units.NewValue(grams, units.Gram)
	}
}

func WithTotalSugars(grams float64) NutritionOption {
	return func(n *Nutrition) {
		n.totalSugars = units.NewValue(grams, units.Gram)
	}
}

func WithProtein(grams float64) NutritionOption {
	return func(n *Nutrition) {
		n.protein = units.NewValue(grams, units.Gram)
	}
}
