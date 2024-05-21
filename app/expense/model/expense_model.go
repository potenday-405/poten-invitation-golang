package model

import (
	"poten-invitation-golang/util"
	"time"
)

type isInvited string

func (i isInvited) GetIntValue() int8 {
	var res int8
	switch i {
	case "invited":
		res = 1
	case "inviting":
		res = 2
	}

	return res
}

type CreateExpense struct {
	UserID     string `json:"user_id"`
	IsInvited  string `json:"is_invited" binding:"required"`
	Name       string `json:"name" binding:"required"`
	EventDate  string `json:"event_date" binding:"required"`
	Expense    int64  `json:"expense" binding:"required"`
	Relation   string `json:"relation"`
	IsAttended int8   `json:"is_attended"`
	Link       string `json:"link"`
}

func (t *CreateExpense) ToEntity() (*Event, *Attendees, error) {
	event := Event{
		UserID:       t.UserID,
		IsInvited:    isInvited(t.IsInvited).GetIntValue(),
		EventDate:    util.StringToTime(t.EventDate).UTC(),
		InvitationID: 1,
	}
	attendees := Attendees{
		Name:        t.Name,
		Relation:    t.Relation,
		Amount:      t.Expense,
		ExpenseType: 1,
		IsAttended:  t.IsAttended,
	}

	return &event, &attendees, nil
}

type UpdateExpense struct {
	EventID    string `json:"event_id" binding:"required"`
	UserID     string `json:"user_id"`
	IsInvited  string `json:"is_invited" binding:"required"`
	Name       string `json:"name" binding:"required"`
	EventDate  string `json:"event_date" binding:"required"`
	Expense    int64  `json:"expense" binding:"required"`
	Relation   string `json:"relation"`
	IsAttended int8   `json:"is_attended"`
	Link       string `json:"link"`
}

func (t *UpdateExpense) ToEntity() (*Event, *Attendees, error) {
	event := Event{
		UserID:       t.UserID,
		IsInvited:    isInvited(t.IsInvited).GetIntValue(),
		EventDate:    util.StringToTime(t.EventDate).UTC(),
		InvitationID: 1,
	}
	attendees := Attendees{
		Name:        t.Name,
		Relation:    t.Relation,
		Amount:      t.Expense,
		ExpenseType: 1,
		IsAttended:  t.IsAttended,
	}

	return &event, &attendees, nil
}

type DeleteExpense struct {
	EventID string `json:"event_id" binding:"required"`
}

type GetExpense struct {
	UserID  string `json:"user_id"`
	EventID string `json:"event_id" binding:"required"`
}

type GetExpenseList struct {
	UserID          string `json:"user_id"`
	IsInvited       string `json:"is_invited" binding:"required"`
	Offset          string `json:"offset"`
	OffsetOrderType int8   `json:"offset_order_type"`
	Order           string `json:"order"`
	Limit           int    `json:"limit"`
	Page            int    `json:"page"`
}

type GetExpenseTotal struct {
	UserID          string `json:"user_id"`
	IsInvited       string `json:"is_invited" binding:"required"`
	Offset          string `json:"offset"`
	OffsetOrderType int8   `json:"offset_order_type"`
}

type GetExpenseSearch struct {
	UserID    string `json:"user_id"`
	IsInvited string `json:"is_invited" binding:"required"`
	Name      string `json:"name" binding:"required"`
}

//=====ENTITY=====

type Event struct {
	EventID      string
	UserID       string
	InvitationID int8
	IsInvited    int8
	InviteStatus string
	EventDate    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (t Event) TableName() string {
	return "event"
}

type Attendees struct {
	AttendeeID  string
	EventID     string
	Name        string
	Relation    string
	Amount      int64
	ExpenseType int8
	IsAttended  int8
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t Attendees) TableName() string {
	return "attendees"
}

type ResponseExpense struct {
	EventId    string    `json:"event_id"`
	UserID     string    `json:"user_id"`
	IsInvited  int8      `json:"is_invited"`
	EventDate  time.Time `json:"event_date"`
	Name       string    `json:"name"`
	Relation   string    `json:"relation"`
	Amount     int64     `json:"amount"`
	IsAttended int8      `json:"is_attended"`
}

type ResponseExpenseTotal struct {
	IsInvited    int8  `json:"is_invited"`
	ExpenseCount int   `json:"expense_count"`
	ExpenseTotal int64 `json:"total_expense"`
}
