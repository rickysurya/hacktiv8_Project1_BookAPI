package main

import (
	"hacktiv8_Project1_BookAPI/database"
	"hacktiv8_Project1_BookAPI/routers"
)

func main() {
	var PORT = ":4000"

	database.StartDB()
	err := routers.StartServer().Run(PORT)
	if err != nil {
		return
	}
}
