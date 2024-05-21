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
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	expense.UserID = userID
	res, err := c.service.CreateExpense(ctx, &expense)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) UpdateExpense(ctx *gin.Context) {
	userID := ctx.Request.Header.Get("user_id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token, user id not exist"})
		return
	}
	var expense model.UpdateExpense
	if err := ctx.ShouldBind(&expense); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	expense.UserID = userID
	res, err := c.service.UpdateExpense(ctx, &expense)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) DeleteExpense(ctx *gin.Context) {
	userID := ctx.Request.Header.Get("user_id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token, user id not exist"})
		return
	}
	var expense model.DeleteExpense
	if err := ctx.ShouldBind(&expense); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := c.service.DeleteExpense(ctx, &expense); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "deleted successfully",
	})
}

func (c *expenseController) GetExpense(ctx *gin.Context) {
	userID := ctx.Request.Header.Get("user_id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token, user id not exist"})
		return
	}
	var expense model.GetExpense
	if err := ctx.ShouldBind(&expense); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	expense.UserID = userID
	res, err := c.service.GetExpense(ctx, &expense)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) GetExpenseList(ctx *gin.Context) {
	userID := ctx.Request.Header.Get("user_id")
	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token, user id not exist"})
		return
	}
	var expense model.GetExpenseList
	if err := ctx.ShouldBind(&expense); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	expense.UserID = userID
	res, err := c.service.GetExpenseList(ctx, &expense)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	ctx.JSON(http.StatusOK, res)
}
