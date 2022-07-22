package admin

import (
	"TFP/database"
	"TFP/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var adminCollection *mongo.Collection = database.OpenCollection(database.Client, "admin")

func (admin *Admin) Save() (*mongo.InsertOneResult, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	emailCount, _ := adminCollection.CountDocuments(ctx, bson.M{"email": admin.Email})

	defer cancel()
	if emailCount > 0 {
		return nil, utils.NewBadRequestError("Email ID already registered")
	}

	phoneCount, _ := adminCollection.CountDocuments(ctx, bson.M{"mobile": admin.Mobile})
	if phoneCount > 0 {
		return nil, utils.NewBadRequestError("Phone number already registered")
	}

	insertNumber, insertErr := adminCollection.InsertOne(ctx, admin)
	if insertErr != nil {
		utils.BuildErrorResponse(insertErr.Error())
	}
	defer cancel()
	return insertNumber, nil
}

func (admin *Admin) GetAll() ([]bson.M, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var allUsers []bson.M

	result, getErr := adminCollection.Find(context.TODO(), bson.M{})
	defer cancel()
	if getErr != nil {
		return nil, utils.NewNotFoundError("Unable to find admin")
	}

	if err := result.All(ctx, &allUsers); err != nil {
		return nil, utils.NewNotFoundError("Unable to find admin")
	}

	return allUsers, nil
}

func (admin *Admin) GetByID() *utils.RestErr {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	err := adminCollection.FindOne(ctx, bson.M{"adminid": admin.AdminID}).Decode(&admin)
	defer cancel()

	if err != nil {
		return utils.NewNotFoundError("Unable to find admin")
	}
	return nil
}

func (admin *Admin) Update() (*mongo.UpdateResult, *utils.RestErr) {
	return nil, nil
}

func (admin *Admin) Delete() (*mongo.DeleteResult, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	filter := bson.M{"adminid": admin.AdminID}
	result, err := adminCollection.DeleteOne(ctx, filter)
	defer cancel()
	if err != nil {
		return nil, utils.NewNotFoundError("Unable to delete admin")
	}
	return result, nil
}

func (admin *Admin) Login() *utils.RestErr {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	err := adminCollection.FindOne(ctx, bson.M{"email": admin.Email, "password": admin.Password}).Decode(&admin)
	defer cancel()

	if err != nil {
		return utils.NewNotFoundError("Invalid Credentials")
	}
	return nil
}
