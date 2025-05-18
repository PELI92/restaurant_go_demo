package service

import (
	"context"
	"database/sql"
	api_errors "demo/api/error"
	"demo/internal/address/mapper"
	"demo/internal/address/repository"
	"demo/internal/address/repository/entity"
	"demo/internal/address/service/domain"
	"demo/internal/address/service/param"
	"errors"
	"github.com/google/uuid"
)

func CreateRestaurantAddress(ctx context.Context, param param.CreateAddressParam) (*domain.Address, *api_errors.APIError) {
	address := entity.Address{
		Street:     param.Street,
		City:       param.City,
		State:      &param.State,
		PostalCode: &param.PostalCode,
		Country:    param.Country,
	}

	id, err := repository.CreateAddress(ctx, address)
	if err != nil {
		return nil, api_errors.NewUnknown(err.Error())
	}

	result, err := repository.GetAddressByID(ctx, *id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, api_errors.NewNotFound("restaurant not found")
		}
		return nil, api_errors.NewUnknown(err.Error())
	}

	addressDomain := mapper.NewAddressDomain(*result)
	return &addressDomain, nil
}

func GetAddressById(ctx context.Context, id uuid.UUID) (*domain.Address, *api_errors.APIError) {
	result, err := repository.GetAddressByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, api_errors.NewNotFound("restaurant not found")
		}
		return nil, api_errors.NewUnknown(err.Error())
	}

	addressDomain := mapper.NewAddressDomain(*result)
	return &addressDomain, nil
}

func GetAddressByRestaurantId(ctx context.Context, id uuid.UUID) (*domain.Address, *api_errors.APIError) {
	result, err := repository.GetAddressByRestaurantId(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, api_errors.NewNotFound("restaurant not found")
		}
		return nil, api_errors.NewUnknown(err.Error())
	}

	addressDomain := mapper.NewAddressDomain(*result)
	return &addressDomain, nil
}

func UpdateAddress(ctx context.Context, id uuid.UUID, param param.UpdateAddressParam) (*domain.Address, *api_errors.APIError) {
	address := entity.Address{
		Street:     param.Street,
		City:       param.City,
		State:      &param.State,
		PostalCode: &param.PostalCode,
		Country:    param.Country,
	}

	result, err := repository.UpdateAddress(ctx, id, address)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, api_errors.NewNotFound("address not found")
		}
		return nil, api_errors.NewUnknown(err.Error())
	}

	addressDomain := mapper.NewAddressDomain(*result)
	return &addressDomain, nil
}

func DeleteAddress(ctx context.Context, id uuid.UUID) *api_errors.APIError {
	err := repository.DeleteAddress(ctx, id)
	if err != nil {
		return api_errors.NewUnknown(err.Error())
	}
	return nil
}
