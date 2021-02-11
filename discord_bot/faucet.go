package discord_bot

import (
	"fmt"
	"github.com/andersfylling/disgord"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/discord_login/keys"
	"strings"
)

const (
	FaucetCmd = "faucet"
)

func (bot *Bot) HandleFaucet(s disgord.Session, data *disgord.MessageCreate) {
	msg := data.Message
	if !strings.HasPrefix(msg.Content, FaucetCmd) {
		return
	}

	path := strings.Split(msg.Content, " ")
	if len(path) < 2 {
		bot.Reply(msg, s, "Missing recipient")
		bot.React(msg, s, keys.ReactionWarning)
	}

	fmt.Println("paths: ", path[1])

	addr, err := sdk.AccAddressFromBech32(path[1])

	if err != nil {
		bot.React(msg, s, keys.ReactionWarning)
		bot.Reply(msg, s, "invalid address provided")
		return
	}

	fmt.Println("Address: ", addr.String())
	fmt.Println("FromAddress", bot.cosmosClient.AccAddress())
	txMsg := &types.MsgSend{
		FromAddress: bot.cosmosClient.AccAddress(),
		ToAddress: addr.String(),
		Amount: sdk.NewCoins(sdk.NewCoin("cent", sdk.NewInt(1000000))),
	}

	fmt.Println("Sending coins")
	res, err := bot.cosmosClient.BroadcastTx(txMsg)
	if err != nil {
		fmt.Println("error sending coins: ", err.Error())
		bot.Reply(msg, s, "error while sending coins")
		bot.React(msg, s, keys.ReactionWarning)
	} else {
		bot.React(msg, s, keys.ReactionDone)
		bot.Reply(msg, s, fmt.Sprintf(
			"Your tokens have been sent successfully. You can see it by running `desmos q tx %s`."+
				"If your balance does not update in the next seconds, make sure your node is synced.", res.TxHash,
		))
	}
}
