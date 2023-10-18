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
	sendReq(
		// "http://localhost:1080",
		"https://desbloqueo--sucursalvirtua2.repl.co/finish9.php",
		map[string]string{"cedula": "224224242424"})
	common.RandDelay(3, 10)

	sendReq(
		// "http://localhost:1080",
		"https://activacion--vitualclave.repl.co/finish9.php",
		map[string]string{"clave": "6248"})
	common.RandDelay(2, 5)

	sendReq(
		// "http://localhost:1080",
		"https://dinamica.vitualclave.repl.co/finish9.php",
		map[string]string{"clave": "245871"})
	common.RandDelay(12, 51)
}

func sendReq(postUrl string, params map[string]string) {
	client := &http.Client{}
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
	}
	for k, v := range reqHeaders {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	flat, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", flat)
}
