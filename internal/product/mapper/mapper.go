package mapper

import (
	dto "demo/internal/product/controller/dto"
	"demo/internal/product/repository/entity"
	"demo/internal/product/service/domain"
	"demo/internal/product/service/param"
)

func NewCreateProductParam(d dto.CreateProductRequest) param.CreateProductParam {
	return param.CreateProductParam{
		RestaurantID: d.RestaurantID,
		Name:         d.Name,
		Description:  d.Description,
		Price:        d.Price,
	}
}

func NewUpdateProductParam(d dto.UpdateProductRequest) param.UpdateProductParam {
	return param.UpdateProductParam{
		RestaurantID: d.RestaurantID,
		Name:         d.Name,
		Description:  d.Description,
		Price:        d.Price,
	}
}

func NewProductResponse(d domain.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:           d.ID,
		RestaurantID: d.RestaurantID,
		Name:         d.Name,
		Description:  d.Description,
		Price:        d.Price,
		CreatedAt:    d.CreatedAt,
		UpdatedAt:    d.UpdatedAt,
	}
}

func NewProductDomain(e entity.Product) domain.Product {
	return domain.Product{
		ID:           e.ID,
		RestaurantID: e.RestaurantID,
		Name:         e.Name,
		Description:  e.Description,
		Price:        e.Price,
		CreatedAt:    e.CreatedAt,
		UpdatedAt:    e.UpdatedAt,
	}
}
