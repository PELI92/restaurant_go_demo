package repository

import (
	"context"
	"database/sql"
	"demo/internal/restaurant/model"
	"fmt"
)

func CreateRestaurant(ctx context.Context, r model.Restaurant) (sql.Result, error) {
	fmt.Println("CreateRestaurant")
	return nil, nil
}

func FetchRestaurants(ctx context.Context) ([]model.Restaurant, error) {
	fmt.Println("FetchRestaurants")
	return []model.Restaurant{}, nil
}
func FetchRestaurantByID(ctx context.Context, id string) (model.Restaurant, error) {
	fmt.Println("FetchRestaurantByID")
	return model.Restaurant{}, nil
}
func ModifyRestaurant(ctx context.Context, id string, r model.Restaurant) (sql.Result, error) {
	fmt.Println("ModifyRestaurant")
	return nil, nil
}
func DeleteRestaurant(ctx context.Context, id string) (sql.Result, error) {
	fmt.Println("DeleteRestaurant")
	return nil, nil
}
