package food

import (
	"github.com/bcicen/go-units"
)

// Nutrition is nutrition facts.
type Nutrition struct {
	kcal         float64
	totalFat     units.Value
	saturatedFat units.Value
	transFat     units.Value
	sodium       units.Value
	totalCarbs   units.Value
	fiber        units.Value
	totalSugars  units.Value
	protein      units.Value
}

// NewNutrition creates new nutrition facts.
func NewNutrition(kcal float64, fat float64, carbs float64, protein float64, options ...NutritionOption) Nutrition {
	nutrition := Nutrition{
		kcal:         0,
		totalFat:     units.NewValue(0, units.Gram),
		saturatedFat: units.NewValue(0, units.Gram),
		transFat:     units.NewValue(0, units.Gram),
		sodium:       units.NewValue(0, units.MilliGram),
		totalCarbs:   units.NewValue(0, units.Gram),
		fiber:        units.NewValue(0, units.Gram),
		totalSugars:  units.NewValue(0, units.Gram),
		protein:      units.NewValue(0, units.Gram),
	}

	options = append(
		[]NutritionOption{WithCalories(kcal), WithTotalFat(fat), WithTotalCarbohydrates(carbs), WithProtein(protein)},
		options...,
	)

	for _, option := range options {
		option(&nutrition)
	}

	return nutrition
}

// Update existing nutrition facts.
func (n Nutrition) Update(options ...NutritionOption) {
	for _, option := range options {
		option(&n)
	}
}

// Scale nutrition facts by a percentage.
func (n Nutrition) Scale(ratio float64) Nutrition {
	return Nutrition{
		kcal:         n.kcal * ratio,
		totalFat:     units.NewValue(n.totalFat.Float()*ratio, n.totalFat.Unit()),
		saturatedFat: units.NewValue(n.saturatedFat.Float()*ratio, n.saturatedFat.Unit()),
		transFat:     units.NewValue(n.transFat.Float()*ratio, n.transFat.Unit()),
		sodium:       units.NewValue(n.sodium.Float()*ratio, n.sodium.Unit()),
		totalCarbs:   units.NewValue(n.totalCarbs.Float()*ratio, n.totalCarbs.Unit()),
		fiber:        units.NewValue(n.fiber.Float()*ratio, n.fiber.Unit()),
		totalSugars:  units.NewValue(n.totalSugars.Float()*ratio, n.totalSugars.Unit()),
		protein:      units.NewValue(n.protein.Float()*ratio, n.protein.Unit()),
	}
}

// Add nutrition facts together.
func (n Nutrition) Add(other Nutrition) Nutrition {
	return Nutrition{
		kcal:         n.kcal + other.kcal,
		totalFat:     units.NewValue(n.totalFat.Float()+other.totalFat.Float(), n.totalFat.Unit()),
		saturatedFat: units.NewValue(n.saturatedFat.Float()+other.saturatedFat.Float(), n.saturatedFat.Unit()),
		transFat:     units.NewValue(n.transFat.Float()+other.transFat.Float(), n.transFat.Unit()),
		sodium:       units.NewValue(n.sodium.Float()+other.sodium.Float(), n.sodium.Unit()),
		totalCarbs:   units.NewValue(n.totalCarbs.Float()+other.totalCarbs.Float(), n.totalCarbs.Unit()),
		fiber:        units.NewValue(n.fiber.Float()+other.fiber.Float(), n.fiber.Unit()),
		totalSugars:  units.NewValue(n.totalSugars.Float()+other.totalSugars.Float(), n.totalSugars.Unit()),
		protein:      units.NewValue(n.protein.Float()+other.protein.Float(), n.protein.Unit()),
	}
}

// Calories are the nutritional calories.
func (n Nutrition) Calories() float64 {
	return n.kcal
}

// TotalFat is the nutritional total fat in grams.
func (n Nutrition) TotalFat() units.Value {
	return n.totalFat
}

// SaturatedFat is the nutritional saturated fat in grams.
func (n Nutrition) SaturatedFat() units.Value {
	return n.saturatedFat
}

// TransFat is the nutritional trans fat in grams.
func (n Nutrition) TransFat() units.Value {
	return n.transFat
}

// Sodium is the nutritional sodium in milligrams.
func (n Nutrition) Sodium() units.Value {
	return n.sodium
}

// TotalCarbohydrates is the nutritional total carbohydrates in grams.
func (n Nutrition) TotalCarbohydrates() units.Value {
	return n.totalCarbs
}

// Fiber is the nutritional fiber in grams.
func (n Nutrition) Fiber() units.Value {
	return n.fiber
}

// TotalSugars is the nutritional total sugar in grams.
func (n Nutrition) TotalSugars() units.Value {
	return n.totalSugars
}

// Protein is the nutritional protein in grams.
func (n Nutrition) Protein() units.Value {
	return n.protein
}
