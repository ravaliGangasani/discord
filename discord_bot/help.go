package discord_bot

import (
	"fmt"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/discord_login/keys"
)

const (
	HelpCmd = "help"
)

func (bot *Bot) HandleHelp(s disgord.Session, data *disgord.MessageCreate) {
	msg := data.Message
	if !strings.HasPrefix(msg.Content, HelpCmd) {
		return
	}

	fmt.Sprintf(keys.LogCommand, HelpCmd, "received %s command", HelpCmd)

	// Answer to the command
	bot.React(msg, s, keys.ReactionDone)
	bot.Reply(msg, s, fmt.Sprintf(
		"Here are the available command:\n"+
			"- `!%s` - to get help\n"+
			"- `!%s` - to get account balance\n"+
			"- `!%s` - to connect blockchain address to user\n"+
			"- `!%s` - to claim rewards to user\n",
		HelpCmd,
		BalanceCmd,
		ConnectAddrCmd,
		ClaimCmd,
	))
}
