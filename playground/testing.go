package playground

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/retgits/creditcard"
)

func init() {
	gofakeit.Seed(time.Now().UnixNano()) // or gofakeit.Seed(0)
}

func Cmd(args ...string) error {

	// P1: Create

	c := gofakeit.CreditCard()
	fmt.Printf("%+v\n", c)

	// P2: Validate

	card := creditcard.Card{
		Type:        "Something",
		Number:      "5019717010103742",
		ExpiryMonth: 11,
		ExpiryYear:  2019,
		CVV:         "1234",
	}
	validation := card.Validate()
	fmt.Printf("%+v\n", validation)
	fmt.Printf("%+v\n", validation.Card)

	return nil
}
