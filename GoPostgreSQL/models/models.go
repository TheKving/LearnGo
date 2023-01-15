package main

import "time"

type TodoList struct {
	ID        int       `db:"id"`
	Task      string    `db:"task"`
	Status    bool      `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	DueDate   time.Time `db:"due_date"`
}
