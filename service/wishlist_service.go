package service

import (
	"TFP/domain/wishlist"
	"TFP/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var (
	WishlistService WishlistServiceInterface = &wishlistService{}
)

type WishlistServiceInterface interface {
	Create(wishlist.Wishlist) (*mongo.InsertOneResult, *utils.RestErr)
	GetUserWishlist(string) ([]wishlist.Wishlist, float64, *utils.RestErr)
	Remove(string, string) (*mongo.DeleteResult, *utils.RestErr)
	ResetWishlist(string) (*mongo.DeleteResult, *utils.RestErr)
	MoveToCart(string) (*mongo.InsertOneResult, *utils.RestErr)
}

type wishlistService struct{}

func (s *wishlistService) Create(wishlist wishlist.Wishlist) (*mongo.InsertOneResult, *utils.RestErr) {
	if validateErr := wishlist.Validate(); validateErr != nil {
		return nil, validateErr
	}
	wishlist.CreatedAt = time.Now()
	wishlist.TotalAmount = float64(wishlist.Price * wishlist.Quantity)

	insertNumber, err := wishlist.Save()
	if err != nil {
		return nil, err
	}
	return insertNumber, nil
}

func (s *wishlistService) GetUserWishlist(userID string) ([]wishlist.Wishlist, float64, *utils.RestErr) {
	wishlist := &wishlist.Wishlist{BuyerID: userID}

	result, totalPrice, err := wishlist.GetUserWishlist()
	if err != nil {
		return nil, 0, err
	}
	return result, totalPrice, nil
}

func (s *wishlistService) Remove(userID, productID string) (*mongo.DeleteResult, *utils.RestErr) {
	result := &wishlist.Wishlist{BuyerID: userID, ProductID: productID}
	deleteID, err := result.Remove()
	if err != nil {
		return nil, err
	}
	return deleteID, nil
}

func (s *wishlistService) ResetWishlist(userID string) (*mongo.DeleteResult, *utils.RestErr) {
	result := &wishlist.Wishlist{BuyerID: userID}
	deletedCount, err := result.ResetWishlist()
	if err != nil {
		return nil, err
	}
	return deletedCount, nil
}

func (s *wishlistService) MoveToCart(userID string) (*mongo.InsertOneResult, *utils.RestErr) {
	//cart := addtocart.AddToCart{BuyerID: userID}
	wishlist := wishlist.Wishlist{BuyerID: userID}
	result, err := wishlist.MoveToCart(userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
