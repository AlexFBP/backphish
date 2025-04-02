package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"
	// "unicode/utf8"
)

type SimpleTerm struct {
	K string // Key
	V string // Value
}

type ReqHandler struct {

	// Cookie jar to be used in subsequent requests
	jar *cookiejar.Jar

	// HTTP Client to be used in subsequent requests
	client *http.Client

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
		r.jar, _ = cookiejar.New(nil)
	} else {
		r.jar = nil
	}
	r.InitClient()
}

func (r *ReqHandler) PrintCookies(u *url.URL) {
	if CanLog(LOG_VERBOSE) {
		n := 0
		for _, cookie := range r.jar.Cookies(u) {
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
	if r.client != nil {
		return
	}
	r.client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	if r.jar != nil {
		r.client.Jar = r.jar
	}
	r.userAgent = RandUserAgent()
}

func (r *ReqHandler) checkClient() {
	if r.client == nil {
		r.InitClient()
	}
}

func (r *ReqHandler) SendPostEncoded(postUrl string, params, additionalHeaders []SimpleTerm, filler interface{}) (err error, lastStep uint8) {
	data := url.Values{}
	for _, v := range params {
		data.Add(v.K, v.V)
	}
	reqHeaders := []SimpleTerm{}
	if params != nil {
		reqHeaders = append(reqHeaders, SimpleTerm{"Content-Type", "application/x-www-form-urlencoded; charset=UTF-8"})
	}
	reqHeaders = append(reqHeaders, additionalHeaders...)
	return r.doRequest("POST", postUrl, strings.NewReader(data.Encode()), reqHeaders, filler)
}

func (r *ReqHandler) SendJSON(target string, payload interface{}, additionalHeaders []SimpleTerm, filler interface{}) (err error, lastStep uint8) {
	b := new(bytes.Buffer)
	if payload != nil {
		err = json.NewEncoder(b).Encode(payload)
		if err != nil {
			return
		}
	}
	reqHeaders := []SimpleTerm{
		{"Content-Type", "application/json"},
	}
	reqHeaders = append(reqHeaders, additionalHeaders...)
	return r.doRequest("POST", target, b, reqHeaders, filler)
}

func (r *ReqHandler) SendGet(getUrl string, urlParams, additionalHeaders []SimpleTerm, filler interface{}) (err error, lastStep uint8) {
	data := url.Values{}
	for _, v := range urlParams {
		data.Add(v.K, v.V)
	}
	coded := data.Encode()
	if CanLog(LOG_VERBOSE) {
		fmt.Println(coded, ":")
	}
	getUrl += "?" + coded
	return r.doRequest("GET", getUrl, nil, additionalHeaders, filler)
}

// To send multipart/form-data
func (r *ReqHandler) SendMultipart(targetURL string, values map[string]io.Reader, additionalHeaders []SimpleTerm, filler interface{}) (err error, lastStep uint8) {
	// Based on https://stackoverflow.com/a/20397167/3180052

	// Prepare a form that you will submit to that URL.
	b := &bytes.Buffer{}
	w := multipart.NewWriter(b)
	for key, rd := range values {
		var fw io.Writer
		if x, ok := rd.(io.Closer); ok {
			defer x.Close()
		}
		// Add a file
		if x, ok := rd.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return
			}
		} else {
			// Add form fields
			if fw, err = w.CreateFormField(key); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, rd); err != nil {
			return
		}

	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	if err = w.Close(); err != nil {
		return
	}

	if CanLog(LOG_VERBOSE) {
		fmt.Printf("body:\n%s\n", b.String())
	}

	additionalHeaders = append(additionalHeaders, SimpleTerm{"Content-Type", w.FormDataContentType()})
	return r.doRequest("POST", targetURL, b, additionalHeaders, filler)
}

func (r *ReqHandler) doRequest(method, urlRequest string, body io.Reader, reqHeaders []SimpleTerm, filler interface{}) (err error, lastStep uint8) {
	r.checkClient()
	var finalUrl string

	if mockServer != "" {
		finalUrl = mockServer
	} else {
		finalUrl = urlRequest
	}

	// Validate URL to be used
	if err = CheckURL(finalUrl); err != nil {
		// err = &CustErr{fmt.Sprintf("Malformed/Wrong URL: '%s'", finalUrl), err}
		return
	}
	lastStep++

	// Prepare request
	r.Request, err = http.NewRequest(method, finalUrl, body)
	if err != nil {
		return
	}
	lastStep++

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
	lastStep++

	// Do request. Retry at most MAX_RETRIES
	retries := uint8(0)
	const MAX_RETRIES = 10
	for {
		r.Response, err = r.client.Do(r.Request)
		if err == nil {
			break
		}
		if retries == MAX_RETRIES {
			// err = &CustErr{fmt.Sprint("Request failed after", MAX_RETRIES, "retries"), err}
			return
		}
		retries++
		if CanLog(LOG_VERBOSE) {
			fmt.Printf("WARN:RET#%d ", retries)
		}
		// Random delay of 1s +/+ 0.5s before next attempt
		RandDelayWindowed(time.Second, time.Second)
	}
	defer r.Response.Body.Close()
	lastStep++

	// Fill supplied response object (if any)
	if filler != nil {
		// _, err = io.ReadAll(resp.Body)
		switch t := filler.(type) {
		case *string:
			var b []byte
			b, err = io.ReadAll(r.Response.Body)
			if err != nil {
				return
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
			err = json.NewDecoder(r.Response.Body).Decode(filler)
			if err != nil {
				return
			}
		}
	}
	lastStep++

	if CanLog(LOG_VERBOSE) {
		fmt.Println("(", r.Response.Status, ")")
	}
	return
}
