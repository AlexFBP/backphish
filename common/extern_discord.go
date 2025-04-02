/**
 * Requests to external/vendor APIs
 */

package common

type Discord struct {
	ReqHandler
	Webhook string
}

type HookData struct {
	User    string `json:"username"`
	Content string `json:"content"`
}

func (m *Discord) SendDiscord(data HookData) {
	m.SendJSON("https://discord.com/api/webhooks/"+m.Webhook, data, nil, nil)
	// use discord instead of discordapp:
	// https://github.com/discord/discord-api-docs/issues/1878
}
