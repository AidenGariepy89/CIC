package user

import (
	"cic/site/pkg/db"
	"fmt"
)

type User struct {
	Id       int
	Username string
	Password string // This will change later
}

func GetUser(userId int) (*User, error) {
	row := db.Db.QueryRow(fmt.Sprintf("select * from user where id = %v", userId))

	var id int
	var username string
	var password string

	err := row.Scan(&id, &username, &password)
	if err != nil {
		return nil, fmt.Errorf("Error scanning user: %w\n", err)
	}

	return &User{id, username, password}, nil
}
