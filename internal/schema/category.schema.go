package schema

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name         string    `gorm:"index;unique;not null"`
	Owner        uuid.UUID `gorm:"type:uuid;not null"`
	DisplayOrder int       `gorm:"column:display_order"`
	User         *User     `gorm:"foreignKey:Owner;references:ID"`
	Products     []Product `gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	ImageUrl     string    `gorm:"default:null"`
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp"`
}
