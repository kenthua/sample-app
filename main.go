package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("index.html")
    t.Execute(w, "Hello World!" + os.Getenv("K8S_ENVIRONMENT"))
}

func main() {
    http.HandleFunc("/", hello)
    fmt.Println("Server running...")
    log.Fatal(http.ListenAndServe(":80", nil))
}

