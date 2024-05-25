package external

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"poten-invitation-golang/domain"
	"time"
)

func GetRouter(expenseController domain.ExpenseController) *gin.Engine {
	gin.ForceConsoleColor()
	r := gin.Default()
	r.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"https://www.gardenr.kr", "http://10.0.4.6", "https://10.0.4.6", "https://tikitakaapi.site"},
			AllowMethods: []string{http.MethodPost, http.MethodPut, http.MethodDelete},
			MaxAge:       12 * time.Hour,
		},
	))
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

	return r
}
