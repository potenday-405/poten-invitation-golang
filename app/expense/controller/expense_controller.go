package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
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
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token, user id not exist"})
		return
	}
	var expense model.CreateExpense
	if err := ctx.ShouldBind(&expense); err != nil {
		log.Printf("error: parameter error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	expense.UserID = userID
	res, err := c.service.CreateExpense(ctx, &expense)
	if err != nil {
		log.Printf("error: CreateExpense API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) UpdateExpense(ctx *gin.Context) {
	userID := ctx.Request.Header.Get("user_id")
	if userID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token, user id not exist"})
		return
	}
	var expense model.UpdateExpense
	if err := ctx.ShouldBind(&expense); err != nil {
		log.Printf("error: parameter error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	expense.UserID = userID
	res, err := c.service.UpdateExpense(ctx, &expense)
	if err != nil {
		log.Printf("error: UpdateExpense API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) DeleteExpense(ctx *gin.Context) {
	userID := ctx.Request.Header.Get("user_id")
	if userID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token, user id not exist"})
		return
	}
	var expense model.DeleteExpense
	if err := ctx.ShouldBind(&expense); err != nil {
		log.Printf("error: parameter error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	expense.UserID = userID
	if err := c.service.DeleteExpense(ctx, &expense); err != nil {
		log.Printf("error: DeleteExpense API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "deleted successfully",
	})
}

func (c *expenseController) GetExpense(ctx *gin.Context) {
	userID := ctx.Request.Header.Get("user_id")
	if userID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token, user id not exist"})
		return
	}
	var expense model.GetExpense
	if err := ctx.ShouldBind(&expense); err != nil {
		log.Printf("error: parameter error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	expense.UserID = userID
	res, err := c.service.GetExpense(ctx, &expense)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("error: GetExpense API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) GetExpenseList(ctx *gin.Context) {
	userID := ctx.Request.Header.Get("user_id")
	if userID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token, user id not exist"})
		return
	}
	var expense model.GetExpenseList
	if err := ctx.ShouldBind(&expense); err != nil {
		log.Printf("error: parameter error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	expense.UserID = userID
	res, err := c.service.GetExpenseList(ctx, &expense)
	if err != nil {
		log.Printf("error: GetExpenseList API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) GetExpenseTotal(ctx *gin.Context) {
	userID := ctx.Request.Header.Get("user_id")
	if userID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token, user id not exist"})
		return
	}
	var expense model.GetExpenseTotal
	if err := ctx.ShouldBind(&expense); err != nil {
		log.Printf("error: parameter error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	expense.UserID = userID
	res, err := c.service.GetExpenseTotal(ctx, &expense)
	if err != nil {
		log.Printf("error: GetExpenseTotal API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) GetExpenseSearch(ctx *gin.Context) {
	userID := ctx.Request.Header.Get("user_id")
	if userID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token, user id not exist"})
		return
	}
	var expense model.GetExpenseSearch
	if err := ctx.ShouldBind(&expense); err != nil {
		log.Printf("error: parameter error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	expense.UserID = userID
	list, err := c.service.GetExpenseSearch(ctx, &expense)
	if err != nil {
		log.Printf("error: GetExpenseSearch API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
	}
	ctx.JSON(http.StatusOK, list)
}
