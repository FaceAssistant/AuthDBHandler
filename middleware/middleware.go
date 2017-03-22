package middleware

import (
    "net/http"
    "net/http/httputil"
    "fmt"
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
