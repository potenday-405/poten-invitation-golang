package model

import "time"

type isInvited string

func (i isInvited) GetIntValue() int8 {
	var res int8
	switch i {
	case "invited":
		res = 1
	case "inviting":
		res = 0
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

func (t CreateExpense) ToEntity() (*Event, *Attendees, error) {
	eventDate, err := time.Parse(time.DateTime, t.EventDate)
	if err != nil {
		return nil, nil, err
	}
	event := Event{
		UserID:    t.UserID,
		IsInvited: isInvited(t.IsInvited).GetIntValue(),
		EventDate: eventDate,
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

func (t CreateExpense) TableName() string {
	return "expense"
}

type Event struct {
	EventID       string
	UserID        string
	InvitationID  int8
	IsInvited     int8
	InvitedStatus string
	EventDate     time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
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
}
