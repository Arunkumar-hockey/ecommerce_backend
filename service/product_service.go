package service

import (
	"TFP/cache"
	"TFP/domain/products"
	"TFP/utils"
	"fmt"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var (
	ProductsService ProductsServiceInterface = &productsService{}
)

type ProductsServiceInterface interface {
	Create(products.Product) (*mongo.InsertOneResult, *utils.RestErr)
	GetAll() ([]primitive.M, *utils.RestErr)
	GetByID(string) (*products.Product, *utils.RestErr)
	GetProductByCategory(string) ([]primitive.M, *utils.RestErr)
	Update(bool, products.Product) (*mongo.UpdateResult, *utils.RestErr)
	Delete(string) (*mongo.DeleteResult, *utils.RestErr)
	GetRedisData(string) (interface{}, error)
	DeleteRedisData(string) (string, error)
	GetAllRedisData() (interface{}, error)
}

type productsService struct{}

func (s *productsService) Create(product products.Product) (*mongo.InsertOneResult, *utils.RestErr) {
	if validateErr := product.Validate(); validateErr != nil {
		return nil, validateErr
	}

	uid := xid.New()
	product.ProductID = uid.String()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	productID := product.ProductID
	cache.SetData(productID, product)

	insertNumber, err := product.Save()
	if err != nil {
		return nil, err
	}
	return insertNumber, nil
}

func (s *productsService) GetAll() ([]primitive.M, *utils.RestErr) {
	product := &products.Product{}
	result, err := product.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *productsService) GetByID(productID string) (*products.Product, *utils.RestErr) {
	result := &products.Product{ProductID: productID}
	if err := result.GetByID(); err != nil {
		fmt.Println("Error in DAO...")
		return nil, err
	}
	return result, nil
}

func (s *productsService) GetProductByCategory(categoryID string) ([]primitive.M, *utils.RestErr) {
	product := &products.Product{CategoryID: categoryID}
	result, err := product.GetProductByCategory()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *productsService) Update(isPartial bool, product products.Product) (*mongo.UpdateResult, *utils.RestErr) {
	result := &products.Product{SellerID: product.SellerID, ProductID: product.ProductID}

	if isPartial {
		if product.ProductName != "" {
			result.ProductName = product.ProductName
		}
		if product.Price != 0 {
			result.Price = product.Price
		}
		if product.PriceType != "" {
			result.PriceType = product.PriceType
		}
		if product.Description != "" {
			result.Description = product.Description
		}
	} else {
		result.ProductName = product.ProductName
		result.Price = product.Price
		result.PriceType = product.PriceType
		result.Description = product.Description
	}
	product.UpdatedAt = time.Now()

	updatedResult, updateErr := result.Update()
	if updateErr != nil {
		return nil, updateErr
	}
	return updatedResult, nil
}

func (s *productsService) Delete(productID string) (*mongo.DeleteResult, *utils.RestErr) {
	result := &products.Product{ProductID: productID}
	deleteID, deleteErr := result.Delete()
	if deleteErr != nil {
		return nil, deleteErr
	}
	return deleteID, nil
}

func (s *productsService) GetRedisData(productID string) (interface{}, error) {
	data, err := cache.GetData(productID)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *productsService) GetAllRedisData() (interface{}, error) {
	data, err := cache.GetAllData()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *productsService) DeleteRedisData(productID string) (string, error) {
	msg, err := cache.DeleteData(productID)
	if err != nil {
		return "", err
	}
	return msg, nil
}
