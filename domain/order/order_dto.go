package order

import (
	"TFP/domain/addtocart"
	"time"
)

type Buycart struct {
	Product_Title string `json:"title"`
	Cdnlink       string `json:"cdn"`
	Description   string `json:"description"`
	OrderStatus   string `json:"orderstatus"`
	Sellerid      string `json:"sellerid"`
	Buyerid       string `json:"buyerid"`
	Productid     string `json:"productid"`
	Price         int    `json:"price"`
	CurrencyType  string `json:"currencyType"`
	Quantity      int    `json:"quantity"`
	TotalAmount   int    `json:"total"`
	ShipInsurance string `json:"shipinsurance"`
	CreatedAt     string `json:"created_At"`
	UpdatedAt     string `json:"updated_At"`
}

type Order struct {
	OrderID       string                `json:"order_id"`
	BuyerID       string                `json:"buyer_id"`
	CartList      []addtocart.AddToCart `json:"cart_list"`
	OrderStatus   string                `json:"order_status"`
	ShipInsurance string                `json:"ship_insurance"`
	OrderPrice    float64               `json:"order_price"`
	OrderedAt     time.Time             `json:"ordered_at"`
	PaymentMethod Payment               `json:"payment_method"`
}

type Payment struct {
	Digital bool `json:"digital"`
	COD     bool `json:"cod"`
}
