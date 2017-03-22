package main

import (
    "net/http"
    "github.com/gorilla/mux"
    "log"
    "os"
    "fmt"
    "github.com/justinas/alice"
    oidc "github.com/coreos/go-oidc"
	"golang.org/x/net/context"
    "fa-db/web"
    "fa-db/model"
    "fa-db/middleware"
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
    fmt.Println("Starting")
    db, err := model.NewDB(fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName))
    if err != nil {
        panic(err)
    }
    fmt.Println("here")
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

    //Use auth middleware
    r := mux.NewRouter().StrictSlash(true)
    r.HandleFunc("/auth", web.AuthHandler(verifier, db))
    r.HandleFunc("/login", web.LoginHandler(verifier, db))
    r.HandleFunc("/loved-one", web.CreateLovedOneHandler(db)).Methods("Post")
    r.HandleFunc("/loved-one", web.GetLovedOneHandler(db)).Methods("Get").Queries("id", "{id:[0-9]+}")
    r.HandleFunc("/loved-one", web.GetLovedOnesForUserHandler(db)).Methods("Get").Queries("user_id","{id:[0-9]+}")
    fmt.Println("listening on 127.0.0.1:8080")
    a := alice.New(middleware.RequestDump).Then(r)
    log.Fatal(http.ListenAndServe(":8080", a))
}
