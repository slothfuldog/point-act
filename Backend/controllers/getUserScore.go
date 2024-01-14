package controllers

import (
	"database/sql"
	"fmt"
)

type UserScore struct {
	username string `json:"username"`
	total    int    `json:"total"`
}

func GetUserScore(db *sql.DB) {
	query := `SELECT * FROM `
	fmt.Printf("%s", query)
}
