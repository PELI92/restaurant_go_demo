package mapper

import (
	"demo/internal/restaurant/controller/dto"
	"demo/internal/restaurant/repository/entity"
	"demo/internal/restaurant/service/domain"
	"demo/internal/restaurant/service/param"
)

func NewCreateRestaurantParam(d dto.CreateRestaurantRequest) param.CreateRestaurantParam {
	return param.CreateRestaurantParam{
		Name:    d.Name,
		Phone:   d.Phone,
		Website: &d.Website,
		Email:   d.Email,
	}
}

func NewUpdateRestaurantParam(d dto.UpdateRestaurantRequest) param.UpdateRestaurantParam {
	return param.UpdateRestaurantParam{
		Name:    d.Name,
		Phone:   d.Phone,
		Website: d.Website,
		Email:   d.Email,
	}
}

func NewRestaurantResponse(d domain.Restaurant) dto.RestaurantResponse {
	return dto.RestaurantResponse{
		ID:        d.ID,
		Name:      d.Name,
		Phone:     d.Phone,
		Website:   d.Website,
		Email:     d.Email,
		Status:    d.Status,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func NewRestaurantDomain(r entity.Restaurant) domain.Restaurant {
	return domain.Restaurant{
		ID:        r.ID,
		Name:      r.Name,
		Phone:     r.Phone,
		Website:   r.Website,
		Email:     r.Email,
		AddressID: r.AddressID,
		Status:    r.Status,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}
