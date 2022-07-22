package service

import (
	"TFP/domain/category"
	"TFP/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var (
	CategoryService CategoryServiceInterface = &categoryService{}
)

type CategoryServiceInterface interface {
	Create(category.Category) (*mongo.InsertOneResult, *utils.RestErr)
	GetAll() ([]primitive.M, *utils.RestErr)
}

type categoryService struct{}

func (s *categoryService) Create(category category.Category) (*mongo.InsertOneResult, *utils.RestErr) {
	if err := category.Validate(); err != nil {
		return nil, err
	}
	category.ID = primitive.NewObjectID()
	category.CategoryID = category.CategoryName
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	insertNumber, insertErr := category.Create()
	if insertErr != nil {
		return nil, insertErr
	}

	return insertNumber, nil
}

func (s *categoryService) GetAll() ([]primitive.M, *utils.RestErr) {
	category := &category.Category{}

	result, err := category.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}
