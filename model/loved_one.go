package model

type LovedOne struct {
    Id string               `json:"id"           db:"id"`
    Name string             `json:"name"         db:"name"`
    Birthday string         `json:"birthday"     db:"birthday"`
    Relationship string     `json:"relationship" db:"relationship"`
    Note string             `json:"note"         db:"note"`
    LastViewed string       `json:"last_viewed"  db:"last_viewed"`
    UserId string           `json:"user_id"      db:"user_id"`
}

func (db *DB) GetLovedOne(id string, userId string) (*LovedOne, error) {
    var l LovedOne
    query := `SELECT * from loved_ones WHERE id=$1 AND user_id=$2;`
    row := db.QueryRowx(query, id, userId)
    err := row.StructScan(&l)
    return &l, err
}

func (db *DB) CreateLovedOne(l *LovedOne) (string, error) {
    query := `INSERT INTO loved_ones(id, name, birthday, relationship, note, last_viewed, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`
    var id string
    err := db.QueryRowx(query, l.Id, l.Name, l.Birthday, l.Relationship, l.Note, l.LastViewed, l.UserId).Scan(&id)
    if err != nil {
        return "", err
    }
    return id, nil
}

func (db *DB) GetAllLovedOnes(userId string) ([]string, error) {
    var lovedOneIds []string
    query := `SELECT id from loved_ones WHERE user_id=$1`

    rows, err := db.Queryx(query, userId)
    if err != nil  {
        return lovedOneIds, err
    }

    for rows.Next() {
        var id string
        err = rows.Scan(&id)
        if err != nil {
            return lovedOneIds, err
        }
        lovedOneIds = append(lovedOneIds, id)
    }
    return lovedOneIds, err
}

func (db *DB) DeleteLovedOne(id string, userId string) error {
    query := `DELETE FROM loved_ones WHERE id=$1 and userId=$2;`
    _, err := db.Exec(query, id, userId)
    if err != nil {
        return err
    }
    return nil
}
