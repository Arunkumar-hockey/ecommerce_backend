package order

import (
	"TFP/database"
	"TFP/domain/addtocart"
	"TFP/utils"
	"context"
	"fmt"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")
var userCartCollection *mongo.Collection = database.OpenCollection(database.Client, "user_cart")

func (order *Order) PlaceOrder(userID string) (*mongo.InsertOneResult, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var orderInfo Order

	cartData, err := userCartCollection.Find(context.TODO(), bson.M{"buyerid": userID})
	defer cancel()
	if err != nil {
		return nil, utils.NewNotFoundError("Unable to find products from cart")
	}

	for cartData.Next(ctx) {
		var cart addtocart.AddToCart
		err := cartData.Decode(&cart)
		if err != nil {
			fmt.Println(err)
		}
		orderInfo.OrderPrice += cart.TotalAmount
		orderInfo.BuyerID = cart.BuyerID
		uid := xid.New()
		orderInfo.OrderID = "Order" + uid.String()
		orderInfo.OrderedAt = time.Now()
		orderInfo.OrderStatus = "Successfully Placed"
		orderInfo.CartList = append(orderInfo.CartList, cart)
	}

	if orderInfo.CartList == nil {
		return nil, utils.NewBadRequestError("No Products found in cart")
	}

	result, insertErr := orderCollection.InsertOne(ctx, orderInfo)
	if insertErr != nil {
		return nil, utils.NewBadRequestError("unable to place order")
	}
	_, err = userCartCollection.DeleteMany(ctx, bson.M{"buyerid": userID})
	if err != nil {
		return nil, utils.NewBadRequestError("unable to reset cart while place order")
	}

	return result, nil
}

func (order *Order) GetOrder() ([]bson.M, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var allOrder []bson.M

	result, getErr := orderCollection.Find(context.TODO(), bson.M{"buyerid": order.BuyerID})
	defer cancel()
	if getErr != nil {
		return nil, utils.NewNotFoundError("No orders found")
	}

	if err := result.All(ctx, &allOrder); err != nil {
		return nil, utils.NewNotFoundError("Unable to find orders")
	}
	return allOrder, nil
}

func (order *Order) GetAllOrder() ([]bson.M, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var allOrder []bson.M

	result, getErr := orderCollection.Find(context.TODO(), bson.M{})
	defer cancel()
	if getErr != nil {
		return nil, utils.NewNotFoundError("No orders found")
	}

	if err := result.All(ctx, &allOrder); err != nil {
		return nil, utils.NewNotFoundError("Unable to find orders")
	}
	return allOrder, nil
}
