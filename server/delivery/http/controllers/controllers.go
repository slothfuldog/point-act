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
type scrHis entity.ScoreHis

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

func InsertActivities(db *sql.DB, c *fiber.Ctx) error {
	var body *scrHis = &scrHis{}
	var usrInf string
	var usrInf2 string

	if err := c.BodyParser(body); err != nil {
		log.Printf("(CONTROLLERS:2001): %s", err)
	}

	queryGet := fmt.Sprintf("SELECT username FROM score_inf WHERE username = '%s'", body.Username)
	err := db.QueryRow(queryGet).Scan(usrInf)
	if err != nil {
		if err == sql.ErrNoRows {
			queryInsrtInf := fmt.Sprintf("INSERT INTO score_his(username, total, reg_dt, updt_dt, his_tot) VALUES (%s, %d, current_date::timestamp, current_date::timestamp, 1);", body.Username, body.Trx_tot)

			result, er := db.Exec(queryInsrtInf)
			if er != nil {
				log.Printf("(CONTOLLERS:2002): %s", er)
			}

			row1, _ := result.RowsAffected()
			if row1 != 1 {
				log.Printf("(CONTOLLERS:2003): Expected 1 row to be inserted.")
			}

			queryGet2 := fmt.Sprintf("SELECT username FROM score_his WHERE username = '%s'", body.Username)
			err2 := db.QueryRow(queryGet2).Scan(usrInf2)
			if err2 != nil {
				if err2 == sql.ErrNoRows {
					queryInsertHis := fmt.Sprintf("INSERT INTO score_his(username, trx_tot, total, reg_dt, his_no, remark, apprv_usr) VALUES ('%s', %d, %d, current_time::timestamp, 1, %s, NULL);", body.Username, body.Trx_tot, body.Trx_tot,
						body.Remark)

					result, err := db.Exec(queryInsertHis)
					if err != nil {
						log.Printf("(CONTROLLERS:2004): %s", err)
					}
					row1, _ := result.RowsAffected()
					if row1 != 1 {
						log.Printf("(CONTROLLERS:2005): Expected 1 row to be inserted")
					}
				}
			}
		} else {
			log.Printf("(CONTROLLERS:2010): %s", err)
		}
	} else {
		queryUpdate := ""
	}

	return nil
}
