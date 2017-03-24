package model

type LovedOne struct {
    Id int                  `json:"id"           db:"id"`
    Name string             `json:"name"         db:"name"`
    Birthday string         `json:"birthday"     db:"birthday"`
    Relationship string     `json:"relationship" db:"relationship"`
    Note string             `json:"note"         db:"note"`
    LastViewed string       `json:"last_viewed"  db:"last_viewed"`
    UserId string              `json:"user_id"      db:"user_id"`
}

func (db *DB) GetLovedOne(id string, userId string) (*LovedOne, error) {
    var l LovedOne
    query := `SELECT * from loved_ones WHERE id=$1 AND user_id=$2;`
    row := db.QueryRowx(query, id, userId)
    err := row.StructScan(&l)
    return &l, err
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

func (db *DB) GetAllLovedOnes(userId string) ([]int, error) {
    var lovedOneIds []int
    query := `SELECT id from loved_ones WHERE user_id=$1`

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
