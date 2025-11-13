package schema

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name       string    `gorm:"index;unique;not null"`
	Description string    `gorm:"default:null"`
	CategoryID  uuid.UUID  `gorm:"type:uuid;not null"`
	Category    *Category  `gorm:"foreignKey:CategoryID;references:ID"`
	ImageUrl   string    `gorm:"default:null"`
	Price      int       `gorm:"not null"`
	Owner      uuid.UUID `gorm:"type:uuid;not null"`
	User       *User     `gorm:"foreignKey:Owner;references:ID"`
	Stock      uint
	CreatedAt  time.Time `gorm:"default:current_timestamp"`
	UpdatedAt  time.Time `gorm:"default:current_timestamp"`
}
