package meal

import (
	"time"

	"github.com/b-sea/meal-planner/food"
)

type Plan struct {
	id   ID
	days map[time.Time][]Meal
}

func NewPlan(id ID, options ...PlanOption) Plan {
	plan := Plan{
		id:   id,
		days: make(map[time.Time][]Meal),
	}

	for _, option := range options {
		option(plan)
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

func (p Plan) NutritionFacts(date time.Time) (food.Nutrition, error) {
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
