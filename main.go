package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/subosito/gotenv"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type jsonProduct struct {
	Recall struct {
		Date               string `json:"release_date"`
		BrandName          string `json:"name"`
		ProductDescription string `json:"product_description"`
		Reason             string `json:"reason"`
		CompanyReleaseLink string `json:"company_release_link"`
	}	`json:"recall"`
}

type Product struct {
	Date               string `xml:"DATE"`
	BrandName          string `xml:"BRAND_NAME"`
	ProductDescription string `xml:"PRODUCT_DESCRIPTION"`
	Reason             string `xml:"REASON"`
	Company            string `xml:"COMPANY"`
	CompanyReleaseLink string `xml:"COMPANY_RELEASE_LINK"`
}

type Recall struct {
	Products []Product `xml:"PRODUCT"`
}

// GET to fda dataset site
func getFdaXml() []byte {
	fdaUrl := "http://www.fda.gov/DataSets/Recalls/Food/Food.xml"
	response, _ := http.Get(fdaUrl)
	xmlFile, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return xmlFile
}

// convert xml to json
func buildJson(xmlFile []byte) string {
		var r Recall
		xml.Unmarshal(xmlFile, &r)

		var oneProduct jsonProduct
		var allProducts []jsonProduct
		for _, value := range r.Products {
			oneProduct.Recall.Date = value.Date
			oneProduct.Recall.BrandName = value.BrandName
			oneProduct.Recall.ProductDescription = value.ProductDescription
			oneProduct.Recall.Reason = value.Reason
			oneProduct.Recall.CompanyReleaseLink = value.CompanyReleaseLink
			allProducts = append(allProducts, oneProduct)
		}

		jsonData, _ := json.Marshal(allProducts)
		var jsonStr = string(jsonData)
		return jsonStr
	}

// // POST json to rails app
func postJson(jsonStr string) {
	str := strings.NewReader(jsonStr)
	gotenv.Load()
	postUrl := os.Getenv("POST_URL")
	request, _ := http.Post(postUrl, "application/json", str)
	fmt.Println("I am getting fda data...", time.Now())
	defer request.Body.Close()
}


func process() {
	xml := getFdaXml()
	recallsJson := buildJson(xml)
	postJson(recallsJson)
}

// initiate process with chron job
func main() {
  ch := gocron.Start()
  gotenv.Load()
  frequencyString := os.Getenv("FREQUENCY")
  frequency, _ := strconv.ParseUint(frequencyString, 10, 64)
  gocron.Every(frequency).Minutes().Do(process)

  <-ch
}
