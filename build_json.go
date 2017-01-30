package main

import (
  "encoding/json"
  "encoding/xml"
)

type jsonRecall struct {
  Recall struct {
    Date               string `json:"release_date"`
    BrandName          string `json:"name"`
    ProductDescription string `json:"product_description"`
    Reason             string `json:"reason"`
    CompanyReleaseLink string `json:"company_release_link"`
  } `json:"recall"`
}

type Product struct {
  Date               string `xml:"DATE"`
  BrandName          string `xml:"BRAND_NAME"`
  ProductDescription string `xml:"PRODUCT_DESCRIPTION"`
  Reason             string `xml:"REASON"`
  Company            string `xml:"COMPANY"`
  CompanyReleaseLink string `xml:"COMPANY_RELEASE_LINK"`
}

type productXml struct {
  Products []Product `xml:"PRODUCT"`
}

// convert xml to json
func buildJson(xmlFile []byte) string {
  var x productXml
  xml.Unmarshal(xmlFile, &x)

  var recall jsonRecall
  var recalls []jsonRecall
  for _, value := range x.Products {
    recall.Recall.Date = value.Date
    recall.Recall.BrandName = value.BrandName
    recall.Recall.ProductDescription = value.ProductDescription
    recall.Recall.Reason = value.Reason
    recall.Recall.CompanyReleaseLink = value.CompanyReleaseLink
    recalls = append(recalls, recall)
  }

  jsonData, _ := json.Marshal(recalls)
  var jsonStr = string(jsonData)
  return jsonStr
}
