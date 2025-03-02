/**
 * Requests to external/vendor APIs
 */

package common

// Discord Structs

type Discord struct {
	WebHook string
}

type hookData struct {
	User    string `json:"username"`
	Content string `json:"content"`
}

// Telegram Structs

type Telegram struct {
	Token string
}

type TgMsg struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
}

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
