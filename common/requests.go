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

	// User Agent to be used in all requests
	userAgent string
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
	if r.Client != nil {
		return
	}
	r.Client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Jar: r.Jar,
	}
	r.userAgent = RandUserAgent()
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
	reqHeaders := []SimpleTerm{}
	if params != nil {
		reqHeaders = append(reqHeaders, SimpleTerm{"Content-Type", "application/x-www-form-urlencoded; charset=UTF-8"})
	}
	reqHeaders = append(reqHeaders, additionalHeaders...)
	r.doRequest("POST", postUrl, strings.NewReader(data.Encode()), reqHeaders, filler)
}

func (r *ReqHandler) SendJSON(target string, payload interface{}, additionalHeaders []SimpleTerm, filler interface{}) {
	r.checkClient()
	b := new(bytes.Buffer)
	if payload != nil {
		err := json.NewEncoder(b).Encode(payload)
		if err != nil {
			log.Fatalf("[FATAL] Data couldn't be JSON serialized: %+v\n", payload)
		}
	}
	reqHeaders := []SimpleTerm{
		{"Content-Type", "application/json"},
	}
	reqHeaders = append(reqHeaders, additionalHeaders...)
	r.doRequest("POST", target, b, reqHeaders, filler)
}

func (r *ReqHandler) SendGet(getUrl string, urlParams, additionalHeaders []SimpleTerm, filler interface{}) {
	r.checkClient()
	data := url.Values{}
	for _, v := range urlParams {
		data.Add(v.K, v.V)
	}
	coded := data.Encode()
	if CanLog(LOG_VERBOSE) {
		fmt.Println(coded, ":")
	}
	getUrl += "?" + coded
	r.doRequest("GET", getUrl, nil, additionalHeaders, filler)
}

func (r *ReqHandler) doRequest(method, urlRequest string, body io.Reader, reqHeaders []SimpleTerm, filler interface{}) {
	var err error
	var finalUrl string

	if mockServer != "" {
		finalUrl = mockServer
	} else {
		finalUrl = urlRequest
	}

	// Validate URL to be used
	if _, err := url.Parse(finalUrl); err != nil {
		log.Fatal("[FATAL] Malformed/Wrong URL:", finalUrl)
	}

	// Prepare request
	r.Request, err = http.NewRequest(method, finalUrl, body)
	if err != nil {
		log.Fatal(err)
	}

	// Add request headers
	for _, h := range reqHeaders {
		r.Request.Header.Add(h.K, h.V)
	}
	commonHeaders := []SimpleTerm{
		{"Pragma", "no-cache"},
		// {"Sec-Ch-Ua", `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`},
		{"User-Agent", r.userAgent},
	}
	for _, h := range commonHeaders {
		r.Request.Header.Add(h.K, h.V)
	}
	if mockServer != "" {
		r.Request.Header.Add("Mocked-Address", urlRequest)
	}

	// Do request. Retry at most MAX_RETRIES
	retries := uint8(0)
	const MAX_RETRIES = 10
	for {
		r.Response, err = r.Client.Do(r.Request)
		if err == nil {
			break
		}
		if retries == MAX_RETRIES {
			log.Fatalln("Request failed after", MAX_RETRIES, "retries. Error:", err)
		}
		retries++
		if CanLog(LOG_VERBOSE) {
			fmt.Printf("WARN:RET#%d ", retries)
		}
	}
	defer r.Response.Body.Close()

	// Fill supplied response object (if any)
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
