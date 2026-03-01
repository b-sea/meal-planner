package food

import "github.com/google/uuid"

// Repository defines all functions required for interacting with persistent food data.
type Repository interface {
	FindFood() ([]*Food, error)
	GetFoodByID(id uuid.UUID) (*Food, error)
	GetFoodByIDs(ids []uuid.UUID) ([]*Food, error)
	CreateFood(food *Food) error
	UpdateFood(food *Food) error
	DeleteFood(id uuid.UUID) error
}
