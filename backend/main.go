package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from Backend!")
}

func main() {
    http.HandleFunc("/", handler)
    log.Println("Backend is running on port 8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("Error starting server: %s", err)
    }
}
