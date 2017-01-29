package main

import (
  "io/ioutil"
  "net/http"
)

// GET to fda dataset site
func getFdaXml() []byte {
  fdaUrl := "http://www.fda.gov/DataSets/Recalls/Food/Food.xml"
  response, _ := http.Get(fdaUrl)
  xmlFile, _ := ioutil.ReadAll(response.Body)
  defer response.Body.Close()
  return xmlFile
}