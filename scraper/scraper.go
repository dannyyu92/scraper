package scraper

import (
	"log"
	"net/http"
	"net/url"
)

func Scrape() {
	proxyString := "https://47.91.179.xxx:443"
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
	resp, err := client.Get("https://example.com")
	if err != nil {
		panic(err.Error())
	}
	log.Println(resp.Body)
}
