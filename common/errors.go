package common

import "fmt"

type CustErr struct {
	Explain string
	Detail  error
}

func (c *CustErr) Error() string {
	return fmt.Sprintf("%s :: %s", c.Explain, c.Detail.Error())
}
