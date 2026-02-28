package dash

type Group int

const (
	VeggieGroup Group = iota
	FruitGroup
	GrainGroup
	DairyGroup
	AnimalProteinGroup
	PlantProteinGroup
	LipidGroup
)

func (g Group) String() string {
	switch g {
	case VeggieGroup:
		return "vegetables"
	case FruitGroup:
		return "fruit"
	case GrainGroup:
		return "grains"
	case DairyGroup:
		return "dairy"
	case AnimalProteinGroup:
		return "animal protein"
	case PlantProteinGroup:
		return "plant protein"
	case LipidGroup:
		return "lipids"
	default:
		return "unknown"
	}
}

type Tallier interface {
	DASHGroup() Group
}

type Requirement struct {
	group Group
	min   float64
	max   float64
	days  int
}

func (r Requirement) Extrapolate(days int) Requirement {
	return Requirement{
		group: r.group,
		min:   r.min / float64(r.days) * float64(days),
		max:   r.max / float64(r.days) * float64(days),
		days:  days,
	}
}

type DASH struct {
	requirements []Requirement
}

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

type Serving struct {
	amount float64
	item   Tallier
}

func NewServing(amount float64, item Tallier) Serving {
	return Serving{
		amount: amount,
		item:   item,
	}
}

type Count struct {
	Group     Group
	Min       float64
	Max       float64
	Total     float64
	Deviation float64
}

func (d *DASH) Tally(servings []Serving, days int) []Count {
	tallies := make(map[Group]float64)

	for _, serving := range servings {
		tallies[serving.item.DASHGroup()] += serving.amount
	}

	result := make([]Count, 0)

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
			Count{
				Group:     requirement.group,
				Min:       requirement.min,
				Max:       requirement.max,
				Total:     count,
				Deviation: deviation,
			},
		)
	}

	return result
}
