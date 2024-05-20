package service

import "poten-invitation-golang/domain"

type expenseService struct {
	repo domain.ExpenseRepository
}

func NewExpenseService(repo domain.ExpenseRepository) domain.ExpenseService {
	return &expenseService{
		repo: repo,
	}
}
