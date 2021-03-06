package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// product entity
type Product struct {
	gorm.Model
	Name          string
	Code          uuid.UUID `gorm:"type:uuid;uniqueIndex:"`
	StockQuantity int       `gorm:"default:0"`
}

// Gorm hook
// Reference: https://gorm.io/docs/hooks.html
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.Code = uuid.New()
	return
}
