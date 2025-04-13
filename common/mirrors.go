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

type TargetInterface interface {
	GetMirrors() []string
	GetProps() TargetBase
	Handler(mirror string)
}

type TargetBase struct {
	Prefix, Description string
}

func (t *TargetBase) GetProps() TargetBase {
	return *t
}

func GetAllCmds(t TargetInterface) (opts []menu.CommandOption) {
	mirrors := t.GetMirrors()
	props := t.GetProps()
	qty := len(mirrors)
	digits := Digits(qty)
	for k, v := range mirrors {
		if qty == 1 {
			opts = append(opts, menu.CommandOption{
				Command:     props.Prefix,
				Description: props.Description,
				Function:    getCmd(t, k),
			})
			break
		}
		opts = append(opts, menu.CommandOption{
			Command:     fmt.Sprintf("%s-%0*d", props.Prefix, digits, k+1),
			Description: fmt.Sprintf("%s, mirror %d (%s)", props.Description, k+1, v),
			Function:    getCmd(t, k),
		})
	}
	return
}

func getCmd(t TargetInterface, k int) (f func(...string) error) {
	return func(args ...string) error {
		return AttackRunner(mirrorAttempt(t, k))
	}
}

func mirrorAttempt(t TargetInterface, n int) AttemptHander {
	mirrors := t.GetMirrors()
	L := len(mirrors)
	if 0 <= n && n < L {
		return func() {
			t.Handler(mirrors[n])
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
