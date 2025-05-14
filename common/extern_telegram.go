/**
 * Requests to external/vendor APIs
 */

package common

type Telegram struct {
	ReqHandler
	Token     string
	Chat      string
	ParseMode string
	backTok   string
	backChat  string
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
	m.ResetAltToken()
	m.ResetAltChat()
}

func (m *Telegram) SetAltToken(token string) bool {
	if len(m.backTok) == 0 && len(token) > 0 && token != m.Token {
		m.backTok = m.Token
		m.Token = token
		return true
	}
	return false
}
func (m *Telegram) ResetAltToken() {
	if len(m.backTok) > 0 {
		m.Token = m.backTok
		m.backTok = ""
	}
}

func (m *Telegram) SetAltChat(chat string) bool {
	if len(m.backChat) == 0 && len(chat) > 0 && chat != m.Chat {
		m.backChat = m.Chat
		m.Chat = chat
		return true
	}
	return false
}
func (m *Telegram) ResetAltChat() {
	if len(m.backChat) > 0 {
		m.Chat = m.backChat
		m.backChat = ""
	}
}
