package discord_bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/andersfylling/disgord"
	"github.com/discord_login/keys"
	"github.com/discord_login/models"
	"gopkg.in/mgo.v2/bson"
)

const (
	ClaimCmd = "claim"
)

func (bot *Bot) HandleClaim(s disgord.Session, data *disgord.MessageCreate) {
	msg := data.Message
	if !strings.HasPrefix(msg.Content, ClaimCmd) {
		return
	}
	path := strings.Split(msg.Content, " ")
	if len(path) < 2 {
		bot.Reply(msg, s, "Missing recipient")
		bot.React(msg, s, keys.ReactionWarning)
	}

	user := models.User{}
	query := bson.M{"id": data.Message.Author.ID.String()}
	err := user.FindOne(query)
	if err != nil {
		bot.React(msg, s, keys.ReactionWarning)
		bot.Reply(msg, s, "Address not connected")
		return
	}

	if user.Claimed {
		bot.React(msg, s, keys.ReactionWarning)
		bot.Reply(msg, s, "Address alredy claimed")
		return
	}

	values := map[string]string{"address": path[1], "denom": "uaut"}
	json_data, err := json.Marshal(values)
	if err != nil {
		bot.React(msg, s, keys.ReactionWarning)
		bot.Reply(msg, s, "error while mashaling data")
		return
	}

	resp, err := http.Post("https://faucet.wouo.autonomy.network/credit", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		bot.React(msg, s, keys.ReactionWarning)
		bot.Reply(msg, s, "error while sending post request")
		return
	}

	fmt.Println(resp, err)

	update := bson.M{
		"$set": bson.M{"claimed": true},
	}

	err = user.FindOneAndUpdate(query, update, true, false, false)
	if err != nil {
		fmt.Println("error: ", err.Error())
		bot.Reply(msg, s, "error while updating user")
		bot.React(msg, s, keys.ReactionWarning)
	} else {
		bot.React(msg, s, keys.ReactionDone)
		bot.Reply(msg, s, "Address successfully claimed")
	}
}
