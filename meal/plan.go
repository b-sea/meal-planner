package meal

import (
	"time"

	"github.com/b-sea/meal-planner/dash"
	"github.com/b-sea/meal-planner/food"
)

const (
	defaultKCalMin = 1950.0
	defaultKCalMax = 2050.0
)

// Plan is a meal plan.
type Plan struct {
	id         ID
	kcalTarget CalorieTarget
	days       map[time.Time][]Meal
}

// NewPlan creates a new meal plan.
func NewPlan(id ID, options ...PlanOption) Plan {
	plan := Plan{
		id: id,
		kcalTarget: CalorieTarget{
			min: defaultKCalMin,
			max: defaultKCalMax,
		},
		days: make(map[time.Time][]Meal),
	}

	for _, option := range options {
		option(&plan)
	}

	return plan
}

type TallyCount struct {
}

// TallyNutrition calculates the total nutritional value of a day.
func (p Plan) TallyNutrition(date time.Time) (food.Nutrition, error) {
	total := food.NewNutrition(0, 0, 0, 0)

	for _, meal := range p.days[date] {
		facts, err := meal.NutritionFacts()
		if err != nil {
			return food.NewNutrition(0, 0, 0, 0), err
		}

		total.Add(facts)
	}

	return total, nil
}

// TallyDASH calculates the meal plan against the DASH diet.
func (p Plan) TallyDASH(diet dash.DASH) ([]dash.TallyCount, error) {
	servings := make([]dash.ServingCount, 0)

	for day := range p.days {
		for _, meal := range p.days[day] {
			for _, ingredient := range meal.ingredients {
				converted, err := ingredient.quantity.Convert(ingredient.item.ServingSize().Unit())
				if err != nil {
					return nil, unitConversionError(err)
				}

				servings = append(servings,
					dash.NewServingCount(converted.Float()/ingredient.item.ServingSize().Float(), &ingredient.item),
				)
			}
		}
	}

	return diet.Tally(servings, len(p.days)), nil
}
