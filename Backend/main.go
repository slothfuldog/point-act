package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("dbEnv"), os.Getenv("dbPort"),
		os.Getenv("dbUser"), os.Getenv("dbPassword"),
		os.Getenv("dbName"))

	db, er := sql.Open("postgres", sqlInfo)
	if er != nil {
		log.Panic(er)
	}

	defer db.Close()

	er = db.Ping()
	if er != nil {
		log.Panic(er)
	}

	fmt.Println("Database Has Sucessfully Connected")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")

}
