package main

import "fmt"
import "net/http"

func main() {
  fmt.Println("Hello World")
  http.HandleFunc("/", handleHomeRoute)

  http.ListenAndServe(":8080", nil)
}

func handleHomeRoute(w http.ResponseWriter, request *http.Request) {
  if request.Method == "GET" {
    fmt.Fprintf(w, "Hello World")
  } else if request.Method == "POST" {
    fmt.Fprintf(w, "Hello World")
  } else {
    w.WriteHeader(405)
  }
}