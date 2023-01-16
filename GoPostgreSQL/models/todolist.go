package models

import "time"

type TodoList struct {
	ID        int `gorm:"autoincrement"`
	Task      string
	Status    bool
	CreatedAt time.Time `gorm:"autoincrement"`
	DueDate   time.Time
}
