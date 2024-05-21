package external

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"poten-invitation-golang/domain"
)

func GetRouter(expenseController domain.ExpenseController) *gin.Engine {
	gin.ForceConsoleColor()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Group("/invitation")
	r.POST("/expense", expenseController.CreateExpense)
	r.PUT("/expense", expenseController.UpdateExpense)
	r.DELETE("/expense", expenseController.DeleteExpense)
	r.GET("/expense", expenseController.GetExpense)
	r.GET("/expenses", expenseController.GetExpenseList)
	r.GET("/expense/total", expenseController.GetExpenseTotal)
	r.GET("/expense/search", expenseController.GetExpenseSearch)

	return r
}
