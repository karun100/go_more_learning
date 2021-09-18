package main
import (
"fmt"
"net/http"
"log"
)
func main() {
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Hello!")
    })


    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)

        helloHandler(w http.ResponseWriter, r *http.Request) {
            if r.URL.Path != "/hello" {
                http.Error(w, "404 not found.", http.StatusNotFound)
                return
            }

            if r.Method != "GET" {
                http.Error(w, "Method is not supported.", http.StatusNotFound)
                return
            }
    }



}
