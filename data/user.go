package data

import (
	"BetterDash/models"
	"BetterDash/scripts"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB = scripts.InitDB()

func CreateUser(user *models.UserData) bool {
	insert_user := `INSERT INTO users(ID, nick, email, password, verify, active) VALUES (?, ?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insert_user)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	_, err = statement.Exec(user.ID, user.Nick, user.Email, user.Password, user.Verify, user.Active)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func SetID() uint32 {
	row, err := db.Query("SELECT seq FROM sqlite_sequence")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var id uint32
	for row.Next() {
		row.Scan(&id)
		id++
	}
	return id
}

func IsExist(user *models.UserData) bool {
	row, err := db.Query(`select count(id)>0 from users where email="` + user.Email + `"`)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var exist uint8
	for row.Next() {
		row.Scan(&exist)
	}
	switch exist {
	case 0:
		return false
	case 1:
		return true
	default:
		return true // true, because if this function return true, client can't create account with this email
	}
	return true // ^
}

func GetPassword(user *models.UserData) string {
	// GET SOME RESEARCH ABOUT DEBUG ENV or something; maybe debug setup
	// test := `SELECT password FROM users where email="` + user.Email + `"`
	// fmt.Println(test) // dupa@dupa.dupa" "INSERT INTO users(nick) VALUES('hacked'); - crash!!!
	row, err := db.Query(`SELECT password FROM users where email="` + user.Email + `"`)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var password string
	for row.Next() {
		row.Scan(&password)
	}
	return password
}

func GetVerifyStrings() []byte {
	row, err := db.Query(`SELECT verify FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	verify := make([]byte, 10)
	for row.Next() {
		row.Scan(&verify)
	}
	return verify
}

func Activate(verify string) {
	_, err := db.Exec(`UPDATE users SET active=1 WHERE verify="` + verify + `"`)
	if err != nil {
		log.Fatal(err)
	}
}

func IsActive(email string) bool {
	row, err := db.Query(`SELECT active FROM users WHERE email="` + email + `"`)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var active bool
	for row.Next() {
		row.Scan(&active)
	}
	return active
}
