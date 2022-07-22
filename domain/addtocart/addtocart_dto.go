package addtocart

import (
	"TFP/utils"
	"strings"
	"time"
)

type AddToCart struct {
	BuyerID            string    `json:"buyer_id"`
	SellerID           string    `json:"seller_id"`
	ProductID          string    `json:"product_id"`
	ProductTitle       string    `json:"product_title"`
	ProductDescription string    `json:"product_description"`
	ProductImage       string    `json:"product_image"`
	Price              int       `json:"price"`
	PriceType          string    `json:"price_type"`
	Quantity           int       `json:"quantity"`
	TotalAmount        float64   `json:"total_amount"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
}

//type Cart struct {
//	Wishlist []wishlist.Wishlist `json:"wishlist"`
//}

func (cart *AddToCart) Validate() *utils.RestErr {
	if strings.TrimSpace(cart.BuyerID) == "" {
		return utils.NewBadRequestError("buyer ID required")
	}
	if strings.TrimSpace(cart.SellerID) == "" {
		return utils.NewBadRequestError("seller ID required")
	}
	if strings.TrimSpace(cart.ProductID) == "" {
		return utils.NewBadRequestError("product ID required")
	}
	if strings.TrimSpace(cart.ProductTitle) == "" {
		return utils.NewBadRequestError("product title required")
	}
	if strings.TrimSpace(cart.ProductDescription) == "" {
		return utils.NewBadRequestError("product description required")
	}
	if strings.TrimSpace(cart.ProductImage) == "" {
		return utils.NewBadRequestError("product image required")
	}
	if cart.Price == 0 {
		return utils.NewBadRequestError("product price required")
	}
	if strings.TrimSpace(cart.PriceType) == "" {
		return utils.NewBadRequestError("price type required")
	}
	if cart.Quantity == 0 {
		return utils.NewBadRequestError("product quantity required")
	}
	return nil
}
