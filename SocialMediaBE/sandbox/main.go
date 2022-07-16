package main

import (
	"fmt"
	"log"

	"github.com/Conor-Fleming/SocialMediaBE/internal/database"
)

func main() {
	c := database.NewClient("db.json")
	err := c.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database ensured!")
}
