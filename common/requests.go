package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type ReqHandler struct {

	// Cookie jar to be used in subsequent requests
	Jar *cookiejar.Jar

	// HTTP Client to be used in subsequent requests
	Client *http.Client

	// Request object of the last HTTP request
	Request *http.Request

	// Response object of the last HTTP request
	Response *http.Response
}

// Defines if a cookie jar will be used in subsequent requests.
// Also calls InitClient()
func (r *ReqHandler) UseJar(use bool) {
	if use {
		jar, _ := cookiejar.New(nil)
		r.Jar = jar
	} else {
		r.Jar = nil
	}
	r.InitClient()
}

// Initializes the HTTP client
func (r *ReqHandler) InitClient() {
	r.Client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Jar: r.Jar,
	}
}

func (r *ReqHandler) checkClient() {
	if r.Client == nil {
		r.InitClient()
	}
}

func (r *ReqHandler) SendPostEncoded(postUrl string, params, additionalHeaders map[string]string, filler interface{}) {
	r.checkClient()
	data := url.Values{}
	for k, v := range params {
		data.Add(k, v)
	}
	if mockServer != "" {
		fmt.Printf("mockServer: %s\n", mockServer)
		postUrl = mockServer
	}
	req, err := http.NewRequest("POST", postUrl, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	reqHeaders := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded; charset=UTF-8",
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
	r.Request = req
	// if filler == nil {
	// 	log.Print("no filler")
	// } else {
	// 	log.Print("with filler")
	// }
	r.doRequest(filler)
}

func (r *ReqHandler) SendJSON(target string, payload interface{}, additionalHeaders map[string]string, filler interface{}) {
	r.checkClient()
	if mockServer != "" {
		fmt.Printf("mockServer: %s\n", mockServer)
		target = mockServer
	}
	b := new(bytes.Buffer)
	if payload != nil {
		err := json.NewEncoder(b).Encode(payload)
		if err != nil {
			log.Fatalf("[FATAL] Data couldn't be JSON serialized: %+v\n", payload)
		}
	}
	req, err := http.NewRequest("POST", target, b)
	if err != nil {
		log.Fatal(err)
	}
	reqHeaders := map[string]string{
		"Content-Type": "application/json",
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
	r.Request = req
	// if filler == nil {
	// 	log.Print("no filler")
	// } else {
	// 	log.Print("with filler")
	// }
	r.doRequest(filler)
}

func (r *ReqHandler) SendGet(getUrl string, params, additionalHeaders map[string]string, filler interface{}) {
	r.checkClient()
	data := url.Values{}
	for k, v := range params {
		data.Add(k, v)
	}
	if mockServer != "" {
		fmt.Printf("mockServer: %s\n", mockServer)
		getUrl = mockServer
	}
	coded := data.Encode()
	// fmt.Print(coded, ":")
	getUrl += "?" + coded
	if _, err := url.Parse(getUrl); err != nil {
		log.Fatal("[FATAL] Malformed/Wrong URL:", getUrl)
	}
	req, err := http.NewRequest("GET", getUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	reqHeaders := map[string]string{
		"Pragma":     "no-cache",
		"Sec-Ch-Ua":  `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`,
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
	}
	for k, v := range reqHeaders {
		req.Header.Add(k, v)
	}
	for k, v := range additionalHeaders {
		req.Header.Add(k, v)
	}
	r.Request = req
	// if filler == nil {
	// 	log.Print("no filler")
	// } else {
	// 	log.Print("with filler")
	// }
	r.doRequest(filler)
}

func (r *ReqHandler) doRequest(filler interface{}) {
	resp, err := r.Client.Do(r.Request)
	r.Response = resp
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// if filler == nil {
	// 	log.Print("do req - no filler")
	// } else {
	// 	log.Print("do req - with filler")
	// }
	if filler != nil {
		// _, err = io.ReadAll(resp.Body)
		err := json.NewDecoder(resp.Body).Decode(filler)
		if err != nil {
			log.Fatal("[json decode]", err)
		}
		// fmt.Printf("%s\n", flat)
	}
	fmt.Println("(", resp.Status, ")")
}
