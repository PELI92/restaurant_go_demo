package service

import (
	errors "demo/api/error"
	"demo/internal/restaurant/model"
	"fmt"
)

func CreateRestaurant(input model.Restaurant) (model.Restaurant, *errors.APIError) {
	fmt.Printf("CreateRestaurant: %+v\n", input)
	return input, nil
}

func GetRestaurants(id string) ([]model.Restaurant, *errors.APIError) {
	fmt.Printf("CreateRestaurant: %s\n", id)
	return []model.Restaurant{}, nil
}

func GetRestaurantByID(id string) (model.Restaurant, *errors.APIError) {
	fmt.Printf("GetRestaurantByID: %s\n", id)
	return model.Restaurant{}, nil
}

func UpdateRestaurant(id string, input model.Restaurant) (model.Restaurant, *errors.APIError) {
	fmt.Printf("UpdateRestaurant: %s > %+v \n", id, input)
	return input, nil
}

func DeleteRestaurant(id string) *errors.APIError {
	fmt.Printf("DeleteRestaurant: %s\n", id)
	return nil
}
