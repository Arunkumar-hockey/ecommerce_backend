package wishlist

import (
	"TFP/utils"
	"strings"
	"time"
)

type Wishlist struct {
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
}

type WishlistInfo struct {
	BuyerID string     `json:"buyer_id"`
	List    []Wishlist `json:"list"`
}

func (wishlist *Wishlist) Validate() *utils.RestErr {
	if strings.TrimSpace(wishlist.BuyerID) == "" {
		utils.NewBadRequestError("buyer id required")
	}
	if strings.TrimSpace(wishlist.SellerID) == "" {
		utils.NewBadRequestError("seller id required")
	}
	if strings.TrimSpace(wishlist.ProductID) == "" {
		utils.NewBadRequestError("product id required")
	}
	if strings.TrimSpace(wishlist.ProductTitle) == "" {
		utils.NewBadRequestError("product title required")
	}
	if strings.TrimSpace(wishlist.ProductDescription) == "" {
		utils.NewBadRequestError("product description required")
	}
	if strings.TrimSpace(wishlist.ProductImage) == "" {
		utils.NewBadRequestError("product image required")
	}
	if wishlist.Price == 0 {
		utils.NewBadRequestError("product price required")
	}
	if strings.TrimSpace(wishlist.PriceType) == "" {
		utils.NewBadRequestError("price type required")
	}
	if wishlist.Quantity == 0 {
		utils.NewBadRequestError("quantity required")
	}
	if wishlist.TotalAmount == 0 {
		utils.NewBadRequestError("total amount required")
	}
	return nil
}
