package model

import (
    "strconv"
)

type LovedOne struct {
    Id int                  `db:"id"`
    Name string             `json:"name"         db:"name"`
    Birthday string         `json:"birthday"     db:"birthday"`
    Relationship string     `json:"relationship" db:"relationship"`
    Note string             `json:"note"         db:"note"`
    LastViewed string       `json:"last_viewed"  db:"last_viewed"`
    UserId int              `json:"user_id"      db:"user_id"`
}

func (db *DB) GetLovedOneByID(rawId string) (*LovedOne, error) {
    var l *LovedOne
    id, err := strconv.Atoi(rawId)
    if err != nil {
        return l, err
    }
    query := `SELECT name from loved_ones WHERE id=$1;`
    row := db.QueryRowx(query, id)

    err = row.Scan(l)
    return l, err
}

func (db *DB) CreateLovedOne(l *LovedOne) (int, error) {
    query := `INSERT INTO loved_ones(name, birthday, relationship, note, last_viewed, user_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`
    var id int
    err := db.QueryRowx(query, l.Name, l.Birthday, l.Relationship, l.Note, l.LastViewed, l.UserId).Scan(&id)
    if err != nil {
        return -1, err
    }
    return id, nil
}

func (db *DB) GetAllLovedOnes(rawUserId string) ([]int, error) {
    var lovedOneIds []int
    query := `SELECT id from loved_ones WHERE user_id=$1`
    userId, err := strconv.Atoi(rawUserId)
    if err != nil {
        return lovedOneIds, err
    }

    rows, err := db.Queryx(query, userId)
    if err != nil  {
        return lovedOneIds, err
    }

    for rows.Next() {
        var id int
        err = rows.Scan(&id)
        if err != nil {
            return lovedOneIds, err
        }
        lovedOneIds = append(lovedOneIds, id)
    }
    return lovedOneIds, err
}
