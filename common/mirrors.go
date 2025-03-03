package common

import (
	"fmt"
	"log"
	"strings"

	"github.com/turret-io/go-menu/menu"
)

type MirrorOptions struct {
	options []string
}

type Target struct {
	Prefix, Description string
	Mirrors             []string
	Handler             func(mirror string)
}

func (t *Target) GetAllCmds() (opts []menu.CommandOption) {
	qty := len(t.Mirrors)
	digits := Digits(qty)
	for k, v := range t.Mirrors {
		if qty == 1 {
			opts = append(opts, menu.CommandOption{
				Command:     t.Prefix,
				Description: t.Description,
				Function:    t.getCmd(k),
			})
			break
		}
		opts = append(opts, menu.CommandOption{
			Command:     fmt.Sprintf("%s-%0*d", t.Prefix, digits, k+1),
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

func (m *MirrorOptions) ParseOptions(opts string) {
	if opts != "" {
		m.options = strings.Split(opts, ",")
		for k, v := range m.options {
			m.options[k] = strings.TrimSpace(v)
		}
	}
}

func (m *MirrorOptions) HasOption(opt string) bool {
	for _, v := range m.options {
		if v == strings.TrimSpace(opt) {
			return true
		}
	}
	return false
}
