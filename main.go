package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	mongodb "main.go/mongoDB"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Some error occured. Err: %s \n", err)
	}

	m, err := mongodb.InitDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(m.Client)
}
