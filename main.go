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
func postRecallJson() {
	fdaUrl := "http://www.fda.gov/DataSets/Recalls/Food/Food.xml"
	response, err := http.Get(fdaUrl)
	xmlFile, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if response.Status == "200 OK" {
		var r Recall
		xml.Unmarshal(xmlFile, &r)

		// convert to JSON
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

		// POST to rails app
		gotenv.Load()
		postUrl := os.Getenv("POST_URL")
		var jsonStr = string(jsonData)
		request, _ := http.Post(postUrl, "application/json", strings.NewReader(jsonStr))
		fmt.Println("I am getting fda data...", time.Now())
		defer request.Body.Close()
	}
	if err != nil {
		panic(err.Error())
	}
}

// initiate process with chron job
func main() {
    ch := gocron.Start()
    gocron.Every(10).Seconds().Do(postRecallJson)

    <-ch
}
