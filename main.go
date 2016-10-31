package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type jsonProduct struct {
	Date               string `json:"release_date"`
	BrandName          string `json:"brand_name"`
	ProductDescription string `json:"product_description"`
	Reason             string `json:"reason"`
	Company            string `json:"company"`
	CompanyReleaseLink string `json:"company_release_link"`
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

func main() {
	fda_url := "http://www.fda.gov/DataSets/Recalls/Food/Food.xml"
	response, err := http.Get(fda_url)
	body, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if response.Status == "200 OK" {
		var r Recall
		err = xml.Unmarshal(body, &r)
		if err != nil {
			panic(err.Error())
		}

		// convert to JSON
		var oneProduct jsonProduct
		var allProducts []jsonProduct

		for _, value := range r.Products {
			oneProduct.Date = value.Date
			oneProduct.BrandName = value.BrandName
			oneProduct.ProductDescription = value.ProductDescription
			oneProduct.Reason = value.Reason
			oneProduct.Company = value.Company
			oneProduct.CompanyReleaseLink = value.CompanyReleaseLink
			allProducts = append(allProducts, oneProduct)
		}

		jsonData, err := json.Marshal(allProducts)
		if err != nil {
			fmt.Println(err)
		}

		sinatra_url := os.Getenv("SINATRA_URL")
		var jsonStr = string(jsonData)
		request, err := http.Post(sinatra_url, "application/json", strings.NewReader(jsonStr))
		if err != nil {
			// handle error
			println(err)
		}
		defer request.Body.Close()
	}
	if err != nil {
		panic(err.Error())
	}
}
