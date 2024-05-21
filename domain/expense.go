package domain

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"poten-invitation-golang/app/expense/model"
)

type ExpenseRepository interface {
	GetTransaction(ctx *gin.Context) *gorm.DB
	CreateEvent(ctx *gin.Context, event *model.Event) error
	CreateAttendee(ctx *gin.Context, attendee *model.Attendees) error
	GetExpenseByEventID(ctx *gin.Context, eventID string) (*model.ResponseExpense, error)
}

type ExpenseService interface {
	CreateExpense(ctx *gin.Context, expense *model.CreateExpense) (*model.ResponseExpense, error)
}

type ExpenseController interface {
	CreateExpense(ctx *gin.Context)
}
