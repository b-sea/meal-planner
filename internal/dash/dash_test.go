package dash_test

import (
	"testing"

	"github.com/b-sea/meal-planner/internal/dash"
	"github.com/stretchr/testify/assert"
)

var _ dash.Tallier = Tallier{}

type Tallier struct {
	DASHGroupResult dash.Group
}

func (t Tallier) DASHGroup() dash.Group {
	return t.DASHGroupResult
}

func TestDASHTally(t *testing.T) {
	t.Parallel()

	servingCounts := []dash.ServingCount{
		dash.NewServingCount(1, Tallier{DASHGroupResult: dash.VeggieGroup}),
		dash.NewServingCount(3, Tallier{DASHGroupResult: dash.VeggieGroup}),

		dash.NewServingCount(1, Tallier{DASHGroupResult: dash.FruitGroup}),

		dash.NewServingCount(4, Tallier{DASHGroupResult: dash.GrainGroup}),
		dash.NewServingCount(5, Tallier{DASHGroupResult: dash.GrainGroup}),
		dash.NewServingCount(4, Tallier{DASHGroupResult: dash.GrainGroup}),

		dash.NewServingCount(2, Tallier{DASHGroupResult: dash.AnimalProteinGroup}),
		dash.NewServingCount(3, Tallier{DASHGroupResult: dash.AnimalProteinGroup}),

		dash.NewServingCount(.6, Tallier{DASHGroupResult: dash.PlantProteinGroup}),

		dash.NewServingCount(20, Tallier{DASHGroupResult: dash.LipidGroup}),
	}

	result := []dash.TallyCount{
		{Group: dash.VeggieGroup, Min: 4, Max: 5, Actual: 4, Deviation: 0},
		{Group: dash.FruitGroup, Min: 4, Max: 5, Actual: 1, Deviation: -3},
		{Group: dash.GrainGroup, Min: 7, Max: 8, Actual: 13, Deviation: 5},
		{Group: dash.DairyGroup, Min: 2, Max: 3, Actual: 0, Deviation: -2},
		{Group: dash.AnimalProteinGroup, Min: 5, Max: 6, Actual: 5, Deviation: 0},
		{Group: dash.PlantProteinGroup, Min: 0.5714285714285714, Max: 0.7142857142857143, Actual: .6, Deviation: 0},
		{Group: dash.LipidGroup, Min: 2, Max: 3, Actual: 20, Deviation: 17},
	}

	assert.Equal(t, dash.New().Tally(servingCounts, 1), result)

}

func TestRequirement(t *testing.T) {
	t.Parallel()

	requirement := dash.New().Requirements()[0]
	assert.Equal(t, requirement.Group(), dash.VeggieGroup)
	assert.Equal(t, requirement.Min(), 4.0)
	assert.Equal(t, requirement.Max(), 5.0)
	assert.Equal(t, requirement.Days(), 1)
}
