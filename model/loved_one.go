package model

import (
    "database/sql"
    "strconv"
)

type LovedOne struct {
    Id int `json:"id"`
    Name string `json:"name"`
}

func (db *DB) GetLovedOneByID(rawId string) (string, error) {
    id, err := strconv.Atoi(rawId)
    if err != nil {
        return "", err
    }
    query := `SELECT name from loved_ones WHERE id=$1;`
    row := db.QueryRowx(query, id)

    var name string
    err = row.Scan(&name)
    return name, err
}

func (db *DB) CreateLovedOne(name string) (sql.Result, error) {
    query := `INSERT INTO loved_ones(name) VALUES ($1);`
    return db.Exec(query, name)
}
