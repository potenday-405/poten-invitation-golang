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
	DeleteExpense(ctx *gin.Context, userID, eventID string) error
	GetExpense(ctx *gin.Context, userID, eventID string) (*model.ResponseExpense, error)
	GetExpenseList(ctx *gin.Context, expense *model.GetExpenseList) ([]*model.ResponseExpense, error)
	GetExpenseTotal(ctx *gin.Context, expense *model.GetExpenseTotal) (*model.ResponseExpenseTotal, error)
	GetExpenseSearch(ctx *gin.Context, expense *model.GetExpenseSearch) ([]*model.ResponseExpense, error)
}

type ExpenseService interface {
	CreateExpense(ctx *gin.Context, expense *model.CreateExpense) (*model.ResponseExpense, error)
	UpdateExpense(ctx *gin.Context, expense *model.UpdateExpense) (*model.ResponseExpense, error)
	DeleteExpense(ctx *gin.Context, expense *model.DeleteExpense) error
	GetExpense(ctx *gin.Context, expense *model.GetExpense) (*model.ResponseExpense, error)
	GetExpenseList(ctx *gin.Context, expense *model.GetExpenseList) ([]*model.ResponseExpense, error)
	GetExpenseTotal(ctx *gin.Context, expense *model.GetExpenseTotal) (*model.ResponseExpenseTotal, error)
	GetExpenseSearch(ctx *gin.Context, expense *model.GetExpenseSearch) ([]*model.ResponseExpense, error)
	CreateExpenseByCSV(ctx *gin.Context, expense *model.CreateExpenseByCSV) error
}

type ExpenseController interface {
	CreateExpense(ctx *gin.Context)
	UpdateExpense(ctx *gin.Context)
	DeleteExpense(ctx *gin.Context)
	GetExpense(ctx *gin.Context)
	GetExpenseList(ctx *gin.Context)
	GetExpenseTotal(ctx *gin.Context)
	GetExpenseSearch(ctx *gin.Context)
	CreateExpenseByCSV(ctx *gin.Context)
}
