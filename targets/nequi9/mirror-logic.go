package nequi9

import (
	"github.com/AlexFBP/backphish/common"
)

type mirrorData struct {
	common.Discord
	common.Telegram
	common.MirrorOptions
}

func mirrData(name string) (d mirrorData) {
	for _, v := range mirrors {
		if name == v[0] {
			d.Webhook, d.Token, d.Chat = v[1], v[2], v[3]
			d.ParseOptions(v[4])
			break
		}
	}
	return
}

func getMirrorNames() (names []string) {
	for _, v := range mirrors {
		names = append(names, v[0])
	}
	return
}
