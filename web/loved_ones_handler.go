package web

import (
    "net/http"
    "fa-db/model"
    "encoding/json"
)

type getLovedOneListOutput struct {
    LovedOnes []int `json:"loved_ones"`
}

type createLovedOneOutput struct {
    Id int `json:"id"`
}

func GetLovedOneHandler(db *model.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        lovedOne, err := db.GetLovedOneByID(r.FormValue("id"))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        err = json.NewEncoder(w).Encode(&lovedOne)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}

func GetLovedOnesForUserHandler(db *model.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        lovedOnes, err := db.GetAllLovedOnes(r.FormValue("user_id"))
        l := &getLovedOneListOutput{LovedOnes: lovedOnes}
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        err = json.NewEncoder(w).Encode(&l)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}

func CreateLovedOneHandler(db *model.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var l model.LovedOne
        err := json.NewDecoder(r.Body).Decode(&l)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }
        id, err := db.CreateLovedOne(&l)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        o := &createLovedOneOutput{Id: id}
        b, err := json.Marshal(&o)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.WriteHeader(http.StatusCreated)
        w.Write(b)
    }
}
