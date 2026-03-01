package meal

// CalorieTarget is the daily target caloric range for a meal plan.
type CalorieTarget struct {
	min float64
	max float64
}

// NewCalorieTarget creates a new daily target caloric range.
func NewCalorieTarget(minimum float64, maximum float64) CalorieTarget {
	return CalorieTarget{
		min: minimum,
		max: maximum,
	}
}

// Min is the daily min target for a meal plan.
func (c CalorieTarget) Min() float64 {
	return c.min
}

// Max is the daily max target for a meal plan.
func (c CalorieTarget) Max() float64 {
	return c.max
}
