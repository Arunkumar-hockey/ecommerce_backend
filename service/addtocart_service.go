package service

import (
	"TFP/domain/addtocart"
	"TFP/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var (
	AddToCartService AddToCartServiceInterface = &addToCartService{}
)

type AddToCartServiceInterface interface {
	Create(addtocart.AddToCart) (*mongo.InsertOneResult, *utils.RestErr)
	GetUserCart(string) ([]addtocart.AddToCart, float64, *utils.RestErr)
	Remove(string, string) (*mongo.DeleteResult, *utils.RestErr)
	ResetCart(string) (*mongo.DeleteResult, *utils.RestErr)
}

type addToCartService struct{}

func (s *addToCartService) Create(cart addtocart.AddToCart) (*mongo.InsertOneResult, *utils.RestErr) {
	if validateErr := cart.Validate(); validateErr != nil {
		return nil, validateErr
	}

	cart.CreatedAt = time.Now()
	cart.UpdatedAt = time.Now()
	cart.TotalAmount = float64(cart.Price * cart.Quantity)
	insertNumber, err := cart.Save()
	if err != nil {
		return nil, err
	}
	return insertNumber, nil
}

func (s *addToCartService) GetUserCart(userID string) ([]addtocart.AddToCart, float64, *utils.RestErr) {
	cart := addtocart.AddToCart{BuyerID: userID}
	result, totalPrice, err := cart.GetUserCart()
	if err != nil {
		return nil, 0, err
	}
	return result, totalPrice, nil
}

func (s *addToCartService) Remove(userID, productID string) (*mongo.DeleteResult, *utils.RestErr) {
	result := &addtocart.AddToCart{BuyerID: userID, ProductID: productID}
	deleteID, deleteErr := result.Remove()
	if deleteErr != nil {
		return nil, deleteErr
	}
	return deleteID, nil
}

func (s *addToCartService) ResetCart(userID string) (*mongo.DeleteResult, *utils.RestErr) {
	result := addtocart.AddToCart{BuyerID: userID}
	deletedCount, deleteErr := result.ResetCart()
	if deleteErr != nil {
		return nil, deleteErr
	}
	return deletedCount, nil
}
