package playground

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/retgits/creditcard"

	"github.com/AlexFBP/backphish/common"
	"github.com/AlexFBP/backphish/targets/unclassified"
)

func init() {
	gofakeit.Seed(time.Now().UnixNano()) // or gofakeit.Seed(0)
}

func pruebaDefer() (equis common.CustErr) {
	defer func() {
		equis.Explain = "entré despues de la muerte"
	}()
	equis.Explain = "entré una vez"
	return
}

func Cmd(args ...string) (err error) {
	fmt.Println("Input args:")
	for k, v := range args {
		fmt.Printf("[%d]:%s\n", k, v)
	}

	// P1: Create
	p := gofakeit.Person()
	user := common.RandUserName(p)
	fmt.Println("user: ", user)

	c := p.CreditCard
	fmt.Printf("%+v\n", c)

	// P2: Validate

	exp := strings.Split(c.Exp, "/")
	var month, year int
	month, err = strconv.Atoi(exp[0])
	if err != nil {
		return
	}
	year, err = strconv.Atoi(exp[1])
	if err != nil {
		return
	}
	year += 2000

	card := creditcard.Card{
		Type:        c.Type,
		Number:      strconv.Itoa(c.Number),
		ExpiryMonth: month,
		ExpiryYear:  year,
		CVV:         c.Cvv,
	}
	validation := card.Validate()
	for _, v := range validation.Errors {
		fmt.Println(v)
	}
	validation.Errors = []string{}
	fmt.Printf("%+v\n", validation)
	fmt.Printf("%+v\n", validation.Card)

	cel := common.GeneraCelColombia()
	fmt.Printf("cel: %s\n", cel)

	qty := 5
	common.TimedCall(time.Second, 60*time.Second, func() (iterateAgain bool) {
		fmt.Println(qty)
		qty--
		if qty == 3 {
			time.Sleep(1500 * time.Millisecond)
		}
		return qty > 0
	})

	valPrueba := pruebaDefer().Explain
	fmt.Println("pruebaDefer().Explain:", valPrueba)

	// TODO: get bool from cli flag
	// For now, testing only alive hosts
	unclassified.CheckHosts(false)

	return nil
}
