package main

import (
    "fmt"
    "net/http"
    "os"
    "errors"
)

func main() {
    http.HandleFunc("/", Handler)
    http.ListenAndServe(":9090", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
    file, err := os.OpenFile("/opt/ready.txt", os.O_RDONLY, 0644)
    if errors.Is(err, os.ErrNotExist) {
        http.Error(w, "false", http.StatusBadRequest)
    } else {
        fmt.Fprintf(w, "true")
    }
    file.Close()
}