package discord_bot

import (
	"fmt"
	"github.com/andersfylling/disgord"
	"github.com/discord_login/keys"
	"log"
	"strings"
)

const (
	HelpCmd = "help"
)
func (bot *Bot) HandleHelp(s disgord.Session, data *disgord.MessageCreate) {
	msg := data.Message
	fmt.Println("MSG: ", msg.ID)
	if !strings.HasPrefix(msg.Content, HelpCmd) {
		return
	}

	log.Printf(keys.LogCommand, HelpCmd, "received %s command", HelpCmd)

	// Answer to the command
	bot.React(msg, s, keys.ReactionDone)
	bot.Reply(msg, s, fmt.Sprintf(
		"Here are the available command:\n"+
			"- `!%s` - to get help\n" +
			"- `!%s` - to get account balance",
		HelpCmd,
		BalanceCmd,
	))
}
