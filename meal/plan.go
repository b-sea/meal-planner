package meal

import (
	"time"

	"github.com/b-sea/meal-planner/dash"
	"github.com/b-sea/meal-planner/food"
)

var (
	defaultKCalMin = 1950.0
	defaultKCalMax = 2050.0
)

type CalorieTarget struct {
	min float64
	max float64
}

type Plan struct {
	id         ID
	kcalTarget CalorieTarget
	days       map[time.Time][]Meal
}

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

func (p Plan) AddDay(date time.Time) {
	p.days[date] = make([]Meal, 0)
}

func (p Plan) RemoveDay(date time.Time) {
	delete(p.days, date)
}

func (p Plan) AddMeal(date time.Time, meal Meal) {
	if _, ok := p.days[date]; !ok {
		p.days[date] = make([]Meal, 0)
	}

	p.days[date] = append(p.days[date], meal)
}

func (p Plan) RemoveMeal(date time.Time, id ID) {
	i := 0

	for _, meal := range p.days[date] {
		if meal.id != id {
			continue
		}

		p.days[date][i] = meal
		i++
	}

	p.days[date] = p.days[date][:i]
}

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

func (p Plan) TallyDASH(diet dash.DASH) ([]dash.Count, error) {
	servings := make([]dash.Serving, 0)

	for day := range p.days {
		for _, meal := range p.days[day] {
			for _, ingredient := range meal.ingredients {
				converted, err := ingredient.quantity.Convert(ingredient.item.ServingSize().Unit())
				if err != nil {
					return nil, err
				}

				dash.NewServing(converted.Float()/ingredient.item.ServingSize().Float(), &ingredient.item)
			}
		}
	}

	return diet.Tally(servings, len(p.days)), nil
}
