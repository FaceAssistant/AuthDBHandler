package web

import (
    "net/http"
    "fa-db/model"
    "encoding/json"
)

//NEED TO UPDATE
type getLovedOneOutput struct {
    Name string `json:"name"`
}

type createLovedOneInput struct {
    Name string `json:"name"`
}

func GetLovedOneHandler(db *model.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        name, err := db.GetLovedOneByID(r.FormValue("id"))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        o := &getLovedOneOutput{
            Name: name,
        }
        err = json.NewEncoder(w).Encode(&o)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}

//WRITE ID FROM SQL RESULT
func CreateLovedOneHandler(db *model.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var i createLovedOneInput
        err := json.NewDecoder(r.Body).Decode(&i)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        result, err = db.CreateLovedOne(i.Name)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusCreated)
    }
}
