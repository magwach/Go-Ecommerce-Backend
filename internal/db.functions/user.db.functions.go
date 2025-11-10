package functions

import (
	"errors"
	"go-ecommerce-app/internal/schema"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserDBFunction interface {
	SignUp(user schema.User) (schema.User, error)
	FindUserByEmail(email string) (schema.User, error)
	FindUserById(id uuid.UUID) (schema.User, error)
	UpdateUser(id uuid.UUID, user schema.User) (schema.User, error)
	CreateBankAccount(details schema.BankAccount) error
}

type userDBFunction struct {
	db *gorm.DB
}

func InitializeUserDBFunction(db *gorm.DB) UserDBFunction {
	return userDBFunction{
		db: db,
	}
}

func (r userDBFunction) SignUp(user schema.User) (schema.User, error) {

	err := r.db.Create(&user).Error

	if err != nil {
		log.Printf("error in creating user: %v", err)
		return schema.User{}, errors.New("error while creating user")
	}

	return user, nil
}

func (r userDBFunction) FindUserByEmail(email string) (schema.User, error) {

	var user schema.User

	err := r.db.First(&user, "email = ?", email).Error

	if err != nil {
		log.Printf("error in finding user: %v", err)
		return schema.User{}, errors.New("user doesn't exist")
	}

	return user, nil
}

func (r userDBFunction) FindUserById(id uuid.UUID) (schema.User, error) {
	var user schema.User

	err := r.db.First(&user, "id = ?", id).Error

	if err != nil {
		log.Printf("error in finding user: %v", err)
		return schema.User{}, errors.New("user doesn't exist")
	}

	return user, nil
}

func (r userDBFunction) UpdateUser(id uuid.UUID, updated schema.User) (schema.User, error) {
	var user schema.User

	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return schema.User{}, errors.New("user not found")
		}
		return schema.User{}, err
	}

	user.FirstName = updated.FirstName
	user.LastName = updated.LastName
	user.Email = updated.Email
	user.Phone = updated.Phone
	user.Password = updated.Password
	user.Code = updated.Code
	user.Expiry = updated.Expiry
	user.Verified = updated.Verified
	user.UserType = updated.UserType
	user.BankAccount = updated.BankAccount

	if err := r.db.Model(&user).
		Select("*").
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(user).Error; err != nil {
		return schema.User{}, errors.New("failed to update user")
	}

	return user, nil
}

func (r userDBFunction) CreateBankAccount(details schema.BankAccount) error {
	return r.db.Create(&details).Error
}
