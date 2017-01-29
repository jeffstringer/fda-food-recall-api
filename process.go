package main

func process() {
  xml := getFdaXml()
  recallsJson := buildJson(xml)
  postJson(recallsJson)
}