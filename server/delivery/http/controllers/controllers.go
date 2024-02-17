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

type updtUsr struct {
	Approver string `json:"approver"`
	Username string `json:"username"`
	His_no   int    `json:"his_no"`
	Trx_tot  int    `json:"trx_tot"`
}

type getData struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Class    string `json:"Class"`
}

type data struct {
	Username  string `json:"username"`
	Trx_tot   string `json:"trx_tot"`
	Remark    string `json:"remark"`
	Apprv_usr string `json:"apprv_usr"`
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
		return c.Status(401).JSON(fiber.Map{
			"status":  401,
			"message": "Request invalid",
		})
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
			queryInsrtInf := fmt.Sprintf("INSERT INTO score_inf(username, total, reg_dt, updt_dt, his_tot) VALUES (%s, 0, current_date::timestamp, current_date::timestamp, 1);", body.Username)

			result, er := db.Exec(queryInsrtInf)
			if er != nil {
				log.Printf("(CONTOLLERS:2002): %s", er)
				return c.Status(401).JSON(fiber.Map{
					"status":  401,
					"message": "Create data failed",
				})
			}

			row1, _ := result.RowsAffected()
			if row1 != 1 {
				log.Printf("(CONTOLLERS:2003): Expected 1 row to be inserted.")
				return c.Status(401).JSON(fiber.Map{
					"status":  401,
					"message": "Create data failed",
				})
			}

			queryGet2 := fmt.Sprintf("SELECT username FROM score_his WHERE username = '%s'", body.Username)
			err2 := db.QueryRow(queryGet2).Scan(usrInf2)
			if err2 != nil {
				if err2 == sql.ErrNoRows {
					queryInsertHis := fmt.Sprintf("INSERT INTO score_his(username, trx_tot, total, reg_dt, his_no, remark, apprv_usr) VALUES ('%s', 0, 0, current_time::timestamp, 1, '%s', NULL);", body.Username,
						body.Remark)

					result, err := db.Exec(queryInsertHis)
					if err != nil {
						log.Printf("(CONTROLLERS:2004): %s", err)
						return c.Status(401).JSON(fiber.Map{
							"status":  401,
							"message": "Create data failed",
						})
					}
					row1, _ := result.RowsAffected()
					if row1 != 1 {
						log.Printf("(CONTROLLERS:2005): Expected 1 row to be inserted")
						return c.Status(401).JSON(fiber.Map{
							"status":  401,
							"message": "Create data failed",
						})
					}
				}
			}
		} else {
			log.Printf("(CONTROLLERS:2010): %s", err)
			return c.Status(401).JSON(fiber.Map{
				"status":  401,
				"message": err,
			})
		}
	} else {

		queryInsertHis := fmt.Sprintf("INSERT INTO score_his(username, trx_tot, total, reg_dt, his_no, remark, apprv_usr) VALUES ('%s', 0, 0, current_time::timestamp, his_no + 1, '%s', NULL);", body.Username,
			body.Remark)

		result, err := db.Exec(queryInsertHis)
		if err != nil {
			log.Printf("(CONTROLLERS:20011): %s", err)
			return c.Status(401).JSON(fiber.Map{
				"status":  401,
				"message": "Create data failed",
			})
		}
		row1, _ := result.RowsAffected()
		if row1 != 1 {
			log.Printf("(CONTROLLERS:20012): Expected 1 row to be inserted")
			return c.Status(401).JSON(fiber.Map{
				"status":  401,
				"message": "Create data failed",
			})
		}

	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "Data sucessfully created",
	})
}

func ApproveActivities(db *sql.DB, c *fiber.Ctx) error {
	var body *updtUsr = &updtUsr{}
	var usrInf string
	var usrInf2 string

	if err := c.BodyParser(body); err != nil {
		log.Printf("(CONTROLLERS:3001): %s", err)
		return c.Status(401).JSON(fiber.Map{
			"status":  401,
			"message": "Invalid request",
		})
	}

	queryGet := fmt.Sprintf("SELECT username FROM bsc_usr_inf WHERE username = '%s'", body.Approver)
	err := db.QueryRow(queryGet).Scan(usrInf)

	if err != nil {
		log.Printf("(CONTROLLERS:3002): %s", err)
		return c.Status(401).JSON(fiber.Map{
			"status":  401,
			"message": "User not found",
		})
	}

	queryGet2 := fmt.Sprintf("SELECT approver FROM score_his WHERE username = '%s' AND his_no = %d", body.Username, body.His_no)

	errs := db.QueryRow(queryGet2).Scan(usrInf2)

	if errs != nil && errs != sql.ErrNoRows {
		log.Printf("(CONTROLLERS:3002): %s", errs)
		return c.Status(401).JSON(fiber.Map{
			"status":  401,
			"message": "Get data failed",
		})
	}

	if usrInf2 != "" {
		log.Printf("(CONTOLLERS:3003): Already approved")
		return c.Status(401).JSON(fiber.Map{
			"status":  401,
			"message": "Already approved",
		})

	}

	queryUpdate := fmt.Sprintf("UPDATE score_inf SET total = total + %d WHERE username = '%s' AND his_no = %d", body.Trx_tot, body.Username, body.His_no)

	result, err := db.Exec(queryUpdate)

	if err != nil {
		log.Printf("(CONTOLLERS:3004): %s", err)
		return c.Status(401).JSON(fiber.Map{
			"status":  401,
			"message": "Create data failed",
		})
	}

	res, _ := result.RowsAffected()

	if res != 1 {
		log.Printf("(CONTROLLERS:3005) : Expected update 1 row")
		return c.Status(401).JSON(fiber.Map{
			"status":  401,
			"message": "Create data failed",
		})
	}

	queryUpdate2 := fmt.Sprintf("UPDATE score_his SET trx_tot = %d, total = total + %d, app WHERE username = '%s' AND his_no = %d", body.Trx_tot, body.Trx_tot, body.Username, body.His_no)

	rest, er := db.Exec(queryUpdate2)

	if er != nil {
		log.Printf("(CONTROLLERS:3006): %s", er)
		return c.Status(401).JSON(fiber.Map{
			"status":  401,
			"message": "Create data failed",
		})
	}

	res2, _ := rest.RowsAffected()

	if res2 != 1 {
		log.Printf("(CONTROLLERS:3007) : Expected update 1 row")
		return c.Status(401).JSON(fiber.Map{
			"status":  401,
			"message": "Create data failed",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "Successfully approved",
	})
}

func GetData(db *sql.DB, c *fiber.Ctx) error {
	var body *getData = &getData{}
	var data1 []data = []data{}

	if err := c.BodyParser(body); err != nil {
		log.Printf("(CONTROLLERS:4001): %s", err)
	}

	if body.Role == "0" || body.Role == "99" {
		querySelect := fmt.Sprintf("SELECT username, trx_tot, remark, apprv_usr FROM FROM score_his a INNER JOIN bsc_usr_inf b ON a.username = b.username WHERE b.class = '%s' AND b.role not in ('0', '99')", body.Class)

		rows, err := db.Query(querySelect)
		if err != nil {
			log.Printf("(CONTROlLERS:4002): %s", err)
		}
		for rows.Next() {
			var data2 *data = &data{}
			if err := rows.Scan(&data2.Username, &data2.Trx_tot, &data2.Remark,
				&data2.Apprv_usr); err != nil {
				return err
			}
			data1 = append(data1, *data2)
		}
	} else {
		querySelect := fmt.Sprintf("SELECT * FROM score_his WHERE username = %s", body.Username)
		rows, err := db.Query(querySelect)
		if err != nil {
			log.Printf("(CONTROlLERS:4003): %s", err)
		}
		for rows.Next() {
			var data2 *data = &data{}
			if err := rows.Scan(&data2.Username, &data2.Trx_tot, &data2.Remark,
				&data2.Apprv_usr); err != nil {
				return err
			}
			data1 = append(data1, *data2)
		}
	}

	return c.Status(200).JSON(fiber.Map{
		"status":   200,
		"response": &data1,
	})
}
