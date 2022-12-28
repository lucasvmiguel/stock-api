package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// product entity
type Product struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `json:"name"`
	Code          uuid.UUID      `gorm:"type:uuid;uniqueIndex:" json:"code"`
	StockQuantity int            `gorm:"default:0" json:"stock_quantity"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

// Gorm hook
// Reference: https://gorm.io/docs/hooks.html
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.Code = uuid.New()
	return
}
