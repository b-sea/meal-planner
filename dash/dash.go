// Package dash defines and implements the DASH diet.
package dash

// Tallier defines functions required for something to be recognized as part of the DASH diet.
type Tallier interface {
	DASHGroup() Group
}

// DASH is a "Dietary Approach to Stop Hypertension" diet.
type DASH struct {
	requirements []Requirement
}

// New creates a new DASH diet.
func New() *DASH {
	diet := &DASH{
		requirements: []Requirement{
			{group: VeggieGroup, min: 4, max: 5, days: 1},        //nolint:mnd
			{group: FruitGroup, min: 4, max: 5, days: 1},         //nolint:mnd
			{group: GrainGroup, min: 7, max: 8, days: 1},         //nolint:mnd
			{group: DairyGroup, min: 2, max: 3, days: 1},         //nolint:mnd
			{group: AnimalProteinGroup, min: 5, max: 6, days: 1}, //nolint:mnd
			{group: PlantProteinGroup, min: 4, max: 5, days: 7},  //nolint:mnd
			{group: LipidGroup, min: 2, max: 3, days: 1},         //nolint:mnd
		},
	}

	return diet
}

// Requirements returns the DASH dietary requirements.
func (d DASH) Requirements() []Requirement {
	return d.requirements
}

// ServingCount is the amount of servings of a partictular DASH diet food item.
type ServingCount struct {
	count float64
	item  Tallier
}

// NewServingCount creates a new DASH serving count.
func NewServingCount(count float64, item Tallier) ServingCount {
	return ServingCount{
		count: count,
		item:  item,
	}
}

// TallyCount is the final tally count for a food group.
type TallyCount struct {
	Group     Group
	Min       float64
	Max       float64
	Actual    float64
	Deviation float64
}

// Tally calculates the given serving counts against a number of days and evaluate if the diet has been met.
func (d DASH) Tally(servingCounts []ServingCount, days int) []TallyCount {
	tallies := make(map[Group]float64)

	for _, servingCount := range servingCounts {
		tallies[servingCount.item.DASHGroup()] += servingCount.count
	}

	result := make([]TallyCount, 0)

	for i := range d.requirements {
		requirement := d.requirements[i].Extrapolate(days)
		count := tallies[requirement.group]
		deviation := 0.0

		switch {
		case count < requirement.min:
			deviation = count - requirement.min
		case count > requirement.max:
			deviation = count - requirement.max
		default:
		}

		result = append(
			result,
			TallyCount{
				Group:     requirement.group,
				Min:       requirement.min,
				Max:       requirement.max,
				Actual:    count,
				Deviation: deviation,
			},
		)
	}

	return result
}
