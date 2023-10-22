package bc01

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/AlexFBP/backphish/common"
)

func Cmd(args ...string) error {
	return common.AttackRunner(attempt)
	// For single attack, comment line above and uncomment this:
	// attempt()
	// return nil
}

func attempt() {
	common.RandDelay(3, 10)
	sendReq(
		// "http://localhost:1080",
		"https://desbloqueo--sucursalvirtua2.repl.co/finish9.php",
		map[string]string{"cedula": fmt.Sprint(common.GeneraNIPcolombia())},
		map[string]string{
			"Host":    "desbloqueo--sucursalvirtua2.repl.co",
			"Origin":  "https://desbloqueo--sucursalvirtua2.repl.co",
			"Referer": "https://desbloqueo--sucursalvirtua2.repl.co/index.html",
		},
	)

	common.RandDelay(2, 5)
	sendReq(
		// "http://localhost:1080",
		"https://activacion--vitualclave.repl.co/finish9.php",
		map[string]string{"clave": fmt.Sprintf("%04d", common.GeneraPin(4))},
		map[string]string{
			"Host":    "activacion--vitualclave.repl.co",
			"Origin":  "https://activacion--vitualclave.repl.co",
			"Referer": "https://activacion--vitualclave.repl.co/index.html",
		},
	)

	common.RandDelay(12, 51)
	sendReq(
		// "http://localhost:1080",
		"https://dinamica.vitualclave.repl.co/finish9.php",
		map[string]string{"clave": fmt.Sprintf("%06d", common.GeneraPin(6))},
		map[string]string{
			"Host":    "dinamica.vitualclave.repl.co",
			"Origin":  "https://dinamica.vitualclave.repl.co",
			"Referer": "https://dinamica.vitualclave.repl.co/index.html",
		},
	)

	// POST https://oblongmajorblocks--mamiamia.repl.co/finish.php
	// tipoCC=master&codigo=5327500310169812&mes=3&a%C3%B1o=2029&cvv=176&cc=&ciudad=&dir=&tel=

	fmt.Println()
}

func sendReq(postUrl string, params, additionalHeaders map[string]string) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	data := url.Values{}
	for k, v := range params {
		data.Add(k, v)
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
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s\n", flat)
	fmt.Print(data.Encode(), ":(", resp.Status, ");")
}
