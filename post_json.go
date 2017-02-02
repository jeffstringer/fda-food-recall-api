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
func postJson(jsonData []byte) {
  recalls := strings.NewReader(string(jsonData))
  gotenv.Load()
  postUrl := os.Getenv("POST_URL")
  request, _ := http.Post(postUrl, "application/json", recalls)
  fmt.Println("I am posting fda json...", time.Now())
  defer request.Body.Close()
}
