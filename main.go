package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand" // "crypto/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type SpamBody struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func main() {
	rand.Seed(time.Now().UnixMilli())
	attempts := 0

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Print("Total Attempts: ")
		log.Println(attempts)
		os.Exit(0)
	}()

	for ; ; attempts++ {
		attempt()
	}
}

func attempt() {
	chat_id := "6332256769"
	tiposDocumento := []string{
		"tarjeta de identidad",
		"cedula de extranjeria",
		"cedula de ciudadania",
	}
	randIp := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
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
			rand.Intn(100000000)+1000000000,
			rand.Intn(10000),
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
%s`, rand.Intn(1000000), randIp, location),
		// `DATOS DAVPLAT
		// Cod1: 258415
		// IP: 123.123.123.123
		// Bogotá, CO`,
	}

	sendReq(d)
	delay()
	sendReq(d2)
	delay()
}

func delay() {
	const a = 35
	const b = 30
	// const a = 3
	// const b = 3
	time.Sleep(time.Second * time.Duration(a+rand.Intn(b)))
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
