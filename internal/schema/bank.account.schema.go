package schema

import (
	"time"

	"github.com/google/uuid"
)

type BankAccount struct {
	ID          uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID      uuid.UUID  `gorm:"type:uuid;unique;not null"`
	User        *User      `gorm:"foreignKey:UserID"`
	BankAccount string     `gorm:"index;unique;not null"`
	SwiftCode   string
	PaymentType string     `gorm:"type:payment_type;not null"`
	CreatedAt   time.Time  `gorm:"default:current_timestamp"`
	UpdatedAt   time.Time  `gorm:"default:current_timestamp"`
}
