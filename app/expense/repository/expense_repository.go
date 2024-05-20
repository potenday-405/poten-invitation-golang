package repository

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"poten-invitation-golang/app/expense/model"
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

func (r *expenseRepository) CreateExpense(ctx *gin.Context, event *model.Event, attendees *model.Attendees) (*model.ResponseExpense, error) {
	if err := r.externalDB.Transaction(func(tx *gorm.DB) error {
		eventAffected := tx.Create(event).RowsAffected
		attendeesAffected := tx.Create(attendees).RowsAffected
		if eventAffected == 0 || attendeesAffected == 0 {
			return errors.New("insert Failed")
		}
		return nil
	}); err != nil {
		return nil, err
	}
	
}
