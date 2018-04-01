package data

import (
	"database/sql"
	"ordinary-system/user/model"

	_ "github.com/go-sql-driver/mysql"
)

var dsn string = "root:123!@#qweQWE@/users"

func CreateUser(u model.User) (model.User, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return model.User{}, err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO user(firstName, lastName, username) VALUES(?,?,?)")
	if err != nil {
		return model.User{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(u.FirstName, u.LastName, u.Username)
	if err != nil {
		return model.User{}, err
	}
	u.Id, err = res.LastInsertId()
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

func GetUsers() ([]model.User, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	users := make([]model.User, 0)

	rows, err := db.Query("SELECT id, firstName, lastName, username FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		u := model.User{}
		err := rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}
