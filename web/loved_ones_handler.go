package web

import (
    "net/http"
    "fa-db/model"
    "encoding/json"
)

type getLovedOneListOutput struct {
    LovedOnes []string `json:"loved_ones"`
}

type getProfileListOutput struct {
    Profiles []model.LovedOne `json:"profiles"`
}

type createLovedOneOutput struct {
    Id string `json:"id"`
}

func GetLovedOneHandler(db *model.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if userId, ok := r.Context().Value("uid").(string); ok {
            lovedOne, err := db.GetLovedOne(r.FormValue("id"), userId)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            err = json.NewEncoder(w).Encode(&lovedOne)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
        } else {
            http.Error(w, "Failed to get subject from context", http.StatusInternalServerError)
            return
        }
    }
}

func GetLovedOnesListHandler(db *model.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if userId, ok := r.Context().Value("uid").(string); ok {
            if r.FormValue("type") != "profile" {
                lovdOnesId, err := db.GetAllLovedOnesId(userId)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }

                if lovdOnesId == nil {
                    lovdOnesId = make([]string, 0)
                }

                l := &getLovedOneListOutput{LovedOnes: lovdOnesId}
                err = json.NewEncoder(w).Encode(&l)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
            } else {
                profiles, err := db.GetAllLovedOnesProfile(userId)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }

                if profiles == nil {
                    profiles = make([]model.LovedOne, 0)
                }

                o := &getProfileListOutput{Profiles: profiles}
                err = json.NewEncoder(w).Encode(&o)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
            }
        } else {
            http.Error(w, "Failed to get subject from context", http.StatusInternalServerError)
            return
        }
    }
}

func CreateLovedOneHandler(db *model.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if userId, ok := r.Context().Value("uid").(string); ok {
            var l model.LovedOne
            err := json.NewDecoder(r.Body).Decode(&l)
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }

            l.UserId = userId
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

        } else {
            http.Error(w, "Failed to get subject from context", http.StatusInternalServerError)
            return
        }
    }
}

func DeleteLovedOneHandler(db *model.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if userId, ok := r.Context().Value("uid").(string); ok {
            err := db.DeleteLovedOne(r.FormValue("id"), userId)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
        } else {
            http.Error(w, "Failed to get subject from context", http.StatusInternalServerError)
            return
        }
    }
}
