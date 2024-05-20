package repository

import (
	"gorm.io/gorm"
	"poten-invitation-golang/domain"
)

type expenseRepository struct {
	externalDB *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) domain.ExpenseRepository {
	return &expenseRepository{
		externalDB: db,
	}
}
