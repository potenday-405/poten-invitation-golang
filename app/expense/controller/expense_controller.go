package controller

import (
	"github.com/gin-gonic/gin"
	"poten-invitation-golang/domain"
)

type expenseController struct {
	service domain.ExpenseService
}

func NewExpenseController(service domain.ExpenseService) domain.ExpenseController {
	return &expenseController{
		service: service,
	}
}

func (c *expenseController) CreateExpense(ctx *gin.Context) {
	ctx.Bind()
}
