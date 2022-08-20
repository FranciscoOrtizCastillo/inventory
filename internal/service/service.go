package service

import (
	"context"

	"github.com/FranciscoOrtizCastillo/inventory/internal/models"
	"github.com/FranciscoOrtizCastillo/inventory/internal/repository"
)

// Service is teh bussines logic os the application
//go:generate mockery --name=Service --output=service --inpackage
type Service interface {
	
	RegisterUser(ctx context.Context, email, name, password string) error
	LoginUser(ctx context.Context, email, password string) (*models.User, error)

	AddUserRole(ctx context.Context, userID, roleID int64) error
	RemoveUserRole(ctx context.Context, userID, roleID int64) error

	GetProducts(ctx context.Context) ([]models.Product, error)
	GetProductByID(ctx context.Context, id int64) (*models.Product, error)

	AddProduct(ctx context.Context, product models.Product, userEmail string) error

}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
