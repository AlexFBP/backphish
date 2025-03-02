package nequi9

import (
	"github.com/AlexFBP/backphish/common"
)

func (m *mirrorData) sendDiscord(h *common.ReqHandler, data hookData) {
	h.SendJSON("https://discord.com/api/webhooks/"+m.Webhook, data, nil, nil)
}

func (m *mirrorData) sendToTelegram(h *common.ReqHandler, msg string) {
	h.SendJSON("https://api.telegram.org/bot"+m.Tk+"/sendMessage", common.TgMsg{
		ChatID:    m.Chat,
		Text:      msg,
		ParseMode: "Markdown",
	}, nil, nil)
}
