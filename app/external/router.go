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
	group := r.Group("/invitation")
	group.POST("/expense", expenseController.CreateExpense)
	group.PUT("/expense", expenseController.UpdateExpense)
	group.DELETE("/expense", expenseController.DeleteExpense)
	group.GET("/expense", expenseController.GetExpense)
	group.GET("/expenses", expenseController.GetExpenseList)
	group.GET("/expense/total", expenseController.GetExpenseTotal)
	group.GET("/expense/search", expenseController.GetExpenseSearch)
	group.POST("/expense/csv", expenseController.CreateExpenseByCSV)

	return r
}
