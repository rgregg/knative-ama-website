package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
)

func main() {
  fs := http.FileServer(http.Dir("./static"))
  http.Handle("/", fs)

  port := os.Getenv("PORT")
  if port == "" {
          port = "8080"
  }

  log.Println(fmt.Sprintf("Listening on :%s...", port))
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}