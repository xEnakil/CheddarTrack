package repository

import (
	"github.com/xenakil/cheddartrack/internal/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction *model.Transaction) error
	GetAllByUser(userID uint) ([]model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(transaction *model.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) GetAllByUser(userID uint) ([]model.Transaction, error) {
    var txns []model.Transaction
    err := r.db.Preload("Category").Where("user_id = ?", userID).Find(&txns).Error
    return txns, err
}
