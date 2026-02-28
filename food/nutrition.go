package food

import (
	"github.com/bcicen/go-units"
)

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

func (n Nutrition) Update(options ...NutritionOption) {
	for _, option := range options {
		option(&n)
	}
}

func (n Nutrition) Scale(ratio float64) Nutrition {
	return Nutrition{
		kcal:         n.kcal * ratio,
		totalFat:     units.NewValue(n.totalFat.Float()+ratio, n.totalFat.Unit()),
		saturatedFat: units.NewValue(n.saturatedFat.Float()+ratio, n.saturatedFat.Unit()),
		transFat:     units.NewValue(n.transFat.Float()+ratio, n.transFat.Unit()),
		sodium:       units.NewValue(n.transFat.Float()+ratio, n.sodium.Unit()),
		totalCarbs:   units.NewValue(n.totalCarbs.Float()+ratio, n.totalCarbs.Unit()),
		fiber:        units.NewValue(n.fiber.Float()+ratio, n.fiber.Unit()),
		totalSugars:  units.NewValue(n.totalSugars.Float()+ratio, n.totalSugars.Unit()),
		protein:      units.NewValue(n.protein.Float()+ratio, n.protein.Unit()),
	}
}

func (n Nutrition) Add(other Nutrition) Nutrition {
	return Nutrition{
		kcal:         n.kcal + other.kcal,
		totalFat:     units.NewValue(n.totalFat.Float()+other.totalFat.Float(), n.totalFat.Unit()),
		saturatedFat: units.NewValue(n.saturatedFat.Float()+other.saturatedFat.Float(), n.saturatedFat.Unit()),
		transFat:     units.NewValue(n.transFat.Float()+other.transFat.Float(), n.transFat.Unit()),
		sodium:       units.NewValue(n.transFat.Float()+other.transFat.Float(), n.sodium.Unit()),
		totalCarbs:   units.NewValue(n.totalCarbs.Float()+other.totalCarbs.Float(), n.totalCarbs.Unit()),
		fiber:        units.NewValue(n.fiber.Float()+other.fiber.Float(), n.fiber.Unit()),
		totalSugars:  units.NewValue(n.totalSugars.Float()+other.totalSugars.Float(), n.totalSugars.Unit()),
		protein:      units.NewValue(n.protein.Float()+other.protein.Float(), n.protein.Unit()),
	}
}

func (n Nutrition) Calories() float64 {
	return n.kcal
}

func (n Nutrition) TotalFat() units.Value {
	return n.totalFat
}

func (n Nutrition) SaturatedFat() units.Value {
	return n.saturatedFat
}

func (n Nutrition) TransFat() units.Value {
	return n.transFat
}

func (n Nutrition) Sodium() units.Value {
	return n.sodium
}

func (n Nutrition) TotalCarbs() units.Value {
	return n.totalCarbs
}

func (n Nutrition) Fiber() units.Value {
	return n.fiber
}

func (n Nutrition) TotalSugars() units.Value {
	return n.totalSugars
}

func (n Nutrition) Protein() units.Value {
	return n.protein
}
