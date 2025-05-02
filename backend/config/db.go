package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func ConnectDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't find .env file")
	}

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUser")
	cfg.Passwd = os.Getenv("DBPass")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = os.Getenv("DBName")

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("[Database] Connection Successful!")
}
