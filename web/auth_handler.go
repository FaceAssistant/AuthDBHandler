package web

import (
    "net/http"
    "golang.org/x/net/context"
    oidc "github.com/coreos/go-oidc"
    "fa-db/model"
    "strconv"
    "fmt"
)

func AuthHandler(v *oidc.IDTokenVerifier, db *model.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        rawIdToken := r.Header.Get("Authorization")
        email, err := AuthAndGetEmail(v, rawIdToken)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        userId, err := db.GetUserId(email)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Authorization", strconv.Itoa(userId))
        w.WriteHeader(http.StatusOK)
    }
}

func AuthAndGetEmail(v *oidc.IDTokenVerifier, rawIdToken string) (string, error) {
    ctx := context.Background()
    idToken, err := v.Verify(ctx, rawIdToken)
    if err != nil {
        fmt.Println("Failed to verify token")
        return "", err
    }

    var claims struct {
        Email string `json:"email"`
    }

    if err = idToken.Claims(&claims); err != nil {
        return "", err
    }

    return claims.Email, nil
}
