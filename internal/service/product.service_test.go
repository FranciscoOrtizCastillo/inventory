package service

import (
	context "context"
	"testing"

	models "github.com/FranciscoOrtizCastillo/inventory/internal/models"
)

func TestAddProduct(t *testing.T) {

	testCases := []struct {
		Name          string
		Product       models.Product
		email         string
		ExpectedError error
	}{
		{
			Name: "AddProduct_Success",
			Product: models.Product{
				Name:        "Test Product 1",
				Description: "Test Description 1",
				Price:       10.0,
			},
			email:         "admin@email.com",
			ExpectedError: nil,
		},
		{
			Name: "AddProduct_InvalidPermissions",
			Product: models.Product{
				Name:        "Test Product 2",
				Description: "Test Description 2",
				Price:       10.0,
			},
			email:         "customer@email.com",
			ExpectedError: ErrInvalidPermissions,
		},
	}

	ctx := context.Background()

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			repo.Mock.Test(t)

			err := s.AddProduct(ctx, tc.Product, tc.email)

			if err != tc.ExpectedError {
				t.Errorf("Expected error %v, got %v", tc.ExpectedError, err)
			}
		})
	}

}
