package food

import (
	"github.com/b-sea/meal-planner/dash"
	"github.com/bcicen/go-units"
)

type ID string

var _ dash.Tallier = (*Food)(nil)

type Food struct {
	id          ID
	name        string
	group       dash.Group
	servingSize units.Value
	facts       *Nutrition
}

func New(id ID, name string, group dash.Group, servingSize units.Value, options ...Option) Food {
	food := Food{
		id:          id,
		name:        "",
		group:       dash.Group(-1),
		servingSize: units.Value{},
	}

	options = append([]Option{WithName(name), WithGroup(group), WithServingSize(servingSize)}, options...)

	for _, option := range options {
		option(&food)
	}

	return food
}

func (f *Food) Update(options ...Option) {
	for _, option := range options {
		option(f)
	}
}

func (f *Food) DASHGroup() dash.Group {
	return f.group
}

func (f *Food) ID() ID {
	return f.id
}

func (f *Food) Name() string {
	return f.name
}

func (f *Food) ServingSize() units.Value {
	return f.servingSize
}

func (f *Food) NutritionFacts(quantity units.Value) (*Nutrition, error) {
	if f.facts == nil {
		return nil, nil
	}

	converted, err := quantity.Convert(f.servingSize.Unit())
	if err != nil {
		return nil, err
	}

	scaled := f.facts.Scale(converted.Float() / f.servingSize.Float())

	return &scaled, nil
}
