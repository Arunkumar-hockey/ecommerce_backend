package seller

import (
	"TFP/database"
	"TFP/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var sellerCollection *mongo.Collection = database.OpenCollection(database.Client, "seller")

func (seller *Seller) Save() (*mongo.InsertOneResult, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	emailCount, _ := sellerCollection.CountDocuments(ctx, bson.M{"email": seller.Email})

	defer cancel()
	if emailCount > 0 {
		return nil, utils.NewBadRequestError("Email ID already registered")
	}

	phoneCount, _ := sellerCollection.CountDocuments(ctx, bson.M{"mobile": seller.Mobile})
	if phoneCount > 0 {
		return nil, utils.NewBadRequestError("Phone number already registered")
	}

	insertNumber, insertErr := sellerCollection.InsertOne(ctx, seller)
	if insertErr != nil {
		utils.BuildErrorResponse(insertErr.Error())
	}
	defer cancel()
	return insertNumber, nil
}

func (seller *Seller) GetAll() ([]bson.M, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var allSellers []bson.M

	result, getErr := sellerCollection.Find(context.TODO(), bson.M{})
	defer cancel()
	if getErr != nil {
		return nil, utils.NewNotFoundError("Unable to find seller")
	}

	if err := result.All(ctx, &allSellers); err != nil {
		return nil, utils.NewNotFoundError("Unable to find seller")
	}

	return allSellers, nil
}

func (seller *Seller) GetByID() *utils.RestErr {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	err := sellerCollection.FindOne(ctx, bson.M{"sellerid": seller.Sellerid}).Decode(&seller)
	defer cancel()

	if err != nil {
		return utils.NewNotFoundError("Unable to find seller")
	}
	return nil
}

func (seller *Seller) Update() (*mongo.UpdateResult, *utils.RestErr) {
	return nil, nil
}

func (seller *Seller) Delete() (*mongo.DeleteResult, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	filter := bson.M{"sellerid": seller.Sellerid}
	result, err := sellerCollection.DeleteOne(ctx, filter)
	defer cancel()
	if err != nil {
		return nil, utils.NewNotFoundError("Unable to delete seller")
	}
	return result, nil
}

func (seller *Seller) Login() *utils.RestErr {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	err := sellerCollection.FindOne(ctx, bson.M{"email": seller.Email, "password": seller.Password}).Decode(&seller)
	defer cancel()

	if err != nil {
		return utils.NewNotFoundError("Invalid Credentials")
	}
	return nil
}
