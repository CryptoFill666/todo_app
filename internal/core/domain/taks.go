package domain

import "time"

type Task struct {
	Id          int
	Version     int
	UserID      int
	Title       string
	Description *string
	Status      bool
	CreatedAt   time.Time
	DoneAt      *time.Time
}
