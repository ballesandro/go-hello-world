package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", helloWorld)
    error := http.ListenAndServe(":" + os.Getenv("PORT"), nil)
    if error != nil {
        log.Fatal("ListenAndServe: ", error)
    }
}

func helloWorld(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "Hello, Predix world!")
}