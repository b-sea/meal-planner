// Package food implements foodstuff.
package food

import (
	"github.com/b-sea/meal-planner/dash"
	"github.com/bcicen/go-units"
	"github.com/google/uuid"
)

var _ dash.Tallier = (*Food)(nil)

// Food is a foodstuff.
type Food struct {
	id          uuid.UUID
	name        string
	group       dash.Group
	servingSize units.Value
	facts       *Nutrition
}

// New creates a new food.
func New(id uuid.UUID, name string, group dash.Group, servingSize units.Value, options ...Option) Food {
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

// Update an existing food.
func (f *Food) Update(options ...Option) {
	for _, option := range options {
		option(f)
	}
}

// DASHGroup returns the DASH diet food group this food belongs to.
func (f *Food) DASHGroup() dash.Group {
	return f.group
}

// ID returns the food id.
func (f *Food) ID() uuid.UUID {
	return f.id
}

// Name returns the food name.
func (f *Food) Name() string {
	return f.name
}

// ServingSize returns the food serving size.
func (f *Food) ServingSize() units.Value {
	return f.servingSize
}

// NutritionFacts return the food nutrition facts.
func (f *Food) NutritionFacts(quantity units.Value) (*Nutrition, error) {
	if f.facts == nil {
		return f.facts, nil
	}

	converted, err := quantity.Convert(f.servingSize.Unit())
	if err != nil {
		return nil, unitConversionError(err)
	}

	scaled := f.facts.Scale(converted.Float() / f.servingSize.Float())

	return &scaled, nil
}
