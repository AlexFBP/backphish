package common

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/turret-io/go-menu/menu"
)

type MirrorOptions struct {
	options []string
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

type TargetInterface interface {
	GetMirrors() []string
	GetProps() TargetBase
	Handler(mirror string)
	MirrorStatus(mirror string) int
	EstimateDuration() time.Duration
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

type TargetBase struct {
	Prefix, Description string
}

func (t *TargetBase) GetProps() TargetBase {
	return *t
}

type TargetSimple struct {
	TargetBase
	mirrors *[]string
}

func (t *TargetSimple) SetMirrors(m *[]string) {
	t.mirrors = m
}

func (t *TargetSimple) GetMirrors() []string {
	return *t.mirrors
}

type TargetWithParams struct {
	TargetBase

	// mirrors is a slice of slices, where each inner slice contains the mirror and its parameters.
	// The quantity of parameters is not fixed.
	mirrors *[][]string
}

func (t *TargetWithParams) SetMirrors(m *[][]string) {
	t.mirrors = m
}

func (t *TargetWithParams) GetMirrors() (names []string) {
	if t.mirrors == nil {
		return
	}
	for _, v := range *t.mirrors {
		names = append(names, v[0])
	}
	return names
}

func (t *TargetWithParams) GetMirrorParams(mirror string) []string {
	if t.mirrors == nil {
		return []string{}
	}
	for _, v := range *t.mirrors {
		if mirror == v[0] {
			return v[1:]
		}
	}
	return []string{}
}
