package service

import (
	errors "demo/api/error"
	"demo/internal/product/model"
	"fmt"
)

func CreateProduct(input model.Product) (model.Product, *errors.APIError) {
	fmt.Printf("CreateProduct: %+v\n", input)
	return input, nil
}

func GetProducts(id string) ([]model.Product, *errors.APIError) {
	fmt.Printf("CreateProduct: %s\n", id)
	return []model.Product{}, nil
}

func GetProductByID(id string) (model.Product, *errors.APIError) {
	fmt.Printf("GetProductByID: %s\n", id)
	return model.Product{}, nil
}

func UpdateProduct(id string, input model.Product) (model.Product, *errors.APIError) {
	fmt.Printf("UpdateProduct: %s > %+v \n", id, input)
	return input, nil
}

func DeleteProduct(id string) *errors.APIError {
	fmt.Printf("DeleteProduct: %s\n", id)
	return nil
}
