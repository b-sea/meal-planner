package meal_test

import (
	"testing"

	"github.com/b-sea/meal-planner/internal/dash"
	"github.com/b-sea/meal-planner/internal/food"
	"github.com/b-sea/meal-planner/internal/meal"
	"github.com/bcicen/go-units"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	cheerios = food.New(
		uuid.MustParse("dd7b49de-a9d0-46b4-be05-400c59f348af"),
		"cheerios",
		dash.GrainGroup,
		units.NewValue(19, units.Gram),
		food.WithNutritionFacts(
			food.NewNutrition(
				70, 2, 14, 2,
				food.WithSodium(95),
				food.WithFiber(2),
				food.WithTotalSugars(1),
			),
		),
	)

	milk = food.New(
		uuid.MustParse("dd7b49de-a9d0-46b4-be05-400c59f348af"),
		"milk",
		dash.DairyGroup,
		units.NewValue(.5, units.Pint),
		food.WithNutritionFacts(
			food.NewNutrition(
				130, 5, 13, 9,
				food.WithSaturatedFat(3),
				food.WithTransFat(.2),
				food.WithSodium(105),
				food.WithTotalSugars(13),
			),
		),
	)

	banana = food.New(
		uuid.MustParse("dd7b49de-a9d0-46b4-be05-400c59f348af"),
		"banana",
		dash.FruitGroup,
		units.NewValue(50, units.Gram),
	)
)

func TestNewMeal(t *testing.T) {
	milkGlass := meal.New(
		uuid.MustParse("b7a0099f-8716-404a-8df6-c8aa3457a85b"), "glass of milk",
		meal.WithIngredients([]meal.Ingredient{
			meal.NewIngredient(uuid.MustParse("8e45f3d8-02fa-4014-8a94-45e9a88145f3"), units.NewValue(.71, units.Pint), milk, 0),
		}),
	)
	assert.Equal(t, uuid.MustParse("b7a0099f-8716-404a-8df6-c8aa3457a85b"), milkGlass.ID())
	assert.Equal(t, "glass of milk", milkGlass.Name())
	assert.Equal(t,
		[]meal.Ingredient{meal.NewIngredient(uuid.MustParse("8e45f3d8-02fa-4014-8a94-45e9a88145f3"), units.NewValue(.71, units.Pint), milk, 0)},
		milkGlass.Ingredients(),
	)

	milkGlass.Update(meal.WithName("updated name"))
	assert.Equal(t, "updated name", milkGlass.Name())
}

func TestMealNutritionFacts(t *testing.T) {
	cereal := meal.New(
		uuid.MustParse("b7a0099f-8716-404a-8df6-c8aa3457a85b"), "big bowl of cereal",
		meal.WithIngredients([]meal.Ingredient{
			meal.NewIngredient(uuid.MustParse("b78e1add-9120-41ef-ba8c-aea6ad2718ef"), units.NewValue(29, units.Gram), cheerios, 0),
			meal.NewIngredient(uuid.MustParse("21debb36-ae05-4f80-b3e1-cefbcde52034"), units.NewValue(.71, units.Pint), milk, 1),
			meal.NewIngredient(uuid.MustParse("54a12544-c21f-4f63-90a1-6bf1d5a7bfaf"), units.NewValue(100, units.Gram), banana, 2),
		}),
	)

	result, err := cereal.NutritionFacts()
	assert.NoError(t, err)
	assert.Equal(t, 291.4421052631579, result.Calories())
	assert.Equal(t, units.NewValue(10.152631578947368, units.Gram), result.TotalFat())
	assert.Equal(t, units.NewValue(4.26, units.Gram), result.SaturatedFat())
	assert.Equal(t, units.NewValue(.284, units.Gram), result.TransFat())
	assert.Equal(t, units.NewValue(294.1, units.MilliGram), result.Sodium())
	assert.Equal(t, units.NewValue(39.82842105263158, units.Gram), result.TotalCarbohydrates())
	assert.Equal(t, units.NewValue(3.0526315789473686, units.Gram), result.Fiber())
	assert.Equal(t, units.NewValue(19.986315789473686, units.Gram), result.TotalSugars())
	assert.Equal(t, units.NewValue(15.832631578947368, units.Gram), result.Protein())
}
