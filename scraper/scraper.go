package scraper

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const HogdonMeadowURL = "https://www.recreation.gov/camping/hodgdon-meadow/r/campgroundDetails.do?contractCode=NRSO&parkId=70929"

func Scrape() {
	//http://proxylist.hidemyass.com/search-1291967#listable
	proxyString := "http://97.77.104.22"
	proxyURL, err := url.Parse(proxyString)
	if err != nil {
		panic("Failed to parse proxy string")
	}

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    0,
		DisableCompression: true,
		Proxy:              http.ProxyURL(proxyURL),
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get(HogdonMeadowURL)
	if err != nil {
		panic(err.Error())
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	bodyString := string(bodyBytes)
	log.Println(bodyString)
}
