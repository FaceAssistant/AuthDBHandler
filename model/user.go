package model

import (
    "database/sql"
)

type User struct {
    Id int        `json:"id"`
    Email string  `json:"email"`
}

func (db *DB) UserExist(email string) bool {
    query := `SELECT * FROM users WHERE email=$1;`
    row := db.QueryRowx(query, email)

    var u User
    err := row.Scan(&u)
    if err == sql.ErrNoRows {
        return false
    } else {
        return true
    }
}

func (db *DB) CreateUser(email string) (sql.Result, error) {
    query := `INSERT INTO users(email) VALUES ($1);`
    return db.Exec(query, email)
}

func (db *DB) GetUserId(email string) (int, error) {
    query := `SELECT id FROM users WHERE email=$1`
    row := db.QueryRowx(query, email)

    var id int
    err := row.Scan(&id)
    return id, err
}
