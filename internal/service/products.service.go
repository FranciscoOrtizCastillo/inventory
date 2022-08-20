package service

import (
	context "context"
	"errors"

	models "github.com/FranciscoOrtizCastillo/inventory/internal/models"
)

var validRolesAddProduct []int64 = []int64{1, 2}
var ErrInvalidPermissions = errors.New("user does not have permissions to add product")

func (s *serv) GetProducts(ctx context.Context) ([]models.Product, error) {

	pp, err := s.repo.GetProducts(ctx)

	if err != nil {
		return nil, err
	}

	products := []models.Product{}

	for _, p := range pp {
		products = append(products, models.Product{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		})
	}

	return products, nil
}

func (s *serv) GetProductByID(ctx context.Context, id int64) (*models.Product, error) {

	P, err := s.repo.GetProductByID(ctx, id)

	if err != nil {
		return nil, err
	}

	product := &models.Product{
		ID:          P.ID,
		Name:        P.Name,
		Description: P.Description,
		Price:       P.Price,
	}

	return product, nil
}

func (s *serv) AddProduct(ctx context.Context, product models.Product, email string) error {

	u, err := s.repo.GetUserByEmail(ctx, email)

	if err != nil {
		return err
	}

	roles, err := s.repo.GetUserRoles(ctx, u.ID)

	if err != nil {
		return err
	}

	userCanAddProduct := false

	for _, r := range roles {
		for _, vr := range validRolesAddProduct {
			if vr == r.RoleID {
				userCanAddProduct = true
			}
		}
	}

	if !userCanAddProduct {
		return ErrInvalidPermissions
	}

	return s.repo.SaveProduct(
		ctx,
		product.Name,
		product.Description,
		product.Price,
		u.ID,
	)
}
