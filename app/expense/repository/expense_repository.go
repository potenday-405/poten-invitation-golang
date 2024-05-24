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
	if err := r.externalDB.Select("a.event_id, a.user_id, a.is_invited, a.event_date, a.link, b.name, b.relation, b.amount, b.is_attended").
		Table("event a").
		Joins("JOIN attendees b ON a.event_id = b.event_id").
		Where("a.event_id = ?", eventID).
		Where("a.invite_status = 'act'").
		Scan(&expense).Error; err != nil {
		return nil, err
	}
	return &expense, nil
}

func (r *expenseRepository) DeleteExpense(ctx *gin.Context, userID, eventID string) error {
	if r.externalDB.Table("event").
		Where("event_id = ?", eventID).
		Where("user_id = ?", userID).
		Update("invite_status", "del").RowsAffected == 0 {
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
		Where("a.invite_status = 'act'").
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
		Where("a.user_id = ?", expense.UserID).
		Where("a.invite_status = 'act'")
	switch expense.IsInvited {
	case "invited":
		db.Where("a.is_invited = ?", 1)
	case "inviting":
		db.Where("a.is_invited = ?", 2)
	}
	eventDate, err := util.StringToTime(expense.Offset)
	if err != nil {
		return nil, err
	}
	if eventDate != nil {
		switch expense.OffsetOrderType {
		case 1:
			db.Where("a.event_date < ?", eventDate)
		case 2:
			db.Where("a.event_date <= ?", eventDate)
		case 3:
			db.Where("a.event_date > ?", eventDate)
		case 4:
			db.Where("a.event_date >= ?", eventDate)
		default:
			db.Where("a.event_date >= ?", eventDate)
		}
	}
	switch expense.Order {
	case "asc":
		db.Order("a.event_date asc")
	case "desc":
		db.Order("a.event_date desc")
	default:
		db.Order("a.event_date desc")
	}
	db.Limit(expense.Limit)
	db.Offset(expense.Limit * (expense.Page - 1))
	err = db.Scan(&expenses).Error
	if err != nil {
		return nil, err
	}
	return expenses, err
}

func (r *expenseRepository) GetExpenseTotal(ctx *gin.Context, expense *model.GetExpenseTotal) (*model.ResponseExpenseTotal, error) {
	var total model.ResponseExpenseTotal
	total.IsInvited = expense.IsInvited
	db := r.externalDB.Select("sum(b.amount) as expense_total, count(1) as expense_count").
		Table("event a").
		Joins("JOIN attendees b ON a.event_id = b.event_id").
		Where("a.user_id = ?", expense.UserID).
		Where("a.invite_status = 'act'")
	eventDate, err := util.StringToTime(expense.Offset)
	if err != nil {
		return nil, err
	}
	if eventDate != nil {
		switch expense.OffsetOrderType {
		case 1:
			db.Where("a.created_at < ?", eventDate)
		case 2:
			db.Where("a.created_at <= ?", eventDate)
		case 3:
			db.Where("a.created_at > ?", eventDate)
		case 4:
			db.Where("a.created_at >= ?", eventDate)
		default:
			db.Where("a.created_at >= ?", eventDate)
		}
	}
	switch expense.IsInvited {
	case "invited":
		db.Where("a.is_invited = 1")
	case "inviting":
		db.Where("a.is_invited = 2")
	default:
		total.IsInvited = "all"
	}
	if err := db.Group("a.user_id").Scan(&total).Error; err != nil {
		return nil, err
	}
	return &total, nil
}

func (r *expenseRepository) GetExpenseSearch(ctx *gin.Context, expense *model.GetExpenseSearch) ([]*model.ResponseExpense, error) {
	var expenses []*model.ResponseExpense
	db := r.externalDB.Select("a.event_id, a.user_id, a.is_invited, a.event_date, b.name, b.relation, b.amount, b.is_attended").
		Table("event a").
		Joins("JOIN attendees b ON a.event_id = b.event_id").
		Where("a.user_id = ?", expense.UserID).
		Where("a.invite_status = 'act'")
	switch expense.IsInvited {
	case "invited":
		db.Where("a.is_invited = 1")
	case "inviting":
		db.Where("a.is_invited = 2")
	}
	db.Where("b.name LIKE ?", "%"+expense.Name+"%")
	switch expense.Order {
	case "asc":
		db.Order("a.event_date asc")
	case "desc":
		db.Order("a.event_date desc")
	default:
		db.Order("a.event_date desc")
	}
	if err := db.Scan(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}
