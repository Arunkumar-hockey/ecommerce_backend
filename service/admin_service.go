package service

import (
	"TFP/domain/admin"
	"TFP/utils"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var (
	AdminService AdminServiceInterface = &adminService{}
)

type AdminServiceInterface interface {
	Create(admin.Admin) (*mongo.InsertOneResult, *utils.RestErr)
	GetAll() ([]primitive.M, *utils.RestErr)
	GetByID(string) (*admin.Admin, *utils.RestErr)
	Delete(string) (*mongo.DeleteResult, *utils.RestErr)
	Login(admin.LoginRequest) (*admin.Admin, *utils.RestErr)
}

type adminService struct{}

func (s *adminService) Create(admin admin.Admin) (*mongo.InsertOneResult, *utils.RestErr) {
	if validateErr := admin.Validate(); validateErr != nil {
		return nil, validateErr
	}

	uid := xid.New()
	admin.AdminID = uid.String()
	//admin.Usertype = "buyer"
	password := utils.HashPasswordMD5(admin.Password)
	admin.Password = password
	admin.Status = 0
	admin.CreatedAt = time.Now()
	admin.UpdatedAt = time.Now()

	insertNumber, err := admin.Save()
	if err != nil {
		return nil, err
	}
	return insertNumber, nil
}

func (s *adminService) GetAll() ([]primitive.M, *utils.RestErr) {
	admin := &admin.Admin{}
	result, err := admin.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *adminService) GetByID(adminID string) (*admin.Admin, *utils.RestErr) {
	result := &admin.Admin{AdminID: adminID}
	if err := result.GetByID(); err != nil {
		return nil, err
	}
	return result, nil
}

// func(s *adminService) Update() () {

// }

func (s *adminService) Delete(adminID string) (*mongo.DeleteResult, *utils.RestErr) {
	result := &admin.Admin{AdminID: adminID}
	deleteID, deleteErr := result.Delete()
	if deleteErr != nil {
		return nil, deleteErr
	}
	return deleteID, nil
}

func (s *adminService) Login(request admin.LoginRequest) (*admin.Admin, *utils.RestErr) {
	result := &admin.Admin{
		Email:    request.Email,
		Password: utils.HashPasswordMD5(request.Password),
	}

	if err := result.Login(); err != nil {
		return nil, err
	}
	return result, nil

}
