package discord_bot

import (
	"encoding/json"
	"fmt"
	"github.com/andersfylling/disgord"
	"github.com/cosmos/cosmos-sdk/types"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	BalanceCmd = "balance"
)

func (bot *Bot) HandleQueryBalance(s disgord.Session, data *disgord.MessageCreate) {
	msg := data.Message

	if !strings.HasPrefix(msg.Content, BalanceCmd) {
		return
	}

	path := strings.Split(msg.Content, " ")
	addr, err := types.AccAddressFromBech32(path[1])
	if err != nil {
		fmt.Println("Invalid address provided", err.Error())
		return
	}

	url:= fmt.Sprintf("http://localhost:1317/bank/balances/%s", addr)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("error while querying balance: ", err.Error())
		return
	}

	bz, _ := ioutil.ReadAll(res.Body)

	var response ResponseBody
	_ = json.Unmarshal(bz, &response)
	bot.Reply(msg, s, string(response.Result))
}

type ResponseBody struct {
	Height int64 `json:"height"`
	Result json.RawMessage `json:"result"`
}

