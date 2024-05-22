package service

import (
	"encoding/json"
	"errors"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"log"
	"poten-invitation-golang/app/expense/model"
	"poten-invitation-golang/app/expense/repository"
	"poten-invitation-golang/app/external"
	"poten-invitation-golang/domain"
	"testing"
)

var testService domain.ExpenseService

func testInitializer() {
	if err := godotenv.Load("../../../env/.env"); err != nil {
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
		Expense:    150000,
		Relation:   "친구",
		IsAttended: 1,
		UserID:     "505f4353-6160-4d27-8898-124d2802cb04",
		EventDate:  "202011111700",
		IsInvited:  "inviting",
		Link:       "naver.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)
	indent, _ := json.MarshalIndent(expense, "", "  ")
	t.Log(string(indent))
}

func TestExpenseService_UpdateExpense(t *testing.T) {
	testInitializer()
	expense, err := testService.UpdateExpense(nil, &model.UpdateExpense{
		EventID:    "2d7cfab2-3a66-4775-90dd-71a129c93c8e",
		UserID:     "505f4353-6160-4d27-8898-124d2802cb04",
		IsInvited:  "invited",
		Name:       "박방구",
		EventDate:  "202404100000",
		Expense:    500000,
		Relation:   "사촌",
		IsAttended: 1,
		Link:       "google.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)
	indent, _ := json.MarshalIndent(expense, "", "  ")
	t.Log(string(indent))
}

func TestExpenseService_GetExpense(t *testing.T) {
	testInitializer()
	eventID := "efb47308-02d3-45ec-a70b-9babc2190ca4"
	userID := "505f4353-6160-4d27-8898-124d2802cb04"
	expense, err := testService.GetExpense(nil, &model.GetExpense{
		EventID: eventID,
		UserID:  userID,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.NoError(t, err)
	assert.EqualValues(t, eventID, expense.EventId)
	assert.EqualValues(t, userID, expense.UserID)
	indent, _ := json.MarshalIndent(expense, "", "  ")
	t.Log(string(indent))
}

func TestExpenseService_GetExpenseList(t *testing.T) {
	testInitializer()
	userID := "505f4353-6160-4d27-8898-124d2802cb04"
	list, err := testService.GetExpenseList(nil, &model.GetExpenseList{
		UserID:          userID,
		IsInvited:       "invited",
		Order:           "desc",
		Offset:          "202404100000",
		OffsetOrderType: 1,
		Limit:           2,
		Page:            2,
	})
	if err != nil {
		t.Fatal(err)
	}
	indent, _ := json.MarshalIndent(list, "", "  ")
	t.Log(string(indent))
}

func TestExpenseService_DeleteExpense(t *testing.T) {
	testInitializer()
	expense, err := testService.CreateExpense(nil, &model.CreateExpense{
		Name:       "coen",
		Expense:    150000,
		Relation:   "친구",
		IsAttended: 1,
		UserID:     "505f4353-6160-4d27-8898-124d2802cb04",
		EventDate:  "202011111700",
		IsInvited:  "inviting",
		Link:       "naver.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	if err = testService.DeleteExpense(nil, &model.DeleteExpense{
		EventID: expense.EventId,
		UserID:  expense.UserID,
	}); err != nil {
		t.Fatal(err)
	}
	getExpense, err := testService.GetExpense(nil, &model.GetExpense{EventID: expense.EventId, UserID: expense.UserID})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatal(err)
	}
	indent, _ := json.MarshalIndent(getExpense, "", "  ")
	t.Log(string(indent))
}

func TestExpenseService_GetExpenseTotal(t *testing.T) {
	testInitializer()
	total, err := testService.GetExpenseTotal(nil, &model.GetExpenseTotal{
		UserID:          "505f4353-6160-4d27-8898-124d2802cb04",
		IsInvited:       "invited",
		Offset:          "202405210000",
		OffsetOrderType: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	indent, _ := json.MarshalIndent(total, "", "  ")
	t.Log(string(indent))
}

func TestExpenseService_GetExpenseSearch(t *testing.T) {
	testInitializer()
	search, err := testService.GetExpenseSearch(nil, &model.GetExpenseSearch{
		UserID:    "505f4353-6160-4d27-8898-124d2802cb04",
		IsInvited: "invited",
		Name:      "박방구",
		Order:     "desc",
	})
	if err != nil {
		t.Fatal(err)
	}
	indent, _ := json.MarshalIndent(search, "", "  ")
	t.Log(string(indent))

}
