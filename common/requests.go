package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func SendPostEncoded(postUrl string, params, additionalHeaders map[string]string, filler interface{}) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	data := url.Values{}
	for k, v := range params {
		data.Add(k, v)
	}
	fmt.Printf("mockServer: %s\n", mockServer)
	if mockServer != "" {
		postUrl = mockServer
	}
	req, err := http.NewRequest("POST", postUrl, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	reqHeaders := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"Pragma":       "no-cache",
		"Sec-Ch-Ua":    `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`,
		"User-Agent":   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
	}
	for k, v := range reqHeaders {
		req.Header.Add(k, v)
	}
	for k, v := range additionalHeaders {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if filler != nil {
		// _, err = io.ReadAll(resp.Body)
		err := json.NewDecoder(resp.Body).Decode(filler)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("%s\n", flat)
	}
	fmt.Print(data.Encode(), ":(", resp.Status, ");")
}
