package category

import (
	"TFP/database"
	"TFP/utils"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var categoryCollection *mongo.Collection = database.OpenCollection(database.Client, "category")

func (category *Category) Create() (*mongo.InsertOneResult, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	insertNumber, insertErr := categoryCollection.InsertOne(ctx, category)
	defer cancel()

	if insertErr != nil {
		return nil, utils.NewBadRequestError("Unable to save the document")
	}

	return insertNumber, nil
}

func (category *Category) GetAll() ([]bson.M, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var allCategory []bson.M

	result, err := categoryCollection.Find(context.TODO(), bson.M{})
	defer cancel()

	if err != nil {
		return nil, utils.NewNotFoundError("Unable to find category")
	}

	if err := result.All(ctx, &allCategory); err != nil {
		return nil, utils.NewNotFoundError("Unable to find category")
	}
	return allCategory, nil
}
