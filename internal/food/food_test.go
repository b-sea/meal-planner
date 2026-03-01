package food_test

import (
	"testing"

	"github.com/b-sea/meal-planner/internal/dash"
	"github.com/b-sea/meal-planner/internal/food"
	"github.com/bcicen/go-units"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFoodNew(t *testing.T) {
	t.Parallel()

	apple := food.New(
		uuid.MustParse("79207b22-0b60-49a1-a9ae-0d683676bdce"),
		"apple",
		dash.FruitGroup,
		units.NewValue(.25, units.Pint),
	)

	assert.Equal(t, uuid.MustParse("79207b22-0b60-49a1-a9ae-0d683676bdce"), apple.ID())
	assert.Equal(t, "apple", apple.Name())
	assert.Equal(t, dash.FruitGroup, apple.DASHGroup())
	assert.Equal(t, units.NewValue(.25, units.Pint), apple.ServingSize())
}

func TestFoodNutritionFacts(t *testing.T) {
	t.Parallel()

	apple := food.New(
		uuid.MustParse("79207b22-0b60-49a1-a9ae-0d683676bdce"),
		"apple",
		dash.FruitGroup,
		units.NewValue(.25, units.Pint),
		food.WithNutritionFacts(
			food.NewNutrition(
				134, 30, 10.2, 5,
				food.WithSaturatedFat(.2),
				food.WithTransFat(2),
				food.WithSodium(45),
				food.WithFiber(11.2),
				food.WithTotalSugars(.1),
			),
		),
	)

	result, err := apple.NutritionFacts(units.NewValue(1, units.Gallon))
	assert.NoError(t, err)
	assert.Equal(t, 4288.0, result.Calories())
	assert.Equal(t, units.NewValue(960, units.Gram), result.TotalFat())
	assert.Equal(t, units.NewValue(6.4, units.Gram), result.SaturatedFat())
	assert.Equal(t, units.NewValue(64, units.Gram), result.TransFat())
	assert.Equal(t, units.NewValue(1440, units.MilliGram), result.Sodium())
	assert.Equal(t, units.NewValue(326.4, units.Gram), result.TotalCarbohydrates())
	assert.Equal(t, units.NewValue(358.4, units.Gram), result.Fiber())
	assert.Equal(t, units.NewValue(3.2, units.Gram), result.TotalSugars())
	assert.Equal(t, units.NewValue(160, units.Gram), result.Protein())

	result, err = apple.NutritionFacts(units.NewValue(1, units.Celsius))
	assert.ErrorIs(t, err, food.ErrUnitConversion)

	apple.Update(food.ClearNutritionFacts())
	result, err = apple.NutritionFacts(units.NewValue(1, units.Celsius))
	assert.NoError(t, err)
	assert.Nil(t, result)
}
