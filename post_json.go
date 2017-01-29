package main

import (
  "fmt"
  "github.com/subosito/gotenv"
  "net/http"
  "os"
  "strings"
  "time"
)

// POST json to rails app
func postJson(jsonStr string) {
  str := strings.NewReader(jsonStr)
  gotenv.Load()
  postUrl := os.Getenv("POST_URL")
  request, _ := http.Post(postUrl, "application/json", str)
  fmt.Println("I am posting fda json...", time.Now())
  defer request.Body.Close()
}
