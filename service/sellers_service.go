package service

import (
	"TFP/domain/seller"
	"TFP/utils"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var (
	SellersService SellersServiceInterface = &sellersService{}
)

type SellersServiceInterface interface {
	Create(seller.Seller) (*mongo.InsertOneResult, *utils.RestErr)
	GetAll() ([]primitive.M, *utils.RestErr)
	GetByID(string) (*seller.Seller, *utils.RestErr)
	Delete(string) (*mongo.DeleteResult, *utils.RestErr)
	Login(seller.LoginRequest) (*seller.Seller, *utils.RestErr)
}

type sellersService struct{}

func (s *sellersService) Create(seller seller.Seller) (*mongo.InsertOneResult, *utils.RestErr) {
	if validateErr := seller.Validate(); validateErr != nil {
		return nil, validateErr
	}

	uid := xid.New()
	seller.Sellerid = uid.String()
	seller.Usertype = "seller"
	password := utils.HashPasswordMD5(seller.Password)
	seller.Password = password
	seller.Status = 0
	seller.CreatedAt = time.Now()
	seller.UpdatedAt = time.Now()

	insertNumber, err := seller.Save()
	if err != nil {
		return nil, err
	}
	return insertNumber, nil
}

func (s *sellersService) GetAll() ([]primitive.M, *utils.RestErr) {
	seller := &seller.Seller{}
	result, err := seller.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *sellersService) GetByID(sellerID string) (*seller.Seller, *utils.RestErr) {
	result := &seller.Seller{Sellerid: sellerID}
	if err := result.GetByID(); err != nil {
		return nil, err
	}
	return result, nil
}

// func(s *usersService) Update() () {

// }

func (s *sellersService) Delete(sellerID string) (*mongo.DeleteResult, *utils.RestErr) {
	result := &seller.Seller{Sellerid: sellerID}
	deleteID, deleteErr := result.Delete()
	if deleteErr != nil {
		return nil, deleteErr
	}
	return deleteID, nil
}

func (s *sellersService) Login(request seller.LoginRequest) (*seller.Seller, *utils.RestErr) {
	result := &seller.Seller{
		Email:    request.Email,
		Password: utils.HashPasswordMD5(request.Password),
	}

	if err := result.Login(); err != nil {
		return nil, err
	}
	return result, nil

}
