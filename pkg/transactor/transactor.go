package transactor

import (
	"context"

	"gorm.io/gorm"
)

type key int

const (
	dbTransaction key = iota
)

// Transactor is a struct that can create transactions
type Transactor struct {
	db *gorm.DB
}

// New creates a new Transactor
func New(DB *gorm.DB) *Transactor {
	return &Transactor{db: DB}
}

// Begin creates a new transaction
func (t *Transactor) Begin(ctx context.Context) context.Context {
	return context.WithValue(ctx, dbTransaction, t.db.Begin())
}

// Commit commits the transaction
func (t *Transactor) Commit(ctx context.Context) {
	transaction, ok := ctx.Value(dbTransaction).(*gorm.DB)
	if !ok {
		return
	}

	transaction.Commit()
}

// Rollback rolls back the transaction
func (t *Transactor) Rollback(ctx context.Context) {
	transaction, ok := ctx.Value(dbTransaction).(*gorm.DB)
	if !ok {
		return
	}

	transaction.Rollback()
}

// Gets the transaction from the context
func DBTransaction(ctx context.Context) *gorm.DB {
	transaction, ok := ctx.Value(dbTransaction).(*gorm.DB)
	if !ok {
		return nil
	}

	return transaction
}
