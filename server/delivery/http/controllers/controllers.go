package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"point/domain/entity"
	com "point/infrastructure/functions"

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
	Class    string `json:"class"`
	Filter   string `json:"filter"`
}

type data struct {
	Username  string `json:"username"`
	Trx_tot   string `json:"trx_tot"`
	Remark    string `json:"remark"`
	Apprv_usr string `json:"apprv_usr"`
	His_no    int    `json:"his_no"`
}

type users entity.Usr
type scrHis entity.ScoreHis

func Login(db *sql.DB, c *fiber.Ctx) error {
	var user users
	var body *usr = &usr{}

	if err := c.BodyParser(body); err != nil {
		com.PrintLog(fmt.Sprintf("(CONTROLLERS:0001): %s", err))
	}

	var available int

	query := fmt.Sprintf("SELECT count(1)::int FROM bsc_usr_inf WHERE username = '%s';", body.Username)

	err := db.QueryRow(query).Scan(&available)

	if err != nil {
		if err == sql.ErrNoRows || available == 0 {
			com.PrintLog(fmt.Sprintf("(CONTROLLERS:0002): %s", err))
			log.Printf("Check")
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
			com.PrintLog(fmt.Sprintf("(CONTROLLERS:0003): %s", err))
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
		com.PrintLog(fmt.Sprintf("(CONTROLLERS:1001): %s", err))
		return c.Status(401).JSON(fiber.Map{
			"status":  401,
			"message": "Request invalid",
		})
	}

	if !body.IsLogin {
		com.PrintLog(fmt.Sprintf("(CONTROLLERS: 1002): user has not login"))
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
			com.PrintLog(fmt.Sprintf("(CONTROLLERS:1003): %s", err))
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
		com.PrintLog(fmt.Sprintf("(CONTROLLERS:2001): %s", err))
	}

	queryGet := fmt.Sprintf("SELECT username FROM score_inf WHERE username = '%s'", body.Username)
	err := db.QueryRow(queryGet).Scan(&usrInf)
	if err != nil {
		if err == sql.ErrNoRows {
			queryInsrtInf := fmt.Sprintf("INSERT INTO score_inf(username, total, reg_dt, updt_dt, his_tot) VALUES ('%s', 0, now()::timestamp at time zone 'Asia/Bangkok', now()::timestamp at time zone 'Asia/Bangkok', 1);", body.Username)

			result, er := db.Exec(queryInsrtInf)
			if er != nil {
				com.PrintLog(fmt.Sprintf("(CONTROLLERS:2002): %s", er))
				return c.Status(401).JSON(fiber.Map{
					"status":  401,
					"message": "Create data failed",
				})
			}

			row1, _ := result.RowsAffected()
			if row1 != 1 {
				com.PrintLog(fmt.Sprintf("(CONTROLLERS:2003): Expected 1 row to be inserted."))
				return c.Status(401).JSON(fiber.Map{
					"status":  401,
					"message": "Create data failed",
				})
			}

			queryGet2 := fmt.Sprintf("SELECT username FROM score_his WHERE username = '%s'", body.Username)
			err2 := db.QueryRow(queryGet2).Scan(&usrInf2)
			if err2 != nil {
				if err2 == sql.ErrNoRows {
					queryInsertHis := fmt.Sprintf("INSERT INTO score_his(username, trx_tot, total, reg_dt, his_no, remark, apprv_usr) VALUES ('%s', 0, 0, now()::timestamp at time zone 'Asia/Bangkok', 1, '%s', NULL);", body.Username,
						body.Remark)

					result, err := db.Exec(queryInsertHis)
					if err != nil {
						com.PrintLog(fmt.Sprintf("(CONTROLLERS:2004): %s", err))
						return c.Status(401).JSON(fiber.Map{
							"status":  401,
							"message": "Create data failed",
						})
					}
					row1, _ := result.RowsAffected()
					if row1 != 1 {
						com.PrintLog(fmt.Sprintf("(CONTROLLERS:2005): Expected 1 row to be inserted"))
						return c.Status(401).JSON(fiber.Map{
							"status":  401,
							"message": "Create data failed",
						})
					}
				}
			}
		} else {
			com.PrintLog(fmt.Sprintf("(CONTROLLERS:2010): %s", err))
			return c.Status(401).JSON(fiber.Map{
				"status":  401,
				"message": err,
			})
		}
	} else {

		queryInsertHis := fmt.Sprintf("INSERT INTO score_his(username, trx_tot, total, reg_dt, his_no, remark, apprv_usr) VALUES ('%s', 0, 0, now()::timestamp at time zone 'Asia/Bangkok', (select MAX(his_no) + 1 from score_his where username = '%s'), '%s', NULL);", body.Username,
			body.Username, body.Remark)

		result, err := db.Exec(queryInsertHis)
		if err != nil {
			com.PrintLog(fmt.Sprintf("(CONTROLLERS:20011): %s", err))
			return c.Status(401).JSON(fiber.Map{
				"status":  401,
				"message": "Create data failed",
			})
		}
		row1, _ := result.RowsAffected()
		if row1 != 1 {
			com.PrintLog(fmt.Sprintf("(CONTROLLERS:2012): Expected 1 row to be inserted"))
			return c.Status(401).JSON(fiber.Map{
				"status":  401,
				"message": "Create data failed",
			})
		}

		queryUpdateInf := fmt.Sprintf("UPDATE score_inf SET his_tot = (SELECT MAX(his_no) FROM score_his WHERE username = '%s'), updt_dt = now()::timestamp at time zone 'Asia/Bangkok' WHERE username = '%s'", body.Username, body.Username)

		resultUp, errUp := db.Exec(queryUpdateInf)

		if errUp != nil {
			com.PrintLog(fmt.Sprintf("(CONTROLLERS:20011): %s", errUp))
			return c.Status(401).JSON(fiber.Map{
				"status":  401,
				"message": "Create data failed",
			})
		}
		row2, _ := resultUp.RowsAffected()
		if row2 != 1 {
			com.PrintLog(fmt.Sprintf("(CONTROLLERS:20012): Expected 1 row to be inserted"))
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
	var scoreHis int

	if err := c.BodyParser(body); err != nil {
		com.PrintLog(fmt.Sprintf("(CONTROLLERS:3001): %s", err))
		return c.Status(401).JSON(fiber.Map{
			"status":  401,
			"message": "Invalid request",
		})
	}

	com.PrintLog(fmt.Sprintf("Approver   =    [%s]", body.Approver))
	com.PrintLog(fmt.Sprintf("Username   =    [%s]", body.Username))
	com.PrintLog(fmt.Sprintf("His_no     =    [%d]", body.His_no))
	com.PrintLog(fmt.Sprintf("Trx_tot    =    [%d]", body.Trx_tot))

	queryGet1 := fmt.Sprintf("SELECT trx_tot FROM score_his WHERE username = '%s' AND his_no = %d", body.Username, body.His_no)

	errss := db.QueryRow(queryGet1).Scan(&scoreHis)

	if errss != nil && errss != sql.ErrNoRows {
		com.PrintLog(fmt.Sprintf("(CONTROLLERS:3011): %s", errss))
		return c.Status(401).JSON(fiber.Map{
			"status":  401,
			"message": "Get data failed",
		})
	}

	queryGet := fmt.Sprintf("SELECT username FROM bsc_usr_inf WHERE username = '%s'", body.Approver)
	err := db.QueryRow(queryGet).Scan(&usrInf)

	if err != nil {
		com.PrintLog(fmt.Sprintf("(CONTROLLERS:3002): %s", err))
		return c.Status(401).JSON(fiber.Map{
			"status":  401,
			"message": "User not found",
		})
	}

	queryGet2 := fmt.Sprintf("SELECT COALESCE(apprv_usr,'') FROM score_his WHERE username = '%s' AND his_no = %d", body.Username, body.His_no)

	errs := db.QueryRow(queryGet2).Scan(&usrInf2)

	if errs != nil && errs != sql.ErrNoRows {
		com.PrintLog(fmt.Sprintf("(CONTROLLERS:3011): %s", errs))
		return c.Status(401).JSON(fiber.Map{
			"status":  401,
			"message": "Get data failed",
		})
	}

	// if usrInf2 != "" {
	// 	com.PrintLog(fmt.Sprintf("(CONTROLLERS:3003): Already approved"))
	// 	return c.Status(401).JSON(fiber.Map{
	// 		"status":  401,
	// 		"message": "Already approved",
	// 	})

	// }

	if scoreHis < 1 {
		queryUpdate := fmt.Sprintf("UPDATE score_inf SET total = (SELECT total FROM score_inf WHERE username = '%s') + %d WHERE username = '%s'", body.Username, body.Trx_tot, body.Username)

		result, err := db.Exec(queryUpdate)

		com.PrintLog(fmt.Sprintf("%s", queryUpdate))

		if err != nil {
			com.PrintLog(fmt.Sprintf("(CONTROLLERS:3004): %s", err))
			return c.Status(401).JSON(fiber.Map{
				"status":  401,
				"message": "Create data failed",
			})
		}

		res, _ := result.RowsAffected()

		if res != 1 {
			com.PrintLog(fmt.Sprintf("(CONTROLLERS:3005) : Expected update 1 row"))
			com.PrintLog(fmt.Sprintf("Effected Row   = [%d]", res))
			return c.Status(401).JSON(fiber.Map{
				"status":  401,
				"message": "Create data failed",
			})
		}
	} else {
		queryUpdate := fmt.Sprintf("UPDATE score_inf SET total = (SELECT total FROM score_inf WHERE username = '%s') - %d + %d WHERE username = '%s'", body.Username, scoreHis, body.Trx_tot, body.Username)

		result, err := db.Exec(queryUpdate)

		com.PrintLog(fmt.Sprintf("%s", queryUpdate))

		if err != nil {
			com.PrintLog(fmt.Sprintf("(CONTROLLERS:3004): %s", err))
			return c.Status(401).JSON(fiber.Map{
				"status":  401,
				"message": "Create data failed",
			})
		}

		res, _ := result.RowsAffected()

		if res != 1 {
			com.PrintLog(fmt.Sprintf("(CONTROLLERS:3005) : Expected update 1 row"))
			com.PrintLog(fmt.Sprintf("Effected Row   = [%d]", res))
			return c.Status(401).JSON(fiber.Map{
				"status":  401,
				"message": "Create data failed",
			})
		}
	}

	queryUpdate2 := fmt.Sprintf("UPDATE score_his SET trx_tot = %d, total = total + %d, apprv_usr = '%s' WHERE username = '%s' AND his_no = %d", body.Trx_tot, body.Trx_tot, body.Approver, body.Username, body.His_no)

	rest, er := db.Exec(queryUpdate2)

	com.PrintLog(fmt.Sprintf("%s", queryUpdate2))

	if er != nil {
		com.PrintLog(fmt.Sprintf("(CONTROLLERS:3006): %s", er))
		return c.Status(401).JSON(fiber.Map{
			"status":  401,
			"message": "Create data failed",
		})
	}

	res2, _ := rest.RowsAffected()

	if res2 != 1 {
		com.PrintLog(fmt.Sprintf("(CONTROLLERS:3007) : Expected update 1 row"))
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

	com.PrintLog(fmt.Sprintf("=========GetData START======"))

	if err := c.BodyParser(body); err != nil {
		com.PrintLog(fmt.Sprintf("(CONTROLLERS:4001): %s", err))
	}

	if body.Role == "1" || body.Role == "99" {
		querySelect := fmt.Sprintf("SELECT b.username, his_no, trx_tot, remark, COALESCE(apprv_usr, 'NOT APPROVED') FROM score_his AS a INNER JOIN bsc_usr_inf AS b ON a.username = b.username WHERE b.class = '%s' AND b.role not in ('0', '99')", body.Class)

		if body.Filter == "scored" {
			querySelect = fmt.Sprintf("%s AND total <> 0", querySelect)
		} else if body.Filter == "unscored" {
			querySelect = fmt.Sprintf("%s AND (total = 0 OR total IS NULL)", querySelect)
		}

		rows, err := db.Query(querySelect)
		if err != nil {
			com.PrintLog(fmt.Sprintf("(CONTROlLERS:4002): %s", err))
			return c.Status(404).JSON(fiber.Map{
				"status":  404,
				"message": "Data Not Found",
			})
		}
		for rows.Next() {
			var data2 *data = &data{}
			if err := rows.Scan(&data2.Username, &data2.His_no, &data2.Trx_tot, &data2.Remark,
				&data2.Apprv_usr); err != nil {
				return err
			}
			data1 = append(data1, *data2)
		}
	} else {
		querySelect := fmt.Sprintf("SELECT username, his_no, trx_tot, remark, COALESCE(apprv_usr, 'NOT APPROVED') FROM score_his WHERE username = '%s'", body.Username)

		if body.Filter == "scored" {
			querySelect = fmt.Sprintf("%s AND total <> 0", querySelect)
		} else if body.Filter == "unscored" {
			querySelect = fmt.Sprintf("%s AND (total = 0 OR total IS NULL)", querySelect)
		}

		com.PrintLog(fmt.Sprintf("%s", querySelect))

		rows, err := db.Query(querySelect)
		if err != nil {
			com.PrintLog(fmt.Sprintf("(CONTROlLERS:4003): %s", err))
			return c.Status(401).JSON(fiber.Map{
				"status":  404,
				"message": "Data Not Found",
			})
		}
		for rows.Next() {
			var data2 *data = &data{}
			if err := rows.Scan(&data2.Username, &data2.His_no, &data2.Trx_tot, &data2.Remark,
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
