package dash

// Requirement defines a dietary requirement for a food group.
type Requirement struct {
	group Group
	min   float64
	max   float64
	days  int
}

// Group returns the dietary requirement food group.
func (r Requirement) Group() Group {
	return r.group
}

// Min returns the dietary requirement min servings.
func (r Requirement) Min() float64 {
	return r.min
}

// Max returns the dietary requirement max servings.
func (r Requirement) Max() float64 {
	return r.max
}

// Days returns the dietary requirement number of days.
func (r Requirement) Days() int {
	return r.days
}

// Extrapolate recalculates a food group requirement for a different amount of days.
func (r Requirement) Extrapolate(days int) Requirement {
	return Requirement{
		group: r.group,
		min:   r.min / float64(r.days) * float64(days),
		max:   r.max / float64(r.days) * float64(days),
		days:  days,
	}
}
