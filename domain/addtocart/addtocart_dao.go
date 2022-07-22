package addtocart

import (
	"TFP/database"
	"TFP/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var userCartCollection *mongo.Collection = database.OpenCollection(database.Client, "user_cart")

func (cart *AddToCart) Save() (*mongo.InsertOneResult, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	insertNumber, insertErr := userCartCollection.InsertOne(ctx, cart)
	defer cancel()

	if insertErr != nil {
		return nil, utils.NewBadRequestError("Unable to add products to cart")
	}
	return insertNumber, nil
}

func (cart *AddToCart) GetUserCart() ([]AddToCart, float64, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var userCart []AddToCart
	var totalPrice float64

	result, err := userCartCollection.Find(context.TODO(), bson.M{"buyerid": cart.BuyerID})
	defer cancel()

	for result.Next(ctx) {
		var cart AddToCart
		err := result.Decode(&cart)
		if err != nil {
			fmt.Println("error")
		}
		totalPrice += cart.TotalAmount
		userCart = append(userCart, cart)
	}

	if err != nil {
		return nil, 0, utils.NewNotFoundError("Unable to list cart")
	}

	if err := result.All(ctx, &userCart); err != nil {
		return nil, 0, utils.NewNotFoundError("Unable to list cart")
	}
	return userCart, totalPrice, nil
}

func (cart *AddToCart) Remove() (*mongo.DeleteResult, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	filter := bson.M{"buyerid": cart.BuyerID, "productid": cart.ProductID}
	result, err := userCartCollection.DeleteOne(ctx, filter)
	defer cancel()
	if err != nil {
		return nil, utils.NewNotFoundError("Unable to remove from cart")
	}
	return result, nil
}

func (cart *AddToCart) ResetCart() (*mongo.DeleteResult, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	filter := bson.M{"buyerid": cart.BuyerID}
	result, err := userCartCollection.DeleteMany(ctx, filter)
	defer cancel()
	if result.DeletedCount == 0 {
		return nil, utils.NewNotFoundError("cart already empty")
	}
	if err != nil {
		return nil, utils.NewNotFoundError("Unable to reset cart")
	}
	return result, nil
}
