package dash

// Group is a DASH food group.
type Group int

// VeggieGroup, et al. are the various DASH food groups.
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
