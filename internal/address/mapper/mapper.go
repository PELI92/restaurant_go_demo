package mapper

import (
	dto "demo/internal/address/controller/dto"
	"demo/internal/address/repository/entity"
	"demo/internal/address/service/domain"
	"demo/internal/address/service/param"
)

func NewCreateAddressParam(input dto.CreateAddressRequest) param.CreateAddressParam {
	return param.CreateAddressParam{
		Street:     input.Street,
		City:       input.City,
		State:      input.State,
		PostalCode: input.PostalCode,
		Country:    input.Country,
	}
}

func NewAddressResponse(domain domain.Address) dto.AddressResponse {
	return dto.AddressResponse{
		ID:         domain.ID,
		City:       domain.City,
		State:      domain.State,
		PostalCode: domain.PostalCode,
		Country:    domain.Country,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

func NewUpdateAddressParam(input dto.UpdateAddressRequest) param.UpdateAddressParam {
	return param.UpdateAddressParam{
		Street:     input.Street,
		City:       input.City,
		State:      input.State,
		PostalCode: input.PostalCode,
		Country:    input.Country,
	}
}

func NewAddressDomain(address entity.Address) domain.Address {
	return domain.Address{
		ID:         address.ID,
		City:       address.City,
		State:      *address.State,
		PostalCode: *address.PostalCode,
		Country:    address.Country,
		CreatedAt:  address.CreatedAt,
		UpdatedAt:  address.UpdatedAt,
	}
}
