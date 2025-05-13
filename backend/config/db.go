package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var Db *sql.DB

func initDBConfig() *mysql.Config {
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
	return cfg
}

var Cfg = initDBConfig()

func ConnectDB() {
	cfg := Cfg

	var err error
	Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := Db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("[Database] Connection Successful!")
}

func SeedDb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't find .env file")
	}

	cfg := Cfg

	Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := Db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	seedSQL := "INSERT INTO `user` (name, email, password) VALUES " +
		"('Alice', 'alice@example.com', 'hashedpassword1'), " +
		"('Bob', 'bob@example.com', 'hashedpassword2');" +

		"INSERT INTO `product` (name, `desc`, price) VALUES " +
		"('T-Shirt', 'Comfortable cotton t-shirt', 19.99), " +
		"('Coffee Mug', 'Ceramic mug with a cool design', 9.99), " +
		"('Notebook', 'Hardcover lined notebook', 12.50);" +

		"INSERT INTO `order` (user_id) VALUES (1), (2);" +

		"INSERT INTO `order_items` (order_id, product_id, quantity) VALUES " +
		"(1, 1, 2), " +
		"(1, 2, 1), " +
		"(2, 3, 3);"

	_, err = Db.Exec(seedSQL)
	if err != nil {
		log.Fatalf("Error seeding DB: %v", err)
	}

	log.Println("Database seeded successfully!")
}
