package service

import (
    "github.com/xenakil/cheddartrack/internal/model"
    "github.com/xenakil/cheddartrack/internal/repository"
)

type TransactionService interface {
	Create(input model.CreateTransactionInput) error
	GetAll(userID uint) ([]model.TransactionResponse, error)
}

type transactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(r repository.TransactionRepository) TransactionService {
	return &transactionService{r}
}

func (s *transactionService) Create(input model.CreateTransactionInput) error {
	txn := &model.Transaction{
		UserID: input.UserId,
		CategoryID: input.CategoryID,
		Amount: input.Amount,
		Currency: input.Currency,
		Description: input.Description,
		Timestamp: input.Timestamp,
	}
	
	return s.repo.Create(txn)
}

func (s *transactionService) GetAll(userID uint) ([]model.TransactionResponse, error) {
	txns, err := s.repo.GetAllByUser(userID)
	if err != nil {
		return nil, err
	}

	var res []model.TransactionResponse
	for _, t := range txns {
		res = append(res, model.TransactionResponse{
			ID: t.ID,
			Amount: t.Amount,
			Currency: t.Currency,
			Description: t.Description,
			Timestamp: t.Timestamp,
			Category: t.Category.Name,
			Type: t.Category.Type,
		})
	}

	return res, nil
}

