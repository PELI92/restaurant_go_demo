package model

// Product represents an item available in the catalog.
// swagger:model Product
type Product struct {
	// ID is the unique identifier for the product.
	// example: "123e4567-e89b-12d3-a456-426614174000"
	ID string `json:"id,omitempty"`

	// URL is the link to the product resource or image.
	// example: "https://example.com/api/products/123e4567-e89b-12d3-a456-426614174000"
	Name string `json:"url"`

	// Description gives a human-readable summary of the product.
	// example: "Ergonomic wireless mouse with adjustable DPI and long battery life."
	Description string `json:"description"`
}
