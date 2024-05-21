package repository

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"poten-invitation-golang/app/expense/model"
	"poten-invitation-golang/domain"
	"poten-invitation-golang/util"
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

func (r *expenseRepository) DeleteExpense(ctx *gin.Context, eventID string) error {
	if r.externalDB.Table("event").
		Update("invite_status = ?", "del").
		Where("event_id = ?", eventID).RowsAffected == 0 {
		return errors.New("invalid parameter event_id")
	}
	return nil
}

func (r *expenseRepository) GetExpense(ctx *gin.Context, userID, eventID string) (*model.ResponseExpense, error) {
	var expense model.ResponseExpense
	err := r.externalDB.Select("a.event_id, a.user_id, a.is_invited, a.event_date, b.name, b.relation, b.amount, b.is_attended").
		Table("event a").
		Joins("JOIN attendees b ON a.event_id = b.event_id").
		Where("a.event_id = ?", eventID).
		Where("a.user_id = ?", userID).
		First(&expense).Error
	if err != nil {
		return nil, err
	}
	return &expense, nil
}

func (r *expenseRepository) GetExpenseList(ctx *gin.Context, expense *model.GetExpenseList) ([]*model.ResponseExpense, error) {
	var expenses []*model.ResponseExpense
	db := r.externalDB.Select("a.event_id, a.user_id, a.is_invited, a.event_date, b.name, b.relation, b.amount, b.is_attended").
		Table("event a").
		Joins("JOIN attendees b ON a.event_id = b.event_id").
		Where("a.user_id = ?", expense.UserID)
	if expense.Offset != "" {
		switch expense.OffsetOrderType {
		case 1:
			db.Where("a.event_date < ?", util.StringToTime(expense.Offset).UTC())
		case 2:
			db.Where("a.event_date <= ?", util.StringToTime(expense.Offset).UTC())
		case 3:
			db.Where("a.event_date > ?", util.StringToTime(expense.Offset).UTC())
		case 4:
			db.Where("a.event_date >= ?", util.StringToTime(expense.Offset).UTC())
		default:
			db.Where("a.event_date >= ?", util.StringToTime(expense.Offset).UTC())
		}
	}
	switch expense.Order {
	case "asc":
		db.Order("a.created_at asc")
	case "desc":
		db.Order("a.created_at desc")
	default:
		db.Order("a.created_at desc")
	}
	db.Limit(expense.Limit)
	db.Offset(expense.Limit * (expense.Page - 1))
	err := db.Scan(&expenses).Error
	if err != nil {
		return nil, err
	}
	return expenses, err
}
