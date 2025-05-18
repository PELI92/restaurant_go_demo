package param

type CreateAddressParam struct {
	Street     string `json:"street" `
	City       string `json:"city" `
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

type UpdateAddressParam struct {
	Street     string `json:"street" `
	City       string `json:"city" `
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}
