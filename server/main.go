package main

import (
	"log"
	"os"
	"point/delivery/http"

	_ "github.com/lib/pq"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("(MAIN:1000): ", err)
	}
	currentDir := path
	app := http.NewHttpDelivery(currentDir)

	app.Listen(":3000")

}
