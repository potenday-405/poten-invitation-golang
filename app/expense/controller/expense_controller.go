package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"poten-invitation-golang/app/expense/model"
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
	userID := ctx.Request.Header.Get("user_id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token, user id not exist"})
		return
	}
	var expense model.CreateExpense
	if err := ctx.ShouldBind(&expense); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	expense.UserID = userID
	res, err := c.service.CreateExpense(ctx, &expense)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}
