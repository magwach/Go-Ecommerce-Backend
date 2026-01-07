package controllers

import (
	"errors"
	"fmt"
	"go-ecommerce-app/configs"
	functions "go-ecommerce-app/internal/db.functions"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/schema"
	"go-ecommerce-app/pkg/notification"
	"log"
	"time"

	"github.com/google/uuid"
)

type UserContoller struct {
	DB     functions.UserDBFunction
	Auth   helper.Auth
	Config configs.AppConfig
}

func (s UserContoller) SignUp(input dto.UserSignUp) (string, error) {

	if len(input.Password) < 6 {
		return "", errors.New("password is too short")
	}

	hashedPassword, err := helper.HashPassword(input.Password)

	if err != nil {
		return "", errors.New("failed to hash password")
	}

	user, err := s.DB.SignUp(schema.User{
		Email:    input.Email,
		Password: hashedPassword,
		Phone:    input.Phone,
	})

	if err != nil {
		return "", err
	}

	token, err := s.Auth.GenerateJWT(helper.JWTRequirements{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.UserType,
	})

	if err != nil {
		return "", err
	}
	return token, nil
}

func (s UserContoller) FindUserByEmail(email string) (*schema.User, error) {

	user, err := s.DB.FindUserByEmail(email)

	if err != nil {
		return &schema.User{}, err
	}

	return &user, nil
}

func (s UserContoller) Login(email, password string) (string, error) {

	user, err := s.DB.FindUserByEmail(email)

	if err != nil {
		return "", errors.New("user not found")
	}

	valid := helper.CheckPassword(user.Password, password)

	if !valid {
		return "", errors.New("invalid credentials")
	}

	token, err := s.Auth.GenerateJWT(helper.JWTRequirements{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.UserType,
	})

	if err != nil {
		return "", err
	}
	return token, nil
}

func (s UserContoller) GetVerificationCode(u *schema.User) error {

	user, err := s.DB.FindUserById(u.ID)

	if err != nil {
		return errors.New("cannot find user")
	}

	if user.Verified {
		return errors.New("user is already verified")
	}

	token, err := helper.SecureNumericCode(6)

	if err != nil {
		return errors.New("failed to generate token")
	}

	user.Code = token

	user.Expiry = time.Now().Add(time.Minute * 10)

	user, err = s.DB.UpdateUser(user.ID, user)

	if err != nil {
		return err
	}

	notification := notification.InitializeNotification(s.Config)

	message := fmt.Sprintf("Your verification code is: %v", user.Code)

	if err := notification.SendSMS(user.Phone, message); err != nil {
		log.Println("Failed to send SMS:", err)
		return err
	}

	return nil
}

func (s UserContoller) VerifyCode(id uuid.UUID, input dto.UserVerifyCode) error {
	user, err := s.DB.FindUserById(id)

	log.Println("user vERIFIED", user.Verified)

	if err != nil {
		return errors.New("cannot find user")
	}

	if user.Verified {
		return errors.New("user is already verified")
	}

	if time.Now().After(user.Expiry) {
		return errors.New("expired token")
	}

	if input.Code != user.Code {
		return errors.New("invalid token")
	}

	user.Code = ""
	user.Expiry = time.Time{}
	user.Verified = true

	_, err = s.DB.UpdateUser(user.ID, user)

	if err != nil {
		return err
	}

	return nil
}

func (s UserContoller) CreateProfile(id uuid.UUID, input any) error {
	return nil
}

func (s UserContoller) GetProfile(id uuid.UUID) (*schema.User, error) {
	return nil, nil
}

func (s UserContoller) UpdateProfile(id uuid.UUID, input any) error {
	return nil
}

func (s UserContoller) BecomeSeller(id uuid.UUID, input dto.BecomeASeller) (string, error) {

	existingUser, err := s.DB.FindUserById(id)

	if err != nil {
		return "", errors.New("cannot find user")
	}

	if existingUser.UserType == schema.SELLER {
		return "", errors.New("user is already seller")
	}

	existingUser.FirstName = input.FirstName
	existingUser.LastName = input.LastName
	existingUser.Phone = input.PhoneNumber
	existingUser.UserType = schema.SELLER

	seller, err := s.DB.UpdateUser(id, existingUser)

	if err != nil {
		return "", err
	}

	token, err := s.Auth.GenerateJWT(helper.JWTRequirements{
		UserID: seller.ID,
		Email:  seller.Email,
		Role:   seller.UserType,
	})

	if err != nil {
		return "", err
	}

	err = s.DB.CreateBankAccount(schema.BankAccount{
		UserID:      seller.ID,
		BankAccount: input.BankAccountNumber,
		SwiftCode:   input.SwiftCode,
		PaymentType: input.PaymentType,
	})

	if err != nil {
		return "", errors.New("failed to create the bank account")
	}

	return token, nil
}

func (s UserContoller) FindCart(id uuid.UUID) (*schema.Cart, error) {
	return nil, nil
}

func (s UserContoller) CreateCart(input any, u *schema.User) (*schema.Cart, error) {
	return nil, nil
}

func (s UserContoller) CreateOrder(u *schema.User) (int, error) {
	return 0, nil
}

func (s UserContoller) GetOrders(u *schema.User) (*schema.Cart, error) {
	return nil, nil
}

func (s UserContoller) GetOrderById(id, UserId uuid.UUID) (*schema.Cart, error) {
	return nil, nil
}
