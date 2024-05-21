package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
	event, attendee, err := expense.ToEntity()
	if err != nil {
		return nil, err
	}

	// uuid 생성
	eventID := uuid.New().String()
	attendeeID := uuid.New().String()
	event.EventID = eventID
	attendee.EventID = eventID
	attendee.AttendeeID = attendeeID

	// 생성한다.
	if err = s.repo.GetTransaction(ctx).Transaction(func(tx *gorm.DB) error {
		if affected := tx.Create(event).RowsAffected; affected == 0 {
			return errors.New("event create failed")
		}
		if affected := tx.Create(attendee).RowsAffected; affected == 0 {
			return errors.New("attendee create failed")
		}
		return nil
	}); err != nil {
		return nil, err
	}

	// 데이터 받아온다
	res, err := s.repo.GetExpenseByEventID(ctx, eventID)
	if err != nil {
		return nil, err
	}

	// TODO 점수준다(유저에 요청).

	// return 한다.
	return res, err
}
