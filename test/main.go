package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Read enviroment var of .env
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	//Exec command hostname -f to obtain the hostname of server
	hostname, _ := exec.Command("hostname", "-f").Output()
	//Obtain the public ip address of server
	PUBLICIPADDRESS1, err := exec.Command("curl", "ifconfig.co").Output()
	/*PUBLICIPADDRESS2, _ := exec.Command("curl", "ifconfig.me").Output()
	PUBLICIPADDRESS3, _ := exec.Command("curl icanhazip.com").Output()*/

	if err != nil {
		fmt.Println(string(PUBLICIPADDRESS1))
		log.Fatal("Not give ip")
	}

	//Print env vars, hostname and IP (fails in work, in sv works?)
	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(string(hostname))
	fmt.Println(string(PUBLICIPADDRESS1))
	/*fmt.Println(string(PUBLICIPADDRESS2))
	fmt.Println(string(PUBLICIPADDRESS3))*/

	//to do, get public ip address from other service
	//update the dynamic ip address to google domain service with curl or other?
	//test in production?
}

//curl ifconfig.co
// curl ifconfig.me
// curl icanhazip.com
