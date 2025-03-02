/**
 * Requests to external/vendor APIs
 */

package common

// Discord Structs

type Discord struct {
	WebHook string
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
