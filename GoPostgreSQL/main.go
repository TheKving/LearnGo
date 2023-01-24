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
	//_ = conn

	var option, id int
	MenuOptions()
	for {
		fmt.Scan(&option)
		if option == 0 {
			fmt.Println("Good bye!")
			break
		}
		switch option {
		case 1:
			CreateTask(conn, "New Task by program", true, time.Now())
			//CreateTask(conn *gorm.DB, task string, status bool, duedate time.Time)
			MenuOptions()
		case 2:
			ReadTask(conn)
			MenuOptions()
		case 3:
			fmt.Printf("Enter the ID to search task: ")
			fmt.Scan(&id)
			searchTask := ReadTaskByID(conn, id)
			if searchTask.ID != 0 {
				fmt.Println(searchTask)
			}
			MenuOptions()
		case 4:
			var statusUser int
			var textTask string
			var status bool

			fmt.Printf("Enter the ID to update task: ")
			fmt.Scanln(&id)

			fmt.Printf("\nEnter task to update to do: ")
			fmt.Scan(&textTask)

			fmt.Printf("\nEnter the new status (0 = Done, 1 = To do): ")
			fmt.Scanln(&statusUser)
			if statusUser != 0 {
				status = false
			} else {
				status = true
			}

			searchTask := ReadTaskByID(conn, id)
			if searchTask.ID != 0 {
				UpdateTask(conn, id, textTask, status, time.Now())
				//fmt.Println(ReadTaskByID(conn, id))
			}
			MenuOptions()
		case 5:
			fmt.Printf("ID to delete task: ")
			fmt.Scan(&id)
			DeleteTask(conn, id)
			MenuOptions()
		default:
			MenuOptions()
		}
	}
}

func CreateTask(conn *gorm.DB, task string, status bool, dueDate time.Time) {
	// Llamamos al struct TodoList, le pasamos los parámetros que hay que rellenar.
	newTask := models.TodoList{Task: task, Status: status, DueDate: dueDate}

	res := conn.Create(&newTask)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
}

/*func CreateTask(conn *gorm.DB) {
	// Llamamos al struct TodoList, le pasamos los parámetros que hay que rellenar.
	task := models.TodoList{Task: "Tarea 456", Status: true, DueDate: time.Now()}
	tasknew := models.TodoList{Task: "Tarea 123", Status: true, DueDate: time.Date(2023, 01, 4, 15, 0, 0, 0, time.UTC)}

	//llamamos a la dirección de memoria de task y creamos el insert en la DB
	//res := conn.Create(&task)
	res := conn.Create([]models.TodoList{task, tasknew})
	if res.Error != nil {
		fmt.Println(res.Error)
	}
}*/

func ReadTask(conn *gorm.DB) {
	var task []models.TodoList
	if err := conn.Find(&task).Error; err != nil {
		fmt.Println(err.Error())
	}

	for _, t := range task {
		fmt.Println(t)
	}
}

func ReadTaskByID(conn *gorm.DB, id int) (task models.TodoList) {
	r := conn.Where("ID=?", id).Limit(1).Find(&task)
	if r.Error != nil {
		fmt.Println(r.Error)
		return
	}
	if r.RowsAffected == 0 {
		fmt.Printf("Not exist task with the ID %d\n", id)
		task = models.TodoList{}
	}
	return task
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

func MenuOptions() {
	fmt.Printf("\n*** Menu To Do List ***\n1 - Create task\n2 - Search all task\n3 - Search task\n4 - Update task\n5 - Delete task\n0 - Exit\n***********************\n\nSelect Option: ")
}
