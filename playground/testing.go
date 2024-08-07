package playground

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/retgits/creditcard"

	"github.com/AlexFBP/backphish/common"
)

func init() {
	gofakeit.Seed(time.Now().UnixNano()) // or gofakeit.Seed(0)
}

func Cmd(args ...string) (err error) {

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

	return nil
}
