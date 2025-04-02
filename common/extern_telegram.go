/**
 * Requests to external/vendor APIs
 */

package common

type Telegram struct {
	ReqHandler
	Token     string
	Chat      string
	ParseMode string
}

type TgMsg struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
}

// Send Message to a Telegram Chat
//
// Please take care of not exceeding Telegram fair usage:
// - https://rollout.com/integration-guides/telegram-bot-api/api-essentials
// - https://core.telegram.org/bots/faq#my-bot-is-hitting-limits-how-do-i-avoid-this
func (m *Telegram) SendToTelegram(msg string) {
	tgm := TgMsg{
		ChatID: m.Chat,
		Text:   msg,
	}
	if m.ParseMode != "" {
		tgm.ParseMode = m.ParseMode
	}

	m.SendJSON("https://api.telegram.org/bot"+m.Token+"/sendMessage", tgm, nil, nil)
}
