package dp01

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand" // "crypto/rand"
	"net/http"

	"github.com/AlexFBP/backphish/common"
)

type SpamBody struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func Cmd1(args ...string) error {
	return common.AttackRunner(attempt)
}

func attempt() {
	chat_id := "6332256769"
	tiposDocumento := []string{
		"tarjeta de identidad",
		"cedula de extranjeria",
		"cedula de ciudadania",
	}
	randIp := common.GeneraIP()
	location := "Bogotá, CO"

	// JSON body
	d := SpamBody{
		ChatID: chat_id,
		Text: fmt.Sprintf(`DATOS DAVPLAT
TipoDoc: %s
NumDoc: %d
Clave: %d
IP: %s
%s`,
			tiposDocumento[rand.Intn(len(tiposDocumento))],
			common.GeneraNIPcolombia(),
			common.GeneraPin(4),
			randIp, location),
		// DATOS DAVPLAT
		// TipoDoc: cedula de ciudadania
		// NumDoc: 1010111010
		// Clave: 2365
		// IP: 123.123.123.123
		// Bogotá, CO
	}

	d2 := SpamBody{
		ChatID: chat_id,
		Text: fmt.Sprintf(`DATOS DAVPLAT
Cod1: %d
IP: %s
%s`, common.GeneraPin(6), randIp, location),
		// `DATOS DAVPLAT
		// Cod1: 258415
		// IP: 123.123.123.123
		// Bogotá, CO`,
	}

	sendReq(d)
	common.RandDelay(30, 60)
	sendReq(d2)
	common.RandDelay(30, 60)
}

func sendReq(v any) {

	// HTTP endpoint
	// posturl := "https://jsonplaceholder.typicode.com/posts"
	// posturl := "http://localhost:8080/"
	posturl := "https://api.telegram.org/bot6328508246:AAFU2THEk8nvxuzFLR5C6FfqQNuxmsiaFWk/sendMessage"

	reqBody, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", posturl, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	headers := map[string]string{
		"Content-Type": "application/json",
		"Origin":       "https://ingressar1davidd.sayo1296.repl.co",
		"Pragma":       "no-cache",
		"Referer":      "https://ingressar1davidd.sayo1296.repl.co/",
		"Sec-Ch-Ua":    `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`,
		"User-Agent":   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
	}
	for k, v := range headers {
		req.Header.Set(k, v)
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
