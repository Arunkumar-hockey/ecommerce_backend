package buyer

import (
	"TFP/database"
	"TFP/utils"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "buyer")

func (user *User) Save() (*mongo.InsertOneResult, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	emailCount, _ := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})

	defer cancel()
	if emailCount > 0 {
		return nil, utils.NewBadRequestError("Email ID already registered")
	}

	phoneCount, _ := userCollection.CountDocuments(ctx, bson.M{"mobile": user.Mobile})
	if phoneCount > 0 {
		return nil, utils.NewBadRequestError("Phone number already registered")
	}

	insertNumber, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr != nil {
		utils.BuildErrorResponse(insertErr.Error())
	}
	defer cancel()
	return insertNumber, nil
}

func (user *User) GetAll() ([]bson.M, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var allUsers []bson.M

	result, getErr := userCollection.Find(context.TODO(), bson.M{})
	defer cancel()
	if getErr != nil {
		return nil, utils.NewNotFoundError("Unable to find buyer")
	}

	if err := result.All(ctx, &allUsers); err != nil {
		return nil, utils.NewNotFoundError("Unable to find buyer")
	}

	return allUsers, nil
}

func (user *User) GetByID() *utils.RestErr {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	err := userCollection.FindOne(ctx, bson.M{"userid": user.Userid}).Decode(&user)
	defer cancel()

	if err != nil {
		fmt.Println("false....", err)
		return utils.NewNotFoundError("Unable to find user")
	}
	fmt.Println("true....")
	return nil
}

func (user *User) Update() (*mongo.UpdateResult, *utils.RestErr) {
	return nil, nil
}

func (user *User) Delete() (*mongo.DeleteResult, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	filter := bson.M{"userid": user.Userid}
	result, err := userCollection.DeleteOne(ctx, filter)
	defer cancel()
	if err != nil {
		return nil, utils.NewNotFoundError("Unable to delete user")
	}
	return result, nil
}

func (user *User) Login() *utils.RestErr {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email, "password": user.Password}).Decode(&user)
	defer cancel()

	if err != nil {
		fmt.Println("Error in finding::::", err)
		return utils.NewNotFoundError("Invalid Credentials")
	}
	return nil
}

//Forgot Password
//first api is to check user exist by email. If user exist, trigger OTP to user email and update that otp to database to verify.
//If user doesn't exist, throw a warning that user doesn't exist
func (user *User) CheckUserExist() *utils.RestErr {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	filter := bson.M{"email": user.Email}
	err := userCollection.FindOne(ctx, filter).Decode(&user.Userid)
	defer cancel()
	if err != nil {
		return utils.NewNotFoundError("record doesn't exist")
	}
	opts := options.Update().SetUpsert(true)
	update := bson.M{"$set": bson.M{"otp": user.Otp, "otp_expires": user.OTPExpires}}
	result, updateErr := userCollection.UpdateOne(ctx, filter, update, opts)
	if updateErr != nil {
		return utils.NewNotFoundError("unable to update otp in database")
	}
	if result.ModifiedCount == 0 {
		return utils.NewNotFoundError("unable to update otp in database")
	}
	return nil
}

func (user *User) VerifyOTP() *utils.RestErr {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	filter := bson.M{"userid": user.Userid, "otp": user.Otp}
	err := userCollection.FindOne(ctx, filter).Decode(&user)
	defer cancel()
	if err != nil {
		return utils.NewNotFoundError("unable to verify")
	}
	return nil
}

func (user *User) UpdatePassword(password string) (*mongo.UpdateResult, *utils.RestErr) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	filter := bson.M{"userid": user.Userid}
	opts := options.Update().SetUpsert(true)
	update := bson.M{"$set": bson.M{"password": password}}
	updateResult, err := userCollection.UpdateOne(ctx, filter, update, opts)
	defer cancel()
	if err != nil {
		return nil, utils.NewNotFoundError("unable to update password in database")
	}
	if updateResult.ModifiedCount == 0 {
		return nil, utils.NewNotFoundError("unable to update password in database")
	}
	return updateResult, nil
}
