package mock

import (
	"github.com/b-sea/meal-planner/internal/food"
	"github.com/google/uuid"
)

var _ food.Repository = (*Foodstuff)(nil)

type Foodstuff struct {
	FindFoodResult     []*food.Food
	FindFoodErr        error
	GetFoodByIDResult  *food.Food
	GetFoodByIDErr     error
	GetFoodByIDsResult []*food.Food
	GetFoodByIDsErr    error
	CreateFoodErr      error
	UpdateFoodErr      error
	DeleteFoodErr      error
}

func (m *Foodstuff) FindFood() ([]*food.Food, error) {
	return m.FindFoodResult, m.FindFoodErr
}

func (m *Foodstuff) GetFoodByID(id uuid.UUID) (*food.Food, error) {
	return m.GetFoodByIDResult, m.GetFoodByIDErr
}

func (m *Foodstuff) GetFoodByIDs(ids []uuid.UUID) ([]*food.Food, error) {
	return m.GetFoodByIDsResult, m.GetFoodByIDsErr
}

func (m *Foodstuff) CreateFood(food *food.Food) error {
	return m.CreateFoodErr
}

func (m *Foodstuff) UpdateFood(food *food.Food) error {
	return m.UpdateFoodErr
}

func (m *Foodstuff) DeleteFood(id uuid.UUID) error {
	return m.DeleteFoodErr
}
