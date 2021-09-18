package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.HandleFunc("/bye", Bye)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func Bye(w http.ResponseWriter, r *http.Request) {

fmt.Fprintf(w, "TC, %s!", r.URL.Path[1:])
}
