/**
 * Requests to external/vendor APIs
 */

package common

// Discord Structs

type Discord struct {
	Webhook string
}

type HookData struct {
	User    string `json:"username"`
	Content string `json:"content"`
}

// Telegram Structs

type Telegram struct {
	Token     string
	Chat      string
	ParseMode string
}

type TgMsg struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
}

func (m *Discord) SendDiscord(h *ReqHandler, data HookData) {
	h.SendJSON("https://discord.com/api/webhooks/"+m.Webhook, data, nil, nil)
}

// Send Message to a Telegram Chat
//
// Please take care of not exceeding Telegram fair usage:
// - https://rollout.com/integration-guides/telegram-bot-api/api-essentials
// - https://core.telegram.org/bots/faq#my-bot-is-hitting-limits-how-do-i-avoid-this
func (m *Telegram) SendToTelegram(h *ReqHandler, msg string) {
	tgm := TgMsg{
		ChatID: m.Chat,
		Text:   msg,
	}
	if m.ParseMode != "" {
		tgm.ParseMode = m.ParseMode
	}

	h.SendJSON("https://api.telegram.org/bot"+m.Token+"/sendMessage", tgm, nil, nil)
}
