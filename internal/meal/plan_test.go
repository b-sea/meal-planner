package meal_test

import (
	"testing"
	"time"

	"github.com/b-sea/meal-planner/internal/dash"
	"github.com/b-sea/meal-planner/internal/food"
	"github.com/b-sea/meal-planner/internal/meal"
	"github.com/bcicen/go-units"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPlanTallyNutritionUnder(t *testing.T) {
	t.Parallel()

	day := time.Date(2026, 2, 27, 0, 0, 0, 0, time.UTC)

	plan := meal.NewPlan(
		uuid.MustParse("b774ccf9-bcee-4758-8c76-9709774295a2"),
		meal.WithMeals(
			day,
			[]meal.Meal{
				meal.New(uuid.New(), "cereal",
					meal.WithIngredients([]meal.Ingredient{
						meal.NewIngredient(uuid.New(), units.NewValue(29, units.Gram), cheerios, 0),
						meal.NewIngredient(uuid.New(), units.NewValue(.71, units.Pint), milk, 1),
						meal.NewIngredient(uuid.New(), units.NewValue(100, units.Gram), banana, 2),
					}),
				),
				meal.New(uuid.New(), "glass of milk",
					meal.WithIngredients([]meal.Ingredient{
						meal.NewIngredient(uuid.New(), units.NewValue(1, units.Pint), milk, 1),
					}),
				),
			},
		),
	)

	expected := meal.TallyCount{
		Nutrition: food.NewNutrition(
			551.4421052631578, 20.152631578947368, 65.82842105263158, 33.83263157894737,
			food.WithSaturatedFat(10.26),
			food.WithTransFat(.6839999999999999),
			food.WithSodium(504.1),
			food.WithFiber(3.0526315789473686),
			food.WithTotalSugars(45.986315789473686),
		),
		Target:    meal.NewCalorieTarget(1950, 2050),
		Deviation: -1398.5578947368422,
	}

	result, err := plan.TallyNutrition(day)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestPlanTallyNutritionOver(t *testing.T) {
	t.Parallel()

	day := time.Date(2026, 2, 27, 0, 0, 0, 0, time.UTC)

	plan := meal.NewPlan(
		uuid.MustParse("b774ccf9-bcee-4758-8c76-9709774295a2"),
		meal.WithMeals(
			day,
			[]meal.Meal{
				meal.New(uuid.New(), "cereal",
					meal.WithIngredients([]meal.Ingredient{
						meal.NewIngredient(uuid.New(), units.NewValue(87, units.Gram), cheerios, 0),
						meal.NewIngredient(uuid.New(), units.NewValue(6, units.Pint), milk, 1),
					}),
				),
				meal.New(uuid.New(), "glass of milk",
					meal.WithIngredients([]meal.Ingredient{
						meal.NewIngredient(uuid.New(), units.NewValue(5, units.Pint), milk, 1),
					}),
				),
			},
		),
	)

	expected := meal.TallyCount{
		Nutrition: food.NewNutrition(
			3180.5263157894738, 119.15789473684211, 350.10526315789474, 207.1578947368421,
			food.WithSaturatedFat(66),
			food.WithTransFat(4.4),
			food.WithSodium(2745),
			food.WithFiber(9.157894736842104),
			food.WithTotalSugars(290.57894736842104),
		),
		Target:    meal.NewCalorieTarget(1950, 2050),
		Deviation: 1130.5263157894738,
	}
	expected.Target.Min()
	expected.Target.Max()

	result, err := plan.TallyNutrition(day)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestPlanTallyDASH(t *testing.T) {
	t.Parallel()

	day := time.Date(2026, 2, 27, 0, 0, 0, 0, time.UTC)

	plan := meal.NewPlan(
		uuid.MustParse("b774ccf9-bcee-4758-8c76-9709774295a2"),
		meal.WithMeals(
			day,
			[]meal.Meal{
				meal.New(uuid.New(), "cereal",
					meal.WithIngredients([]meal.Ingredient{
						meal.NewIngredient(uuid.New(), units.NewValue(200, units.Gram), cheerios, 0),
						meal.NewIngredient(uuid.New(), units.NewValue(.71, units.Pint), milk, 1),
						meal.NewIngredient(uuid.New(), units.NewValue(240, units.Gram), banana, 2),
					}),
				),
				meal.New(uuid.New(), "glass of milk",
					meal.WithIngredients([]meal.Ingredient{
						meal.NewIngredient(uuid.New(), units.NewValue(4, units.Pint), milk, 1),
					}),
				),
			},
		),
	)

	expected := []dash.TallyCount{
		{Group: 1, Min: 4, Max: 5, Actual: 0, Deviation: -4},
		{Group: 2, Min: 4, Max: 5, Actual: 4.8, Deviation: 0},
		{Group: 3, Min: 7, Max: 8, Actual: 10.526315789473685, Deviation: 2.526315789473685},
		{Group: 4, Min: 2, Max: 3, Actual: 9.42, Deviation: 6.42},
		{Group: 5, Min: 5, Max: 6, Actual: 0, Deviation: -5},
		{Group: 6, Min: 0.5714285714285714, Max: 0.7142857142857143, Actual: 0, Deviation: -0.5714285714285714},
		{Group: 7, Min: 2, Max: 3, Actual: 0, Deviation: -2},
	}

	result, err := plan.TallyDASH(dash.New())
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
