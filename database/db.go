package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var (
	user     string
	password string
	host     string
	port     string
	dbname   string

	db  *sql.DB
	err error
)

func SetDBEnvironment() {
	user = viper.GetString("database.user")
	password = viper.GetString("database.password")
	host = viper.GetString("database.host")
	port = viper.GetString("database.port")
	dbname = viper.GetString("database.dbname")
}

func Connect() {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	db, err = sql.Open("mysql", conn)
	if err != nil {
		log.Fatal("Error when connecting to database\n" + err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error when connecting to database\n" + err.Error())
	}

	log.Println("Success connecting to database")
}

func DB() *sql.DB {
	return db
}
