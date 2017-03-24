package middleware

import (
    "net/http"
    "net/http/httputil"
    oidc "github.com/coreos/go-oidc"
    "fmt"
    "context"
)

func RequestDump(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        dump, err := httputil.DumpRequest(r, true)
        if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(string(dump))
		next.ServeHTTP(w, r)
    })
}

func AuthRequest(v *oidc.IDTokenVerifier) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            rawIdToken := r.Header.Get("Authorization")
            sub, err := authUser(v, rawIdToken)
            if err != nil {
                http.Error(w, err.Error(), http.StatusUnauthorized)
                return
            }
            newCtx := context.WithValue(r.Context(), "sub", sub)
            next.ServeHTTP(w, r.WithContext(newCtx))
        })
    }
}


func authUser(v *oidc.IDTokenVerifier, rawIdToken string) (string, error) {
    ctx := context.Background()
    idToken, err := v.Verify(ctx, rawIdToken)
    if err != nil {
        fmt.Println("Failed to verify token")
        return "", err
    }

    var claims struct {
        Subject string `json:"sub"`
    }

    if err = idToken.Claims(&claims); err != nil {
        return "", err
    }

    return claims.Subject, nil
}
