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

	// TODO 점수준다(유저에 요청) 주의사항: 생성점수 + 불참1 참여2 점수

	// 데이터 받아온다
	res, err := s.repo.GetExpenseByEventID(ctx, eventID)
	if err != nil {
		return nil, err
	}

	// return 한다.
	return res, err
}

func (s *expenseService) UpdateExpense(ctx *gin.Context, expense *model.UpdateExpense) (*model.ResponseExpense, error) {
	// 여기서 모델별로 찢기
	event, attendee, err := expense.ToEntity()
	if err != nil {
		return nil, err
	}

	// 이전 정보 (점수 계산에 필요)
	oldEvent, err := s.repo.GetExpenseByEventID(ctx, event.EventID)
	if err != nil {
		return nil, err
	}
	if oldEvent == nil {
		return nil, errors.New("invalid event_id")
	}

	// 수정한다.
	if err = s.repo.GetTransaction(ctx).Transaction(func(tx *gorm.DB) error {
		if affected := tx.Updates(event).RowsAffected; affected == 0 {
			return errors.New("event create failed")
		}
		if affected := tx.Updates(attendee).RowsAffected; affected == 0 {
			return errors.New("attendee create failed")
		}
		return nil
	}); err != nil {
		return nil, err
	}

	// 데이터 받아온다
	newEvent, err := s.repo.GetExpenseByEventID(ctx, event.EventID)
	if err != nil {
		return nil, err
	}

	// TODO 점수준다(유저에 요청) 이전정보에서 참석여부 정보 확인 필요. newEvent.IsAttended - oldEvent.IsAttended = 반영 필요한 점수

	return newEvent, nil
}

func (s *expenseService) DeleteExpense(ctx *gin.Context, expense *model.DeleteExpense) error {
	if err := s.repo.DeleteExpense(ctx, expense.EventID); err != nil {
		return err
	}
	return nil
}

func (s *expenseService) GetExpense(ctx *gin.Context, expense *model.GetExpense) (*model.ResponseExpense, error) {
	res, err := s.repo.GetExpense(ctx, expense.EventID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
