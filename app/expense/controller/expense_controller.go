package controller

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"poten-invitation-golang/app/expense/model"
	"poten-invitation-golang/domain"
	"strconv"
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
	var expense model.CreateExpense
	buf := make([]byte, 1024)
	num, _ := ctx.Request.Body.Read(buf)
	reqBody := string(buf[0:num])
	log.Println("request body: " + string(reqBody))
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(reqBody))) // Write body back
	if err := ctx.ShouldBind(&expense); err != nil {
		log.Printf("error: parameter error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}
	expense.UserID = ctx.Request.Header.Get("user_id")
	log.Printf("Create Expense user_id Log: %v", expense.UserID)
	if expense.UserID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": "Invalid token, user id not exist"})
		return
	}
	log.Printf("Create Expense Parameter Log: %v", expense)
	res, err := c.service.CreateExpense(ctx, &expense)
	if err != nil {
		log.Printf("error: CreateExpense API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) UpdateExpense(ctx *gin.Context) {
	var expense model.UpdateExpense
	expense.UserID = ctx.Request.Header.Get("user_id")
	log.Printf("Create Expense user_id Log: %v", expense.UserID)
	if expense.UserID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": "Invalid token, user id not exist"})
		return
	}
	if err := ctx.ShouldBind(&expense); err != nil {
		log.Printf("error: parameter error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}
	log.Printf("Create Expense Parameter Log: %v", expense)
	res, err := c.service.UpdateExpense(ctx, &expense)
	if err != nil {
		log.Printf("error: UpdateExpense API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) DeleteExpense(ctx *gin.Context) {
	var expense model.DeleteExpense
	expense.UserID = ctx.Request.Header.Get("user_id")
	log.Printf("Create Expense user_id Log: %v", expense.UserID)
	if expense.UserID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": "Invalid token, user id not exist"})
		return
	}
	expense.EventID = ctx.Query("event_id")
	if expense.EventID == "" {
		log.Printf("error: parameter event_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": errors.New("parameter event_id not exist")})
		return
	}
	log.Printf("Create Expense Parameter Log: %v", expense)
	if err := c.service.DeleteExpense(ctx, &expense); err != nil {
		log.Printf("error: DeleteExpense API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "deleted successfully",
	})
}

func (c *expenseController) GetExpense(ctx *gin.Context) {
	var expense model.GetExpense
	expense.UserID = ctx.Request.Header.Get("user_id")
	log.Printf("Create Expense user_id Log: %v", expense.UserID)
	if expense.UserID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": "Invalid token, user id not exist"})
		return
	}
	expense.EventID = ctx.Query("event_id")
	if expense.EventID == "" {
		log.Printf("error: event_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": errors.New("event_id not exist")})
		return
	}
	log.Printf("Create Expense Parameter Log: %v", expense)
	res, err := c.service.GetExpense(ctx, &expense)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("error: GetExpense API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": err})
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) GetExpenseList(ctx *gin.Context) {
	var expense model.GetExpenseList
	expense.UserID = ctx.Request.Header.Get("user_id")
	log.Printf("Create Expense user_id Log: %v", expense.UserID)
	if expense.UserID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": "Invalid token, user id not exist"})
		return
	}
	offsetOrderType, _ := strconv.Atoi(ctx.Query("offset_order_type"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	page, _ := strconv.Atoi(ctx.Query("page"))
	expense.IsInvited = ctx.Query("is_invited")
	expense.Offset = ctx.Query("offset")
	expense.OffsetOrderType = int8(offsetOrderType)
	expense.Order = ctx.Query("order")
	expense.Limit = limit
	expense.Page = page
	log.Printf("Create Expense Parameter Log: %v", expense)
	res, err := c.service.GetExpenseList(ctx, &expense)
	if err != nil {
		log.Printf("error: GetExpenseList API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": err})
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) GetExpenseTotal(ctx *gin.Context) {
	var expense model.GetExpenseTotal
	expense.UserID = ctx.Request.Header.Get("user_id")
	log.Printf("Create Expense user_id Log: %v", expense.UserID)
	if expense.UserID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": "Invalid token, user id not exist"})
		return
	}
	if err := ctx.ShouldBind(&expense); err != nil {
		log.Printf("error: parameter error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}
	offsetOrderType, _ := strconv.Atoi(ctx.Query("offset_order_type"))
	expense.IsInvited = ctx.Query("is_invited")
	expense.Offset = ctx.Query("offset")
	expense.OffsetOrderType = int8(offsetOrderType)
	log.Printf("Create Expense Parameter Log: %v", expense)
	res, err := c.service.GetExpenseTotal(ctx, &expense)
	if err != nil {
		log.Printf("error: GetExpenseTotal API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": err})
	}
	ctx.JSON(http.StatusOK, res)
}

func (c *expenseController) GetExpenseSearch(ctx *gin.Context) {
	var expense model.GetExpenseSearch
	expense.UserID = ctx.Request.Header.Get("user_id")
	log.Printf("Create Expense user_id Log: %v", expense.UserID)
	if expense.UserID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": "Invalid token, user id not exist"})
		return
	}
	expense.IsInvited = ctx.Query("is_invited")
	expense.Name = ctx.Query("name")
	expense.Order = ctx.Query("order")
	log.Printf("Create Expense Parameter Log: %v", expense)
	list, err := c.service.GetExpenseSearch(ctx, &expense)
	if err != nil {
		log.Printf("error: GetExpenseSearch API error: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": err})
	}
	ctx.JSON(http.StatusOK, list)
}

func (c *expenseController) CreateExpenseByCSV(ctx *gin.Context) {
	var expense model.CreateExpenseByCSV

	// TODO Remove this Logic
	// TODO 이 부분이 잘 안먹히네?
	//buf := make([]byte, 1024)
	//num, _ := ctx.Request.Body.Read(buf)
	//reqBody := string(buf[0:num])
	//log.Println("request body: " + string(reqBody))
	//ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(reqBody))) // Write body back
	// TODO Remove this Logic

	expense.UserID = ctx.Request.Header.Get("user_id")
	log.Printf("Create Expense user_id Log: %v", expense.UserID)
	if expense.UserID == "" {
		log.Println("error: user_id not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": "Invalid token, user id not exist"})
		return
	}
	if err := ctx.ShouldBind(&expense); err != nil {
		log.Println("error: Error in JSON binding", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": "Error in JSON binding" + err.Error()})
		return
	}
	form, _ := ctx.MultipartForm()
	files := form.File
	for s, headers := range files {
		log.Println("files test!!!")
		log.Println(s)
		log.Println(headers)
	}
	if expense.File == nil {
		log.Println("error: file is not exist")
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": "file is not exist"})
		return
	} else {
		log.Println("file upload test: ", expense.File.Filename)
		log.Println("file upload test: ", expense.File.Header)
		log.Println("file upload test: ", expense.File.Size)
	}
	err := c.service.CreateExpenseByCSV(ctx, &expense)
	if err != nil {
		log.Println("error: CreateExpenseByCSV service failed", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, "Excel Uploaded Successfully")
}
