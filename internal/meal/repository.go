package meal

import "github.com/google/uuid"

// Repository defines all functions required for interacting with persistent meal data.
type Repository interface {
	FindMeal() ([]*Meal, error)
	GetMealByID(id uuid.UUID) (*Meal, error)
	GetMealsByIDs(ids []uuid.UUID) ([]*Meal, error)
	CreateMeal(meal *Meal) error
	UpdateMeal(meal *Meal) error
	DeleteMeal(id uuid.UUID) error
}

// PlanRepository defines all functions required for interacting with persistent meal plan data.
type PlanRepository interface {
	FindPlan() ([]*Plan, error)
	GetPlanByID(id uuid.UUID) (*Plan, error)
	CreatePlan(plan *Plan) error
	UpdatePlan(plan *Plan) error
	DeletePlan(id uuid.UUID) error
}
