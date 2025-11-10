package schema

import (
	"time"

	"github.com/google/uuid"
)

const (
	SELLER = "seller"
	BUYER  = "buyer"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FirstName string
	LastName  string
	Email     string `gorm:"index;unique;not null"`
	Phone     string `gorm:"unique;not null"`
	Password  string
	Code      string
	Expiry    time.Time `gorm:"default:null"`
	Verified  bool      `gorm:"default:false"`
	UserType  string    `gorm:"default:buyer"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`

	BankAccount BankAccount `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;default:null"`
}
