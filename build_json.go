package main

import (
  "encoding/json"
  "encoding/xml"
)

/**
"recall"=>
  {
    "release_date"=>"Sat, 28 Jan 2017 15:09:00 -0500",
    "name"=>"ageLOC TR90",
    "product_description"=>"Protein boost",
    "reason"=>"Undeclared milk",
    "company_release_link"=>"http://www.fda.gov/Safety/Recalls/ucm538725.htm"
  }
}
**/

type jsonRecall struct {
  Recall struct {
    Date               string `json:"release_date"`
    BrandName          string `json:"name"`
    ProductDescription string `json:"product_description"`
    Reason             string `json:"reason"`
    CompanyReleaseLink string `json:"company_release_link"`
  } `json:"recall"`
}

/**
<PRODUCT>
  <DATE>Sat, 28 Jan 2017 15:09:00 -0500</DATE>
  <BRAND_NAME>
    <![CDATA[ ageLOC TR90 ]]>
  </BRAND_NAME>
  <PRODUCT_DESCRIPTION>
    <![CDATA[ Protein boost ]]>
  </PRODUCT_DESCRIPTION>
  <REASON>
    <![CDATA[ Undeclared milk ]]>
  </REASON>
  <COMPANY>
    <![CDATA[ NSE Products, Inc. ]]>
  </COMPANY>
  <COMPANY_RELEASE_LINK>http://www.fda.gov/Safety/Recalls/ucm538725.htm</COMPANY_RELEASE_LINK>
  <PHOTOS_LINK></PHOTOS_LINK>
</PRODUCT>
**/

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
