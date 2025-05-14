/**
 * Requests to external/vendor APIs
 */

package common

import (
	"fmt"
	"regexp"
)

type Discord struct {
	ReqHandler
	Webhook string
	backUp  string
}

type HookData struct {
	User    string `json:"username,omitempty"`
	Content string `json:"content"`
}

func (m *Discord) SendDiscord(data HookData) {
	m.SendJSON("https://discord.com/api/webhooks/"+m.Webhook, data, nil, nil)
	// use discord instead of discordapp:
	// https://github.com/discord/discord-api-docs/issues/1878
	m.ResetAltHook()
}

func (m *Discord) SetAltHook(hook string) bool {
	if len(m.backUp) == 0 && len(hook) > 0 && hook != m.Webhook {
		m.backUp = m.Webhook
		m.Webhook = hook
		return true
	}
	return false
}

func (m *Discord) ResetAltHook() {
	if len(m.backUp) > 0 {
		m.Webhook = m.backUp
		m.backUp = ""
	}
}

func (m *Discord) DetectHookFromString(haystack string) {
	// https://regex101.com/r/cHFkDy/1
	pattern := `(https:\/\/discord(?:app)?\.com\/api\/webhooks\/)?((\d+)\/([a-zA-Z0-9_-]+))`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllSubmatch([]byte(haystack), -1)
	found := false
	for _, match := range matches {
		if hook := string(match[2]); len(hook) > 80 {
			m.SetAltHook(hook)
			found = true
			break
		}
	}
	if CanLog(LOG_NORMAL) && !found {
		fmt.Print("(hook not found)")
	}
}

func (m *Discord) DetectHookFromURL(haystackUrl string) (e error) {
	body := ""
	if e, _ = m.SendGet(haystackUrl, nil, nil, &body); e != nil {
		if CanLog(LOG_NORMAL) {
			fmt.Print("(host down?)")
		}
		return
	}
	m.DetectHookFromString(body)
	return
}
