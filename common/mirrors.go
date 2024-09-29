package common

import (
	"fmt"
	"log"

	"github.com/turret-io/go-menu/menu"
)

type Target struct {
	Prefix, Description string
	Mirrors             []string
	Handler             func(mirror string)
}

func (t *Target) GetAllCmds() (opts []menu.CommandOption) {
	for k, v := range t.Mirrors {
		opts = append(opts, menu.CommandOption{
			Command:     fmt.Sprintf("%s-%d", t.Prefix, k+1),
			Description: fmt.Sprintf("%s, mirror %d (%s)", t.Description, k+1, v),
			Function:    t.getCmd(k),
		})
	}
	return
}

func (t *Target) getCmd(k int) (f func(...string) error) {
	return func(args ...string) error {
		return AttackRunner(t.mirrorAttempt(k))
	}
}

func (t *Target) mirrorAttempt(n int) AttemptHander {
	L := len(t.Mirrors)
	if 0 <= n && n < L {
		return func() {
			t.Handler(t.Mirrors[n])
		}
	}
	return func() {
		log.Fatalf("\n[FATAL] Wrong mirror number, max:%d - got:%d\n", L, n)
	}
}
