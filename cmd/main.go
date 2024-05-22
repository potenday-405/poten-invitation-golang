package main

import (
	"log"
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

	if err := external.GetRouter(expenseController).Run(); err != nil {
		log.Fatalln(err)
	}

}
