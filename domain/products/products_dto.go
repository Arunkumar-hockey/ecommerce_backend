package products

import (
	"TFP/utils"
	"strings"
	"time"
)

type Product struct {
	SellerID     string    `json:"seller_id"`
	ProductID    string    `json:"product_id"`
	CategoryID   string    `json:"category_id"`
	ProductName  string    `json:"product_name"`
	Quantity     int       `json:"quantity"`
	Unit         string    `json:"unit"`
	Price        float64   `json:"price"`
	ProductImage string    `json:"product_image"`
	PriceType    string    `json:"price_type"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (product *Product) Validate() *utils.RestErr {
	if strings.TrimSpace(product.SellerID) == "" {
		return utils.NewBadRequestError("seller ID required")
	}
	if strings.TrimSpace(product.CategoryID) == "" {
		return utils.NewBadRequestError("category ID required")
	}
	if strings.TrimSpace(product.ProductName) == "" {
		return utils.NewBadRequestError("product name required")
	}
	if product.Price == 0 {
		return utils.NewBadRequestError("product price required")
	}
	if strings.TrimSpace(product.ProductImage) == "" {
		return utils.NewBadRequestError("product image required")
	}
	if strings.TrimSpace(product.PriceType) == "" {
		return utils.NewBadRequestError("price type required")
	}
	if strings.TrimSpace(product.Description) == "" {
		return utils.NewBadRequestError("description required")
	}
	return nil
}
