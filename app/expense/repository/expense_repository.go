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

func (r *expenseRepository) GetTransaction(ctx *gin.Context) *gorm.DB {
	return r.externalDB
}

func (r *expenseRepository) CreateEvent(ctx *gin.Context, event *model.Event) error {
	if affected := r.externalDB.Create(event).RowsAffected; affected == 0 {
		return errors.New("event create failed")
	}
	return nil
}

func (r *expenseRepository) CreateAttendee(ctx *gin.Context, attendee *model.Attendees) error {
	if affected := r.externalDB.Create(attendee).RowsAffected; affected == 0 {
		return errors.New("attendee create failed")
	}
	return nil
}

func (r *expenseRepository) GetExpenseByEventID(ctx *gin.Context, eventID string) (*model.ResponseExpense, error) {
	var expense model.ResponseExpense
	if err := r.externalDB.Select("a.event_id, a.user_id, a.is_invited, a.event_date, b.name, b.relation, b.amount, b.is_attended").
		Table("event a").
		Joins("JOIN attendees b ON a.event_id = b.event_id").
		Where("a.event_id = ?", eventID).
		Scan(&expense).Error; err != nil {
		return nil, err
	}
	return &expense, nil
}
