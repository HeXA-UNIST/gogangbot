package intro

import (
	"github.com/fabioxgn/go-bot"
)

func intro(command *bot.Cmd) (msg string, err error) {
	return "Gogangbot", nil
}

func init() {
	bot.RegisterCommand(
		"intro",
		"",
		"",
		intro)
}
