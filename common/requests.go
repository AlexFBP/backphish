package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	// "unicode/utf8"
)

type SimpleTerm struct {
	K string // Key
	V string // Value
}

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
		r.Jar, _ = cookiejar.New(nil)
	} else {
		r.Jar = nil
	}
	r.InitClient()
}

func (r *ReqHandler) PrintCookies(u *url.URL) {
	if CanLog(LOG_VERBOSE) {
		n := 0
		for _, cookie := range r.Jar.Cookies(u) {
			fmt.Printf("%s\t%s\n", cookie.Name, cookie.Value)
			n++
		}
		if n == 0 {
			fmt.Printf("No cookies for %s\n", u.String())
		}
	}
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

func (r *ReqHandler) SendPostEncoded(postUrl string, params, additionalHeaders []SimpleTerm, filler interface{}) {
	r.checkClient()
	data := url.Values{}
	for _, v := range params {
		data.Add(v.K, v.V)
	}
	if mockServer != "" {
		if CanLog(LOG_VERBOSE) {
			fmt.Printf("mockServer: %s\n", mockServer)
		}
		postUrl = mockServer
	}
	var err error
	r.Request, err = http.NewRequest("POST", postUrl, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	reqHeaders := []SimpleTerm{
		{"Pragma", "no-cache"},
		{"Sec-Ch-Ua", `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`},
		{"User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"},
	}
	if params != nil {
		reqHeaders = append(reqHeaders, SimpleTerm{"Content-Type", "application/x-www-form-urlencoded; charset=UTF-8"})
	}
	if additionalHeaders != nil {
		reqHeaders = append(reqHeaders, additionalHeaders...)
	}
	r.doRequest(reqHeaders, filler)
}

func (r *ReqHandler) SendJSON(target string, payload interface{}, additionalHeaders []SimpleTerm, filler interface{}) {
	r.checkClient()
	if mockServer != "" {
		if CanLog(LOG_VERBOSE) {
			fmt.Printf("mockServer: %s\n", mockServer)
		}
		target = mockServer
	}
	b := new(bytes.Buffer)
	if payload != nil {
		err := json.NewEncoder(b).Encode(payload)
		if err != nil {
			log.Fatalf("[FATAL] Data couldn't be JSON serialized: %+v\n", payload)
		}
	}
	var err error
	r.Request, err = http.NewRequest("POST", target, b)
	if err != nil {
		log.Fatal(err)
	}
	reqHeaders := []SimpleTerm{
		{"Content-Type", "application/json"},
		{"Pragma", "no-cache"},
		{"Sec-Ch-Ua", `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`},
		{"User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"},
	}
	if additionalHeaders != nil {
		reqHeaders = append(reqHeaders, additionalHeaders...)
	}
	r.doRequest(reqHeaders, filler)
}

func (r *ReqHandler) SendGet(getUrl string, urlParams, additionalHeaders []SimpleTerm, filler interface{}) {
	r.checkClient()
	data := url.Values{}
	for _, v := range urlParams {
		data.Add(v.K, v.V)
	}
	if mockServer != "" {
		if CanLog(LOG_VERBOSE) {
			fmt.Printf("mockServer: %s\n", mockServer)
		}
		getUrl = mockServer
	}
	coded := data.Encode()
	if CanLog(LOG_VERBOSE) {
		fmt.Println(coded, ":")
	}
	getUrl += "?" + coded
	if _, err := url.Parse(getUrl); err != nil {
		log.Fatal("[FATAL] Malformed/Wrong URL:", getUrl)
	}
	var err error
	r.Request, err = http.NewRequest("GET", getUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	reqHeaders := []SimpleTerm{
		{"Pragma", "no-cache"},
		{"Sec-Ch-Ua", `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`},
		{"User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36"},
	}
	if additionalHeaders != nil {
		reqHeaders = append(reqHeaders, additionalHeaders...)
	}
	r.doRequest(reqHeaders, filler)
}

func (r *ReqHandler) doRequest(reqHeaders []SimpleTerm, filler interface{}) {
	for _, h := range reqHeaders {
		r.Request.Header.Add(h.K, h.V)
	}
	var err error
	r.Response, err = r.Client.Do(r.Request)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Response.Body.Close()
	if filler != nil {
		// _, err = io.ReadAll(resp.Body)
		switch t := filler.(type) {
		case *string:
			b, err := io.ReadAll(r.Response.Body)
			if err != nil {
				log.Fatal(err)
			}
			*t = string(b)
			// for len(b) > 0 {
			// 	r, size := utf8.DecodeRune(b)
			// 	if CanLog(LOG_VERBOSE) {
			// 		fmt.Print(r)
			// 	}
			// 	b = b[size:]
			// }
			if CanLog(LOG_VERBOSE) {
				log.Printf("plain:%s - quoted:%+q\n", b, b)
			}
		default:
			err := json.NewDecoder(r.Response.Body).Decode(filler)
			if err != nil {
				log.Fatalf("[decode] Not a JSON response or unhandled filler type '%T' - Error: %v\n", t, err)
			}
		}
	}
	if CanLog(LOG_VERBOSE) {
		fmt.Println("(", r.Response.Status, ")")
	}
}
