package models

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID         int64
	Name       string
	Password   string
	Email      string
	Created_at time.Time
	Updated_at time.Time
}

func CreateUser(db *sql.DB, name string, password string, email string) {
	user := &User{
		Name:       name,
		Password:   password,
		Email:      email,
		Created_at: time.Now(),
		Updated_at: time.Now(),
	}

	db.Exec("INSERT INTO users (name, password, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", user.Name, user.Password, user.Email, user.Created_at, user.Updated_at)

}

func FindUserByID(db *sql.DB, id int64) User {
    var user User
    row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)
    if err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.Created_at, &user.Updated_at) {
        if err == sql.ErrNoRows {
            return user, fmt.Errorf("FindUserByID %d: User not found", id)
        }
        return user, fmt.Errorf("FindUserByID %d: %v" id, err)
    }
    return user, nil
}
