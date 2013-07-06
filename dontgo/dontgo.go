package dontgo

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/dontgo", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "/dontgo/")
}