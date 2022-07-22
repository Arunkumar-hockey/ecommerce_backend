package service

import (
	"TFP/domain/buyer"
	"TFP/utils"
	"fmt"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/smtp"
	"strings"

	//"net/smtp"
	"go.mongodb.org/mongo-driver/mongo"
	//"TFP/helper"
	"time"
)

var (
	UsersService UsersServiceInterface = &usersService{}
)

type UsersServiceInterface interface {
	Create(buyer.User) (*mongo.InsertOneResult, *utils.RestErr)
	GetAll() ([]primitive.M, *utils.RestErr)
	GetByID(string) (*buyer.User, *utils.RestErr)
	Delete(string) (*mongo.DeleteResult, *utils.RestErr)
	Login(buyer.LoginRequest) (*buyer.User, *utils.RestErr)
	CheckUserExist(string) (*string, *utils.RestErr)
	VerifyOTP(buyer.VerifyOTP) (*string, *utils.RestErr)
	UpdatePassword(string, string) (*mongo.UpdateResult, *utils.RestErr)
}

type usersService struct{}

func (s *usersService) Create(user buyer.User) (*mongo.InsertOneResult, *utils.RestErr) {
	if validateErr := user.Validate(); validateErr != nil {
		return nil, validateErr
	}

	uid := xid.New()
	user.Userid = uid.String()
	user.Usertype = "buyer"
	password := utils.HashPasswordMD5(user.Password)
	user.Password = password
	user.Status = 0
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// token, refreshToken, _ := helper.GenerateAlltokens(user.Email, user.FirstName, user.LastName, user.Userid)
	// user.Token = token
	// user.Refresh_Token = refreshToken

	// sEnc := b64.StdEncoding.EncodeToString([]byte(user.Userid))
	// fmt.Println("Encrypt :",sEnc)
	// encryptvalue := sEnc

	// if user.Userid != "" {
	// 	link :=  "https://tfp.psibertechsolutions.com/activate_account?token="+encryptvalue
	// 	from := "tfpsmtp@gmail.com"
	// 	password := "tfpsmtp@123"

	// 	// Receiver email address.
	// 	to := []string{
	// 	  user.Email,
	// 	}

	// 	// smtp server configuration.
	// 	smtpHost := "smtp.gmail.com"
	// 	smtpPort := "587"

	// 	// Message.
	// 	message := []byte(otp)

	// 	// Authentication.
	// 	auth := smtp.PlainAuth("", from, password, smtpHost)

	// 	// Sending email.
	// 	emailsentErr := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	// 	if emailsentErr != nil {
	// 		fmt.Println("Failed to send.......")
	// 	  fmt.Println(emailsentErr)
	// 	  return nil, utils.NewInternalServerError("Unable to send mail")
	// 	}
	// }

	insertNumber, err := user.Save()
	if err != nil {
		return nil, err
	}
	return insertNumber, nil
}

func (s *usersService) GetAll() ([]primitive.M, *utils.RestErr) {
	user := &buyer.User{}
	result, err := user.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) GetByID(userID string) (*buyer.User, *utils.RestErr) {
	result := &buyer.User{Userid: userID}
	if err := result.GetByID(); err != nil {
		fmt.Println("Error in DAO...")
		return nil, err
	}
	return result, nil
}

// func(s *usersService) Update() () {

// }

func (s *usersService) Delete(userID string) (*mongo.DeleteResult, *utils.RestErr) {
	result := buyer.User{Userid: userID}
	deleteID, deleteErr := result.Delete()
	if deleteErr != nil {
		return nil, deleteErr
	}
	return deleteID, nil
}

func (s *usersService) Login(request buyer.LoginRequest) (*buyer.User, *utils.RestErr) {
	result := &buyer.User{
		Email:    request.Email,
		Password: utils.HashPasswordMD5(request.Password),
	}

	if err := result.Login(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) CheckUserExist(emailID string) (*string, *utils.RestErr) {
	result := &buyer.User{Email: emailID}

	otp := strings.TrimSpace(utils.GenerateOTP())

	from := "tfpsmtp@gmail.com"
	password := "tfpsmtp@123"

	// Receiver email address.
	to := []string{
		emailID,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(otp)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	emailSentErr := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if emailSentErr != nil {
		fmt.Println("Error::::", emailSentErr)
		return nil, utils.NewInternalServerError("Unable to send email right now")
	}

	result.Otp = otp
	result.OTPExpires = time.Now().Add(time.Minute * 5)
	if err := result.CheckUserExist(); err != nil {
		return nil, err
	}
	return &result.Userid, nil
}

func (s *usersService) VerifyOTP(request buyer.VerifyOTP) (*string, *utils.RestErr) {
	result := &buyer.User{
		Email: request.Email,
		Otp:   request.OTP,
	}

	expirationOTPTime := result.OTPExpires
	currentTime := time.Now()

	if expirationOTPTime.Before(currentTime) {
		return nil, utils.NewBadRequestError("otp expired")
	}

	if err := result.VerifyOTP(); err != nil {
		return nil, err
	}
	return &result.Userid, nil
}

func (s *usersService) UpdatePassword(userID, password string) (*mongo.UpdateResult, *utils.RestErr) {
	result := &buyer.User{Userid: userID}
	hashedPassword := utils.HashPasswordMD5(password)
	result.Password = hashedPassword
	result.UpdatedAt = time.Now()

	updateResult, err := result.UpdatePassword(password)
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}
