package wishlist

import (
	"TFP/database"
	"TFP/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var wishlistCollection *mongo.Collection = database.OpenCollection(database.Client, "wishlist")
var userCartCollection *mongo.Collection = database.OpenCollection(database.Client, "user_cart")

func (wishlist *Wishlist) Save() (*mongo.InsertOneResult, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	insertNumber, err := wishlistCollection.InsertOne(ctx, wishlist)
	defer cancel()
	if err != nil {
		return nil, utils.NewBadRequestError("Unable to add products to wishlist")
	}
	return insertNumber, nil
}

func (wishlist *Wishlist) GetUserWishlist() ([]Wishlist, float64, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var userWishList []Wishlist
	var totalPrice float64

	result, err := wishlistCollection.Find(context.TODO(), bson.M{"buyerid": wishlist.BuyerID})
	defer cancel()

	for result.Next(ctx) {
		var wishlist Wishlist
		err := result.Decode(&wishlist)
		if err != nil {
			fmt.Println("error")
		}
		totalPrice += wishlist.TotalAmount
		userWishList = append(userWishList, wishlist)
	}

	if err != nil {
		return nil, 0, utils.NewNotFoundError("Unable to list wishlist")
	}

	if err := result.All(ctx, &userWishList); err != nil {
		return nil, 0, utils.NewNotFoundError("Unable to list wishlist")
	}
	return userWishList, totalPrice, nil
}

func (wishlist *Wishlist) Remove() (*mongo.DeleteResult, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	filter := bson.M{"buyerid": wishlist.BuyerID, "productid": wishlist.ProductID}
	result, err := wishlistCollection.DeleteOne(ctx, filter)
	defer cancel()
	if err != nil {
		return nil, utils.NewNotFoundError("Unable to remove from wishlist")
	}
	return result, nil
}

func (wishlist *Wishlist) ResetWishlist() (*mongo.DeleteResult, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	filter := bson.M{"buyerid": wishlist.BuyerID, "productid": wishlist.ProductID}
	result, err := wishlistCollection.DeleteMany(ctx, filter)
	defer cancel()
	if result.DeletedCount == 0 {
		return nil, utils.NewNotFoundError("wishlist already empty")
	}
	if err != nil {
		return nil, utils.NewNotFoundError("Unable to reset wishlist")
	}
	return result, nil
}

func (wishlist *Wishlist) MoveToCart(userId string) (*mongo.InsertOneResult, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var cartInfo WishlistInfo

	filter := bson.M{"buyerid": userId}
	wishlistData, err := wishlistCollection.Find(ctx, filter)
	defer cancel()

	if wishlistData == nil {
		return nil, utils.NewBadRequestError("No Products found in wishlist")
	}
	if err != nil {
		return nil, utils.NewNotFoundError("Unable to find products from cart")
	}

	for wishlistData.Next(ctx) {
		var wishlist Wishlist
		err := wishlistData.Decode(&wishlist)
		if err != nil {
			fmt.Println("ERROR::::", err)
		}
		cartInfo.BuyerID = userId
		cartInfo.List = append(cartInfo.List, wishlist)
	}

	result, insertErr := userCartCollection.InsertOne(ctx, cartInfo)
	if insertErr != nil {
		return nil, utils.NewBadRequestError("error occured while adding to cart")
	}

	_, err = wishlistCollection.DeleteMany(ctx, bson.M{"buyerid": userId})
	if err != nil {
		return nil, utils.NewBadRequestError("unable to reset cart while place order")
	}

	return result, nil
}
