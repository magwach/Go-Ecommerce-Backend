package controllers

import (
	"go-ecommerce-app/internal/dormain"
	"go-ecommerce-app/internal/dto"
	"log"

	"github.com/google/uuid"
)

type UserContoller struct{}

func (s UserContoller) SignUp(input dto.UserSignUp) (string, error) {

	log.Println(input)

	return "test-token", nil
}

func (s UserContoller) FindUserByEmail(email string) (*dormain.User, error) {
	return nil, nil
}

func (s UserContoller) Login(input any) (string, error) {
	return "", nil
}

func (s UserContoller) GetVerificationCode(u *dormain.User) (int, error) {
	return 0, nil
}

func (s UserContoller) VerifyCode(u *dormain.User) error {
	return nil
}

func (s UserContoller) CreateProfile(id uuid.UUID, input any) error {
	return nil
}

func (s UserContoller) GetProfile(id uuid.UUID) (*dormain.User, error) {
	return nil, nil
}

func (s UserContoller) UpdateProfile(id uuid.UUID, input any) error {
	return nil
}

func (s UserContoller) BecomeSeller(id uuid.UUID, input any) (string, error) {
	return "", nil
}

func (s UserContoller) FindCart(id uuid.UUID) (*dormain.Cart, error) {
	return nil, nil
}

func (s UserContoller) CreateCart(input any, u *dormain.User) (*dormain.Cart, error) {
	return nil, nil
}

func (s UserContoller) CreateOrder(u *dormain.User) (int, error) {
	return 0, nil
}

func (s UserContoller) GetOrders(u *dormain.User) (*dormain.Cart, error) {
	return nil, nil
}

func (s UserContoller) GetOrderById(id, UserId uuid.UUID) (*dormain.Cart, error) {
	return nil, nil
}
