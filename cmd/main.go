package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"poten-invitation-golang/app/expense/controller"
	"poten-invitation-golang/app/expense/repository"
	"poten-invitation-golang/app/expense/service"
	"poten-invitation-golang/app/external"
	"poten-invitation-golang/util"
)

func main() {
	if err := util.EnvInitializer(); err != nil {
		panic(err)
	}
	db := external.NewDB()
	expenseRepository := repository.NewExpenseRepository(db)
	expenseService := service.NewExpenseService(expenseRepository)
	expenseController := controller.NewExpenseController(expenseService)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
