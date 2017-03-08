package web

import (
    "net/http"
    "fa-db/model"
    "encoding/json"
)

type CreateUserInput struct {
    Email string `json:"email"`
}

func CreateUserHandler(db *model.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var i CreateUserInput
        err := json.NewDecoder(r.Body).Decode(&i)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        if db.UserExist(i.Email) {
            w.Write([]byte("User exists. OK"))
        } else {
            _, err = db.CreateUser(i.Email)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            w.WriteHeader(http.StatusCreated)
            w.Write([]byte("User:" + i.Email + " created."))
        }
    }
}
