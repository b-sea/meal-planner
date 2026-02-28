package dash_test

import (
	"fmt"
	"testing"

	"github.com/b-sea/meal-planner/dash"
	"github.com/b-sea/meal-planner/food"
	"github.com/bcicen/go-units"
	"github.com/stretchr/testify/assert"
)

var apple food.Food

func init() {
	appleNutrition := food.NewNutrition(95, .3, 25.1, .473, food.WithFiber(4.37), food.WithTotalSugars(18.9))
	apple = food.New(
		food.ID("1"), "apple", dash.FruitGroup, units.NewValue(.5, units.Pint),
		food.WithNutritionFacts(&appleNutrition),
	)
}

func TestHMM(t *testing.T) {
	diet := dash.New()
	result := diet.Tally([]dash.Serving{
		dash.NewServing(1.5, &apple), dash.NewServing(30, &apple),
	}, 7)
	fmt.Printf("%+v", result)

}

func TestUHH(t *testing.T) {
	data, err := apple.NutritionFacts(units.NewValue(1, units.Quart))
	assert.NoError(t, err)
	fmt.Printf("%+v", data)
}
