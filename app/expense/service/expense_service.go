package service

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"poten-invitation-golang/app/expense/model"
	"poten-invitation-golang/domain"
	"poten-invitation-golang/util"
	"strconv"
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
		eventResult := tx.Create(event)
		if eventResult.Error != nil {
			return eventResult.Error
		}
		if eventResult.RowsAffected == 0 {
			return errors.New("event create failed")
		}
		attendeeResult := tx.Create(attendee)
		if attendeeResult.Error != nil {
			return attendeeResult.Error
		}
		if attendeeResult.RowsAffected == 0 {
			return errors.New("event create failed")
		}
		return nil
	}); err != nil {
		return nil, err
	}

	url := "http://" + os.Getenv("USER_SERVER") + "/user/score"
	body, err := json.Marshal(util.UserScore{Method: http.MethodPost, IsAttended: int(expense.IsAttended), InvitationType: "Wedding"})
	if err != nil {
		return nil, err
	}
	client, err := util.RestClient(http.MethodPost, url, expense.UserID, body)
	if err != nil && client != 200 {
		return nil, errors.New("user score request failed")
	}

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
		// TODO UpdateColumn으로 변경 필요. 현재는 ID도 update하기에 문제가 생길 수 있다
		eventResult := tx.Table("event a").Where("a.event_id = ?", event.EventID).Where("a.user_id = ?", event.UserID).Updates(&event)
		if eventResult.Error != nil {
			return eventResult.Error
		}
		if eventResult.RowsAffected == 0 {
			return errors.New("event update failed")
		}
		attendeeResult := tx.Where("event_id = ?", event.EventID).Updates(attendee)
		if attendeeResult.Error != nil {
			return attendeeResult.Error
		}
		if attendeeResult.RowsAffected == 0 {
			return errors.New("attendee update failed")
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
	url := "http://" + os.Getenv("USER_SERVER") + "/user/score"
	body, err := json.Marshal(util.UserScore{Method: http.MethodPost, IsAttended: int(expense.IsAttended), InvitationType: "Wedding"})
	if err != nil {
		return nil, err
	}
	client, err := util.RestClient(http.MethodPut, url, expense.UserID, body)
	if err != nil && client != 200 {
		return nil, errors.New("user score request failed")
	}

	return newEvent, nil
}

func (s *expenseService) DeleteExpense(ctx *gin.Context, expense *model.DeleteExpense) error {
	// 이전 정보 (점수 계산에 필요)
	oldEvent, err := s.repo.GetExpenseByEventID(ctx, expense.EventID)
	if err != nil {
		return err
	}
	if err := s.repo.DeleteExpense(ctx, expense.UserID, expense.EventID); err != nil {
		return err
	}
	// TODO 점수 삭제(유저에 요청) 이전 정보에서 참석여부 정보 확인 + 생성 점수 삭제 요청
	url := "http://" + os.Getenv("USER_SERVER") + "/user/score"
	body, err := json.Marshal(util.UserScore{Method: http.MethodPost, IsAttended: int(oldEvent.IsAttended), InvitationType: "Wedding"})
	if err != nil {
		return err
	}
	client, err := util.RestClient(http.MethodDelete, url, expense.UserID, body)
	if err != nil && client != 200 {
		return errors.New("user score request failed")
	}
	return nil
}

func (s *expenseService) GetExpense(ctx *gin.Context, expense *model.GetExpense) (*model.ResponseExpense, error) {
	res, err := s.repo.GetExpense(ctx, expense.UserID, expense.EventID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *expenseService) GetExpenseList(ctx *gin.Context, expense *model.GetExpenseList) ([]*model.ResponseExpense, error) {
	if expense.Limit > 100 {
		expense.Limit = 100
	}
	if expense.Limit == 0 {
		expense.Limit = 10
	}
	if expense.Page == 0 {
		expense.Page = 1
	}
	list, err := s.repo.GetExpenseList(ctx, expense)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *expenseService) GetExpenseTotal(ctx *gin.Context, expense *model.GetExpenseTotal) (*model.ResponseExpenseTotal, error) {
	total, err := s.repo.GetExpenseTotal(ctx, expense)
	if err != nil {
		return nil, err
	}
	return total, nil
}

func (s *expenseService) GetExpenseSearch(ctx *gin.Context, expense *model.GetExpenseSearch) ([]*model.ResponseExpense, error) {
	list, err := s.repo.GetExpenseSearch(ctx, expense)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *expenseService) CreateExpenseByCSV(ctx *gin.Context, expense *model.CreateExpenseByCSV) error {
	var events []*model.Event
	var attendees []*model.Attendees
	file, err := expense.File.Open()
	if err != nil {
		log.Fatalln("Error in opening file")
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatalln("Error in reading file")
	}
	for i := range records {
		if i <= 2 {
			continue
		}
		if records[i][0] == "" {
			continue
		}
		// TODO goroutine 사용하도록 refactoring 해도 좋을듯
		var event model.Event
		var attendee model.Attendees
		time, err := util.StringToTime(records[i][2])
		if err != nil {
			return err
		}
		amount, err := strconv.Atoi(records[i][1])
		if err != nil {
			return err
		}
		eventID := uuid.New().String()
		event.EventID = eventID
		event.UserID = expense.UserID
		event.EventDate = *time
		event.IsInvited = 2
		event.InviteStatus = "act"
		event.InvitationID = 1
		attendee.EventID = eventID
		attendee.AttendeeID = uuid.New().String()
		attendee.Name = records[i][1]
		attendee.Amount = int64(amount)
		switch records[i][4] {
		case "Y":
			attendee.IsAttended = 2
		case "N":
			attendee.IsAttended = 1
		default:
			return errors.New("invalid is_attended type please write in Y or N")
		}
		attendee.ExpenseType = 1
		events = append(events, &event)
		attendees = append(attendees, &attendee)
	}
	for i := range events {
		if err = s.repo.GetTransaction(ctx).Transaction(func(tx *gorm.DB) error {
			eventResult := tx.Create(events[i])
			if eventResult.Error != nil {
				return eventResult.Error
			}
			if eventResult.RowsAffected == 0 {
				return errors.New("event create failed")
			}
			attendeeResult := tx.Create(attendees[i])
			if attendeeResult.Error != nil {
				return attendeeResult.Error
			}
			if attendeeResult.RowsAffected == 0 {
				return errors.New("event create failed")
			}
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}
