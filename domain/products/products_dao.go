package products

import (
	"TFP/database"
	"TFP/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var productCollection *mongo.Collection = database.OpenCollection(database.Client, "products")

func (product *Product) Save() (*mongo.InsertOneResult, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	insertNumber, insertErr := productCollection.InsertOne(ctx, product)
	defer cancel()

	if insertErr != nil {
		return nil, utils.NewBadRequestError("Unable to save the document")
	}

	return insertNumber, nil
}

func (product *Product) GetAll() ([]bson.M, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var allProducts []bson.M

	result, err := productCollection.Find(context.TODO(), bson.M{})
	defer cancel()

	if err != nil {
		return nil, utils.NewNotFoundError("Unable to find products")
	}

	if err := result.All(ctx, &allProducts); err != nil {
		return nil, utils.NewNotFoundError("Unable to find products")
	}
	return allProducts, nil
}

func (product *Product) GetByID() *utils.RestErr {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	err := productCollection.FindOne(ctx, bson.M{"productid": product.ProductID}).Decode(&product)
	defer cancel()

	if err != nil {
		fmt.Println("false....", err)
		return utils.NewNotFoundError("Unable to find product")
	}
	return nil
}

func (product *Product) GetProductByCategory() ([]bson.M, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var allProducts []bson.M
	result, err := productCollection.Find(context.TODO(), bson.M{"categoryid": product.CategoryID})
	defer cancel()

	if err != nil {
		return nil, utils.NewNotFoundError("Unable to find products")
	}
	if err := result.All(ctx, &allProducts); err != nil {
		return nil, utils.NewNotFoundError("Unable to find products")
	}
	return allProducts, nil
}

func (product *Product) Update() (*mongo.UpdateResult, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	filter := bson.M{"sellerid": product.SellerID, "productid": product.ProductID}
	opts := options.Update().SetUpsert(true)
	update := bson.M{"set": bson.M{"productname": product.ProductName, "price": product.Price, "pricetype": product.PriceType, "description": product.Description}}

	result, err := productCollection.UpdateOne(
		ctx,
		filter,
		update,
		opts,
	)
	defer cancel()
	if err != nil {
		return nil, utils.NewNotFoundError("Unable to update product")
	}
	return result, nil
}

func (product *Product) Delete() (*mongo.DeleteResult, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	filter := bson.M{"productid": product.ProductID}
	result, err := productCollection.DeleteOne(ctx, filter)
	defer cancel()
	if err != nil {
		return nil, utils.NewNotFoundError("Unable to delete product")
	}
	return result, nil
}
