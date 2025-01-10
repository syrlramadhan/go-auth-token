package config

import (
	"database/sql"
	"fmt"
	"os"
	"time"
	"tugas-pemrograman-web/util"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var env = godotenv.Load()
var Host = os.Getenv("DB_HOST")
var Port = os.Getenv("DB_PORT")
var AppPort = os.Getenv("APP_PORT")
var User = os.Getenv("DB_USER")
var Pass = os.Getenv("DB_PASSWORD")
var DBName = os.Getenv("DB_NAME")

func ConnectToDatabase() (*sql.DB, error) {
	util.SentPanicIfError(env)

	mysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", User, Pass, Host, Port, DBName)
	db, err := sql.Open("mysql", mysql)
	util.SentPanicIfError(err)

	err = db.Ping()
	util.SentPanicIfError(err)

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
