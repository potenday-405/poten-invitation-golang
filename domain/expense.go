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
	DeleteExpense(ctx *gin.Context, eventID string) error
}

type ExpenseService interface {
	CreateExpense(ctx *gin.Context, expense *model.CreateExpense) (*model.ResponseExpense, error)
	UpdateExpense(ctx *gin.Context, expense *model.UpdateExpense) (*model.ResponseExpense, error)
	DeleteExpense(ctx *gin.Context, expense *model.DeleteExpense) error
}

type ExpenseController interface {
	CreateExpense(ctx *gin.Context)
	UpdateExpense(ctx *gin.Context)
	DeleteExpense(ctx *gin.Context)
}
