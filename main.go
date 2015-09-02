package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/HeXA-UNIST/gogangbot/commands/intro"
	_ "github.com/HeXA-UNIST/gogangbot/commands/memo"
	"github.com/fabioxgn/go-bot"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		token = "xoxb-10051409156-5SHRSBxDUNfOC5fuSUo3XnBd"
	}

	bot.RunSlack(token)
}
