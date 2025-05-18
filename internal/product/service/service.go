package service

import (
	"context"
	api_errors "demo/api/error"
	"demo/internal/product/service/domain"
	"demo/internal/product/service/param"
	"fmt"
	"github.com/google/uuid"
)

func CreateProduct(ctx context.Context, param param.CreateProductParam) (domain.Product, *api_errors.APIError) {
	fmt.Printf("CreateProduct: %+v\n", param)
	return domain.Product{}, nil
}

func GetProductsPaginated(ctx context.Context, limit, offset int) ([]domain.Product, *api_errors.APIError) {
	fmt.Printf("GetProductsPaginated: limit=%d, offset=%d\n", limit, offset)
	return []domain.Product{}, nil
}

func GetProductByID(ctx context.Context, id uuid.UUID) (*domain.Product, *api_errors.APIError) {
	fmt.Printf("GetProductByID: %s\n", id)
	return &domain.Product{}, nil
}

func UpdateProduct(ctx context.Context, id uuid.UUID, param param.UpdateProductParam) (domain.Product, *api_errors.APIError) {
	fmt.Printf("UpdateProduct: %s > %+v \n", id, param)
	return domain.Product{}, nil
}

func DeleteProduct(ctx context.Context, id uuid.UUID) *api_errors.APIError {
	fmt.Printf("DeleteProduct: %s\n", id)
	return nil
}
