package meal

type CalorieTarget struct {
	min float64
	max float64
}

func (c CalorieTarget) Min() float64 {
	return c.min
}

func (c CalorieTarget) Max() float64 {
	return c.max
}
