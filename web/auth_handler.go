package web

import (
    "net/http"
    "encoding/json"
)

type AuthOutput struct {
    UserId string `json:"user_id"`
}

func AuthHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if userId, ok := r.Context().Value("uid").(string); ok {
            o := &AuthOutput{UserId: userId}
            err := json.NewEncoder(w).Encode(&o)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
        } else  {
            http.Error(w, "Failed to get user id from context.", http.StatusInternalServerError)
            return
        }
    }
}

