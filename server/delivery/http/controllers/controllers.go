package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"point/domain/entity"

	"github.com/gofiber/fiber/v2"
)

type usr struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type keepUsr struct {
	Username string `json:"username"`
	IsLogin  bool   `json:"isLogin"`
}

type users entity.Usr

func Login(db *sql.DB, c *fiber.Ctx) error {
	var user users
	var body *usr = &usr{}

	if err := c.BodyParser(body); err != nil {
		log.Printf("(CONTROLLERS:0001): %s", err)
	}

	var available int

	query := fmt.Sprintf("SELECT count(1)::int FROM bsc_usr_inf WHERE username = '%s';", body.Username)

	err := db.QueryRow(query).Scan(&available)

	if err != nil {
		if err == sql.ErrNoRows || available == 0 {
			log.Printf("(CONTROLLERS:0002): %s", err)
			return c.Status(400).JSON(fiber.Map{
				"status":  400,
				"message": "Username or password not found",
			})
		}
	}

	query2 := fmt.Sprintf("SELECT username, name, class, role, reg_dt, updt_dt, sts FROM bsc_usr_inf WHERE username = '%s' AND password = '%s'", body.Username, body.Password)

	err = db.QueryRow(query2).Scan(&user.Username, &user.Name,
		&user.Class, &user.Role, &user.Reg_dt,
		&user.Updt_dt, &user.Sts)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("(CONTROLLERS:0003): %s", err)
			return c.Status(400).JSON(fiber.Map{
				"status":  400,
				"message": "Username or password not found",
			})
		}
	}

	user.IsLogin = true

	return c.Status(200).JSON(fiber.Map{
		"status":   200,
		"response": &user,
	})
}

func KeepLogin(db *sql.DB, c *fiber.Ctx) error {
	var user users
	var body *keepUsr = &keepUsr{}

	if err := c.BodyParser(body); err != nil {
		log.Printf("(CONTROLLERS:1001): %s", err)
	}

	if !body.IsLogin {
		log.Printf("(CONTOLLERS: 1002): user has not login")
		return c.Status(401).JSON(fiber.Map{
			"status":  401,
			"message": "User has not login",
		})
	}

	query := fmt.Sprintf("SELECT username, name, class, role, reg_dt, updt_dt, sts FROM bsc_usr_inf WHERE username = '%s'", body.Username)

	err := db.QueryRow(query).Scan(&user.Username, &user.Name, &user.Class,
		&user.Role, &user.Reg_dt, &user.Updt_dt, &user.Sts)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("(CONTROLLERS:1003): %s", err)
			user.IsLogin = false
			return c.Status(400).JSON(fiber.Map{
				"status":  400,
				"message": "User not found, logging out..",
			})
		}
	}

	user.IsLogin = true

	return c.Status(200).JSON(fiber.Map{
		"status":   200,
		"response": &user,
	})

}
