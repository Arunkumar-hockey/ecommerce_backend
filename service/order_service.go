package service

import (
	"TFP/domain/order"
	"TFP/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	OrderService OrderServiceInterface = &orderService{}
)

type OrderServiceInterface interface {
	PlaceOrder(string) (*mongo.InsertOneResult, *utils.RestErr)
	GetOrder(string) ([]bson.M, *utils.RestErr)
	GetAllOrder() ([]bson.M, *utils.RestErr)
}

type orderService struct{}

func (s *orderService) PlaceOrder(userID string) (*mongo.InsertOneResult, *utils.RestErr) {
	order := &order.Order{}

	result, err := order.PlaceOrder(userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *orderService) GetOrder(userID string) ([]bson.M, *utils.RestErr) {
	order := &order.Order{BuyerID: userID}
	result, err := order.GetOrder()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *orderService) GetAllOrder() ([]bson.M, *utils.RestErr) {
	order := &order.Order{}
	result, err := order.GetAllOrder()
	if err != nil {
		return nil, err
	}
	return result, nil
}
