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

func TestNewIngredient(t *testing.T) {
	milk := food.New(
		uuid.MustParse("79207b22-0b60-49a1-a9ae-0d683676bdce"),
		"milk",
		dash.DairyGroup,
		units.NewValue(.5, units.Pint),
	)

	ingredient := meal.NewIngredient(
		uuid.MustParse("8e45f3d8-02fa-4014-8a94-45e9a88145f3"),
		units.NewValue(.71, units.Pint),
		milk,
		3,
	)

	assert.Equal(t, uuid.MustParse("8e45f3d8-02fa-4014-8a94-45e9a88145f3"), ingredient.ID())
	assert.Equal(t, milk, ingredient.Item())
	assert.Equal(t, units.NewValue(.71, units.Pint), ingredient.Quantity())
	assert.Equal(t, 3, ingredient.Order())
}
