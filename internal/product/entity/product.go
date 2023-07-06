// entity package is a package that contains the entities of the domain product
package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// product entity
type Product struct {
	ID            int `gorm:"primaryKey"`
	Name          string
	Code          uuid.UUID `gorm:"uniqueIndex"`
	StockQuantity int       `gorm:"default:0"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

// Gorm hook
// Reference: https://gorm.io/docs/hooks.html
func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.Code = uuid.New()
	return
}
