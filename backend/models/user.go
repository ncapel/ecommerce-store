package models

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/ncapel/ecommerce-store/utils"
)

type User struct {
	ID        int64
	Name      string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserPatch struct {
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

func CreateUser(db *sql.DB, name string, password string, email string) (int64, error) {
	user := &User{
		Name:      name,
		Password:  password,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	hashedPassword := utils.HashAndSalt([]byte(user.Password))
	result, err := db.Exec("INSERT INTO user (name, password, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", user.Name, hashedPassword, user.Email, user.CreatedAt, user.UpdatedAt)
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
	row := db.QueryRow("SELECT * FROM user WHERE id = ?", id)
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

	_, err = db.Exec("DELETE FROM user WHERE id = ?", user.ID)
	if err != nil {
		return 0, user.Name, fmt.Errorf("DeleteUser: %v", err)
	}

	return user.ID, "Deleted: " + user.Name, nil
}

func UpdateUserPatch(db *sql.DB, id int64, p UserPatch) error {
	if db == nil {
		return fmt.Errorf("db is nil")
	}

	sets := []string{}
	args := []interface{}{}

	if p.Name != nil {
		sets = append(sets, "name = ?")
		args = append(args, *p.Name)
	}
	if p.Email != nil {
		sets = append(sets, "email = ?")
		args = append(args, *p.Email)
	}
	if p.Password != nil {
		hashedPassword := utils.HashAndSalt([]byte(*p.Password))
		sets = append(sets, "password = ?")
		args = append(args, hashedPassword)
	}

	sets = append(sets, "updated_at = ?")
	args = append(args, time.Now())

	if len(sets) == 1 {
		return nil
	}

	query := fmt.Sprintf(
		"UPDATE user SET %s WHERE id = ?",
		strings.Join(sets, ", "),
	)
	args = append(args, id)

	res, err := db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("UpdateUserPatch exec: %v", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("UpdateUserPatch rows: %v", err)
	}
	if rows == 0 {
		return fmt.Errorf("no user found with id %d", id)
	}
	return nil
}
