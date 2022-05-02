package discord_bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/andersfylling/disgord"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/discord_login/keys"
)

const (
	SendCmd = "send"
)

func (bot *Bot) HandleFaucet(s disgord.Session, data *disgord.MessageCreate) {
	msg := data.Message
	if !strings.HasPrefix(msg.Content, SendCmd) {
		return
	}
	fmt.Println(data.Message.Author.Username, data.Message.Author.ID)
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

	values := map[string]string{"address": string(addr), "denom": "uaut"}
	json_data, err := json.Marshal(values)

	resp, err := http.Post("https://faucet.wouo.autonomy.network/credit", "application/json",
		bytes.NewBuffer(json_data))

	// txMsg := &types.MsgSend{
	// 	FromAddress: bot.cosmosClient.AccAddress(),
	// 	ToAddress:   addr.String(),
	// 	Amount:      sdk.NewCoins(sdk.NewCoin("cent", sdk.NewInt(1000000))),
	// }

	// res, err := bot.cosmosClient.BroadcastTx(txMsg)
	fmt.Println("response:", resp)
	if err != nil {
		fmt.Println("error: ", err.Error())
		bot.Reply(msg, s, "error while sending coins")
		bot.React(msg, s, keys.ReactionWarning)
	} else {
		bot.React(msg, s, keys.ReactionDone)
		bot.Reply(msg, s, fmt.Sprintf("%v", resp))
	}
}
