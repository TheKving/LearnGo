package main

import (
	"fmt"
	"log"

	"connection"
)

func errorFatal(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func main() {
	conn, err := connection.GetConnection("postgres", "postgres", "postgres", "postgres", "5432")
	errorFatal(err)
	_ = conn
	fmt.Println("Connected to BD")
}
