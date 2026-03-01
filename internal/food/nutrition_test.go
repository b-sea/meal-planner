package food_test

import (
	"testing"

	"github.com/b-sea/meal-planner/internal/food"
	"github.com/bcicen/go-units"
	"github.com/stretchr/testify/assert"
)

func TestNutritionScale(t *testing.T) {
	t.Parallel()

	nutrition := food.NewNutrition(
		134, 30, 10.2, 5,
		food.WithSaturatedFat(.2),
		food.WithTransFat(2),
		food.WithSodium(45),
		food.WithFiber(11.2),
		food.WithTotalSugars(.1),
	)

	result := nutrition.Scale(2.25)
	assert.Equal(t, 301.5, result.Calories())
	assert.Equal(t, units.NewValue(67.5, units.Gram), result.TotalFat())
	assert.Equal(t, units.NewValue(.45, units.Gram), result.SaturatedFat())
	assert.Equal(t, units.NewValue(4.5, units.Gram), result.TransFat())
	assert.Equal(t, units.NewValue(101.25, units.MilliGram), result.Sodium())
	assert.Equal(t, units.NewValue(22.95, units.Gram), result.TotalCarbohydrates())
	assert.Equal(t, units.NewValue(25.2, units.Gram), result.Fiber())
	assert.Equal(t, units.NewValue(.225, units.Gram), result.TotalSugars())
	assert.Equal(t, units.NewValue(11.25, units.Gram), result.Protein())
}

func TestNutritionAdd(t *testing.T) {
	t.Parallel()

	first := food.NewNutrition(
		5, .5, 2, 1.4,
		food.WithSaturatedFat(.01),
		food.WithTransFat(.3),
		food.WithSodium(13),
		food.WithFiber(8.7),
		food.WithTotalSugars(2),
	)

	result := first.Add(
		food.NewNutrition(
			2, 1, 0, 2.4,
			food.WithSaturatedFat(.24),
			food.WithTransFat(1.2),
			food.WithSodium(21),
			food.WithFiber(7.5),
			food.WithTotalSugars(1.2),
		),
	)

	assert.Equal(t, 7.0, result.Calories())
	assert.Equal(t, units.NewValue(1.5, units.Gram), result.TotalFat())
	assert.Equal(t, units.NewValue(.25, units.Gram), result.SaturatedFat())
	assert.Equal(t, units.NewValue(1.5, units.Gram), result.TransFat())
	assert.Equal(t, units.NewValue(34, units.MilliGram), result.Sodium())
	assert.Equal(t, units.NewValue(2, units.Gram), result.TotalCarbohydrates())
	assert.Equal(t, units.NewValue(16.2, units.Gram), result.Fiber())
	assert.Equal(t, units.NewValue(3.2, units.Gram), result.TotalSugars())
	assert.Equal(t, units.NewValue(3.8, units.Gram), result.Protein())
}
