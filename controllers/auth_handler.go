package controllers
import (
    "net/http"
    "encoding/json"
    "golang.org/x/net/context"
    oidc "github.com/coreos/go-oidc"
    "fa-db/models"
)

type IdToken struct {
    RawIDToken string `json:"id_token"`
}

type AuthOutput struct {
    UserId int `json:"user_id"`
}

func AuthHandler(v *oidc.IDTokenVerifier, db *models.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx := context.Background()

        var t *IdToken
        err := json.NewDecoder(r.Body).Decode(&t)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        idToken, err := v.Verify(ctx, t.RawIDToken)
        if err != nil {
			http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
			return
        }

        var claims struct {
            Email string `json:"email"`
        }

        if err = idToken.Claims(&claims); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
			return
        }

        userId, err := db.GetUserId(claims.Email)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        o := &AuthOutput{UserId:userId}
        err = json.NewEncoder(w).Encode(&o)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}
