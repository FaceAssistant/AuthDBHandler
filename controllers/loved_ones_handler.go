package controllers

import (
    "net/http"
    "fa-db/models"
    "encoding/json"
)

type getLovedOneInput struct {
    Id int `json:"id"`
}

type getLovedOneOutput struct {
    Name string `json:"name"`
}

type createLovedOneInput struct {
    Name string `json:"name"`
}

func GetLovedOneHandler(db *models.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var i getLovedOneInput
        err := json.NewDecoder(r.Body).Decode(&i)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        name, err := db.GetLovedOneByID(i.Id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        var o getLovedOneInput
        err = json.NewEncoder(w).Encode(&o)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}

func CreateLovedOneHandler(db *models.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var i createLovedOneInput
        err := json.NewDecoder(r.Body).Decode(&i)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        _, err = db.CreateLovedOne(i.Name)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        w.WriteHeader(http.StatusCreated)
    }
}
