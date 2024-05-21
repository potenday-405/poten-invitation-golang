package service

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"poten-invitation-golang/app/expense/model"
	"poten-invitation-golang/app/expense/repository"
	"poten-invitation-golang/app/external"
	"poten-invitation-golang/domain"
	"testing"
)

var testService domain.ExpenseService

func testInitializer() {
	if err := godotenv.Load("../../../.env"); err != nil {
		log.Fatal(err)
	}
	db := external.NewDB()
	expenseRepository := repository.NewExpenseRepository(db)
	testService = NewExpenseService(expenseRepository)
}

func TestExpenseService_CreateExpense(t *testing.T) {
	testInitializer()
	expense, err := testService.CreateExpense(nil, &model.CreateExpense{
		Name:       "coen",
		Expense:    50000,
		Relation:   "친구",
		IsAttended: 1,
		UserID:     "505f4353-6160-4d27-8898-124d2802cb04",
		EventDate:  "2020-11-11 17:00:00",
		IsInvited:  "invited",
		Link:       "naver.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	indent, _ := json.MarshalIndent(expense, "", "  ")
	t.Log(string(indent))
}
