package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "log"
    "os"
    "fmt"
    oidc "github.com/coreos/go-oidc"
	"golang.org/x/net/context"
    "fa-db/controllers"
    "fa-db/models"
)

var (
    clientID = os.Getenv("OAUTH2_CLIENT_ID")
    clientSecret = os.Getenv("OAUTH2_CLIENT_SECRET")

    dbUser = os.Getenv("DB_USER")
    dbPass = os.Getenv("DB_PASS")
    dbName = os.Getenv("DB_NAME")
    dbPort = os.Getenv("DB_PORT")
    dbHost = os.Getenv("DB_HOST")
)

func main() {
    db, err := models.NewDB(fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName))
    if err != nil {
        panic(err)
    }
    ctx := context.Background()
    provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
    if err != nil {
        log.Fatal(err)
    }

    oidcConfig := &oidc.Config{
        ClientID: clientID,
        SkipNonceCheck: true,
    }
    verifier := provider.Verifier(oidcConfig)

    r := mux.NewRouter().StrictSlash(true)
    r.HandleFunc("/auth", controllers.AuthHandler(verifier, db))
    r.HandleFunc("/create-user", controllers.CreateUserHandler(db))
    fmt.Println("listening on 127.0.0.1:8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
