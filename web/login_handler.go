package web

import (
    "net/http"
    "fa-db/model"
    oidc "github.com/coreos/go-oidc"
)

func LoginHandler(v *oidc.IDTokenVerifier, db *model.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        rawIdToken := r.Header.Get("Authorization")
        email, err := AuthAndGetEmail(v, rawIdToken)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        if db.UserExist(email) {
            w.Write([]byte("User exists. OK"))
        } else {
            _, err = db.CreateUser(email)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            w.WriteHeader(http.StatusCreated)
            w.Write([]byte("User:" + email + " created."))
        }
    }
}
