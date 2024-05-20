package service

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"poten-invitation-golang/app/expense/model"
	"poten-invitation-golang/domain"
)

type expenseService struct {
	repo domain.ExpenseRepository
}

func NewExpenseService(repo domain.ExpenseRepository) domain.ExpenseService {
	return &expenseService{
		repo: repo,
	}
}

func (s *expenseService) CreateExpense(ctx *gin.Context, expense *model.CreateExpense) (*model.ResponseExpense, error) {
	// 여기서 모델별로 찢기
	event, attendees, err := expense.ToEntity()
	if err != nil {
		return nil, err
	}

	// uuid 생성
	eventID := uuid.New().String()
	attendeeID := uuid.New().String()
	event.EventID = eventID
	attendees.EventID = eventID
	attendees.AtendeeID = attendeeID

	// 생성한다.
	s.repo.CreateExpense(ctx, event, attendees)

	// 점수준다(유저에 요청).

	// return 한다.
}
