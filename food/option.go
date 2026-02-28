package food

import (
	"github.com/b-sea/meal-planner/dash"
	"github.com/bcicen/go-units"
)

// Option is a create or update option for food.
type Option func(f *Food)

// WithName sets the food name.
func WithName(name string) Option {
	return func(f *Food) {
		f.name = name
	}
}

// WithGroup sets the DASH food group.
func WithGroup(group dash.Group) Option {
	return func(f *Food) {
		f.group = group
	}
}

// WithServingSize sets the food serving side.
func WithServingSize(value units.Value) Option {
	return func(f *Food) {
		f.servingSize = value
	}
}

// WithNutritionFacts sets the food nutrition facts.
func WithNutritionFacts(facts *Nutrition) Option {
	return func(f *Food) {
		f.facts = facts
	}
}

// NutritionOption is a create or update option for nutrition facts.
type NutritionOption func(n *Nutrition)

// WithCalories sets the nutritional calories.
func WithCalories(kcal float64) NutritionOption {
	return func(n *Nutrition) {
		n.kcal = kcal
	}
}

// WithTotalFat sets the nutritional total fat in grams.
func WithTotalFat(grams float64) NutritionOption {
	return func(n *Nutrition) {
		n.totalFat = units.NewValue(grams, units.Gram)
	}
}

// WithSaturatedFat sets the nutritional saturated fat in grams.
func WithSaturatedFat(grams float64) NutritionOption {
	return func(n *Nutrition) {
		n.saturatedFat = units.NewValue(grams, units.Gram)
	}
}

// WithTransFat sets the nutritional trans fat in grams.
func WithTransFat(grams float64) NutritionOption {
	return func(n *Nutrition) {
		n.totalFat = units.NewValue(grams, units.Gram)
	}
}

// WithSodium sets the nutritional sodium in milligrams.
func WithSodium(milligrams float64) NutritionOption {
	return func(n *Nutrition) {
		n.sodium = units.NewValue(milligrams, units.MilliGram)
	}
}

// WithTotalCarbohydrates sets the nutritional total carbohydrates in grams.
func WithTotalCarbohydrates(grams float64) NutritionOption {
	return func(n *Nutrition) {
		n.totalCarbs = units.NewValue(grams, units.Gram)
	}
}

// WithFiber sets the nutritional fiber in grams.
func WithFiber(grams float64) NutritionOption {
	return func(n *Nutrition) {
		n.fiber = units.NewValue(grams, units.Gram)
	}
}

// WithTotalSugars sets the nutritional total sugar in grams.
func WithTotalSugars(grams float64) NutritionOption {
	return func(n *Nutrition) {
		n.totalSugars = units.NewValue(grams, units.Gram)
	}
}

// WithProtein sets the nutritional protein in grams.
func WithProtein(grams float64) NutritionOption {
	return func(n *Nutrition) {
		n.protein = units.NewValue(grams, units.Gram)
	}
}
