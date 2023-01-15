package models

import "time"

type TodoList struct {
	ID        int `gorm:"autoincrement"`
	Task      string
	Status    bool
	CreatedAt time.Time `gorm:"autoincrement"`
	DueDate   time.Time
}

/*

	ID        int `gorm:"-"`
	Task      string
	Status    bool
	CreatedAt time.Time `gorm:"-"`
	DueDate   time.Time
	ID        int       `db:"id"`
	Task      string    `db:"task"`
	Status    bool      `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	DueDate   time.Time `db:"due_date"`


En mi base de datos postgres el id SERIAL PRIMARY KEY como lo leo con GORM? debo declararlo en el type struct? y como añado datos si lo declaro?

gorm:"primary_key;auto_increment" en el campo ID de tu estructura de modelo. Esto indicará a GORM que el c
*/
