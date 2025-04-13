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
	v := target.GetMirrorParams(name)
	d.Webhook, d.Token, d.Chat = v[0], v[1], v[2]
	d.ParseOptions(v[3])
	return
}
