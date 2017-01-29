package main

import (
	"github.com/jasonlvhit/gocron"
	"github.com/subosito/gotenv"
	"os"
	"strconv"
)

// initiate process with chron job
func main() {
  ch := gocron.Start()
  gotenv.Load()
  frequencyString := os.Getenv("FREQUENCY")
  frequency, _ := strconv.ParseUint(frequencyString, 10, 64)
  gocron.Every(frequency).Seconds().Do(process)

  <-ch
}
