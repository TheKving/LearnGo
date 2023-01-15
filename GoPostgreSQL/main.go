package main

import (
	"fmt"
	"log"
	"time"

	"github.com/TheKving/LearnGo/tree/main/GoPostgreSQL/connection"
	"github.com/TheKving/LearnGo/tree/main/GoPostgreSQL/models"
	"gorm.io/gorm"
)

func errorFatal(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func main() {
	conn, err := connection.GetConnection("localhost", "postgres", "postgres", "postgres", "5432")
	errorFatal(err)
	_ = conn
	//CreateTask(conn)
	fmt.Println("-----------")
	ReadTask(conn)
	fmt.Println("-----------")
	ReadTaskByID(conn, 11)
	fmt.Println("-----Update table------")
	UpdateTask(conn, 11, "Esta tarea se modifico", false, time.Now())
	ReadTaskByID(conn, 11)
	fmt.Println("-----Delete task------")
	DeleteTask(conn, 1)
	ReadTask(conn)
	fmt.Println("Connected to BD")
}

func CreateTask(conn *gorm.DB) {
	// Llamamos al struct TodoList, le pasamos los parámetros que hay que rellenar.
	task := models.TodoList{Task: "Tarea 456", Status: true, DueDate: time.Now()}
	tasknew := models.TodoList{Task: "Tarea 123", Status: true, DueDate: time.Date(2023, 01, 4, 15, 0, 0, 0, time.UTC)}

	//llamamos a la dirección de memoria de task y creamos el insert en la DB
	//res := conn.Create(&task)
	res := conn.Create([]models.TodoList{task, tasknew})
	if res.Error != nil {
		fmt.Println(res.Error)
	}
}

func ReadTask(conn *gorm.DB) {
	var task []models.TodoList
	if err := conn.Find(&task).Error; err != nil {
		fmt.Println(err.Error())
	}

	for _, t := range task {
		fmt.Println(t)
	}
}

func ReadTaskByID(conn *gorm.DB, id int) {
	var task models.TodoList
	r := conn.Where("ID=?", id).Limit(1).Find(&task)
	if r.Error != nil {
		fmt.Println(r.Error)
		return
	}
	if r.RowsAffected == 0 {
		fmt.Printf("Not exist task with the ID %d\n", id)
		return
	}
	fmt.Println(task)
}

func UpdateTask(conn *gorm.DB, id int, taskTxt string, status bool, dueDate time.Time) {
	var task = models.TodoList{Task: taskTxt, Status: status, DueDate: dueDate}
	r := conn.Select("Task").Where("ID = ?", id).Updates(&task)
	if r.Error != nil {
		fmt.Println(r.Error)
		return
	}
}

func DeleteTask(conn *gorm.DB, id int) {
	r := conn.Where("ID = ?", id).Delete(&models.TodoList{})
	if r.Error != nil {
		fmt.Println(r.Error)
		return
	}
	if r.RowsAffected == 0 {
		fmt.Printf("No deleted task with the ID %d\n", id)
	}
}
