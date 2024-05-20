package domain

import (
	"github.com/gin-gonic/gin"
	"poten-invitation-golang/app/expense/model"
)

type ExpenseRepository interface {
	CreateExpense(ctx *gin.Context, event *model.Event, attendees *model.Attendees) (*model.ResponseExpense, error)
}

type ExpenseService interface {
	CreateExpense(ctx *gin.Context, expense *model.CreateExpense) (*model.ResponseExpense, error)
}

type ExpenseController interface {
	CreateExpense(ctx *gin.Context)
}
