package models

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID        int64
	Name      string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateUser(db *sql.DB, name string, password string, email string) (int64, error) {
	user := &User{
		Name:      name,
		Password:  password,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := db.Exec("INSERT INTO users (name, password, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", user.Name, user.Password, user.Email, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return 0, fmt.Errorf("CreateUser: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("CreateUser: %v", err)
	}

	return id, nil
}

func FindUserByID(db *sql.DB, id int64) (User, error) {
	var user User
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	if err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("FindUserByID %d: User not found", id)
		}
		return user, fmt.Errorf("FindUserByID %d: %v", id, err)
	}
	return user, nil
}

func DeleteUser(db *sql.DB, u User) (int64, string, error) {
	user, err := FindUserByID(db, u.ID)
	if err != nil {
		return 0, "", err
	}

	_, err = db.Exec("DELETE FROM users WHERE id = ?", user.ID)
	if err != nil {
		return 0, user.Name, fmt.Errorf("DeleteUser: %v", err)
	}

	return user.ID, "Deleted: " + user.Name, nil
}
