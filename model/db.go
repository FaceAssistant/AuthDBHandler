package model

import (
    _ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

type DB struct {
    *sqlx.DB
}

func NewDB(dataSourceName string) (*DB, error) {
    db, err := sqlx.Connect("postgres", dataSourceName)
    if err != nil {
        return nil, err
    }

    return &DB{db}, nil
}
