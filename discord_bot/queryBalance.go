package discord_bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/andersfylling/disgord"
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
	if path[1] == "" {
		bot.Reply(msg, s, "No address provided")
		return
	}

	url := fmt.Sprintf("http://20.102.97.37:1317/bank/balances/%s", path[1])
	res, err := http.Get(url)
	if err != nil {
		bot.Reply(msg, s, "error while fetching balance")
		return
	}

	bz, _ := ioutil.ReadAll(res.Body)

	var response ResponseBody
	_ = json.Unmarshal(bz, &response)
	bot.Reply(msg, s, string(response.Result))
}

type ResponseBody struct {
	Height int64           `json:"height"`
	Result json.RawMessage `json:"result"`
}
