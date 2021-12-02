package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/joho/godotenv"
	"api-gofiber/test/config"
	"api-gofiber/test/helpers"
	"api-gofiber/test/models"
)

func main() {
	errEnv := godotenv.Load()

	if errEnv != nil {
		fmt.Println(errEnv)
		os.Exit(1)
	}

	var dsn map[string]string = map[string]string{
		"DB_HOST" : os.Getenv("DB_HOST"),
		"DB_PORT" : os.Getenv("DB_PORT"),
		"DB_NAME" : os.Getenv("DB_NAME"),
		"DB_USER" : os.Getenv("DB_USER"),
		"DB_PASSWORD" : os.Getenv("DB_PASSWORD"),
	}

	database, errDb := config.Database(dsn)

	if errDb == nil {
		for i := 0; i < 3; i++ {
			hash, _ := helpers.HashPassword("password")
			var id string = "user" + strconv.Itoa(i)

			user := models.User{
				Name:     id,
				Email:    id + "@gmail.com",
				Password: hash,
			}

			result := database.Create(&user)

			if result.Error == nil {
				fmt.Println("Success Memasukan " + strconv.Itoa(int(result.RowsAffected)) + " Data")
			} else {
				fmt.Print(result.Error)
				os.Exit(1)
			}
		}

		for i := 0; i < 30; i++ {
			var id string = "data" + strconv.Itoa(i)

			data := models.Data{
				Name:  id,
				Phone: nil,
			}

			result := database.Create(&data)

			if result.Error == nil {
				fmt.Println("Success Memasukan " + strconv.Itoa(int(result.RowsAffected)) + " Data")
			} else {
				fmt.Print(result.Error)
				os.Exit(1)
			}
		}
		fmt.Println("Success Seed")
	}
}
