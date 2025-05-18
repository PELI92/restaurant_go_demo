package service

import (
	"context"
	"database/sql"
	api_errors "demo/api/error"
	"demo/internal/restaurant/mapper"
	"demo/internal/restaurant/repository"
	"demo/internal/restaurant/repository/entity"
	"demo/internal/restaurant/service/domain"
	"demo/internal/restaurant/service/param"
	"errors"
	"github.com/google/uuid"
)

func CreateRestaurant(ctx context.Context, param param.CreateRestaurantParam) (*domain.Restaurant, *api_errors.APIError) {
	restaurant := entity.Restaurant{
		Name:    param.Name,
		Phone:   &param.Phone,
		Website: param.Website,
		Email:   &param.Email,
	}

	id, err := repository.CreateRestaurant(ctx, restaurant)
	if err != nil {
		return nil, api_errors.NewUnknown(err.Error())
	}

	result, err := repository.GetRestaurantByID(ctx, *id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, api_errors.NewNotFound("restaurant not found")
		}
		return nil, api_errors.NewUnknown(err.Error())
	}

	restaurantDomain := mapper.NewRestaurantDomain(*result)
	return &restaurantDomain, nil
}

func GetRestaurantsPaginated(ctx context.Context, limit, offset int) ([]domain.Restaurant, *api_errors.APIError) {
	results, err := repository.GetRestaurantsPaginated(ctx, limit, offset)
	if err != nil {
		return nil, api_errors.NewUnknown(err.Error())
	}

	restaurants := make([]domain.Restaurant, 0, len(results))
	for _, e := range results {
		restaurants = append(restaurants, mapper.NewRestaurantDomain(e))
	}
	return restaurants, nil
}

func GetRestaurantByID(ctx context.Context, id uuid.UUID) (*domain.Restaurant, *api_errors.APIError) {
	result, err := repository.GetRestaurantByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, api_errors.NewNotFound("restaurant not found")
		}
		return nil, api_errors.NewUnknown(err.Error())
	}

	restaurantDomain := mapper.NewRestaurantDomain(*result)
	return &restaurantDomain, nil
}

func UpdateRestaurant(ctx context.Context, id uuid.UUID, param param.UpdateRestaurantParam) (*domain.Restaurant, *api_errors.APIError) {
	restaurant := entity.Restaurant{
		Name:    *param.Name,
		Phone:   param.Phone,
		Website: param.Website,
		Email:   param.Email,
	}

	result, err := repository.UpdateRestaurant(ctx, id, restaurant)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, api_errors.NewNotFound("restaurant not found")
		}
		return nil, api_errors.NewUnknown(err.Error())
	}

	restaurantDomain := mapper.NewRestaurantDomain(*result)
	return &restaurantDomain, nil
}

func DeleteRestaurant(ctx context.Context, id uuid.UUID) *api_errors.APIError {
	err := repository.DeleteRestaurant(ctx, id)
	if err != nil {
		return api_errors.NewUnknown(err.Error())
	}
	return nil
}
