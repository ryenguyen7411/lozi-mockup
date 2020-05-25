package main

import (
	"fmt"
	"log"
	"main/helpers"
)

func main() {
	helpers.OpenDB()
	defer helpers.CloseDB()

	// _, err := helpers.Create([]byte("today"), []byte("Monday"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	today, err := helpers.Delete([]byte("today"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted: Today is", today)

	// Initialize router
	// configs.HandleRouter()
}
