package discord_bot

import (
	"errors"
	"fmt"
	"strings"

	"github.com/andersfylling/disgord"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/discord_login/keys"
	"github.com/discord_login/models"
	"gopkg.in/mgo.v2/bson"
)

const (
	ConnectAddrCmd = "connectaddr"
)

func (bot *Bot) HandleConnectAddres(s disgord.Session, data *disgord.MessageCreate) {
	msg := data.Message
	if !strings.HasPrefix(msg.Content, ConnectAddrCmd) {
		return
	}
	path := strings.Split(msg.Content, " ")
	if len(path) < 2 {
		bot.Reply(msg, s, "Missing recipient")
		bot.React(msg, s, keys.ReactionWarning)
	}

	addr, err := sdk.AccAddressFromBech32(path[1])
	if err != nil {
		bot.React(msg, s, keys.ReactionWarning)
		bot.Reply(msg, s, "invalid address provided")
		return
	}

	user := models.User{}
	query := bson.M{"id": data.Message.Author.ID.String()}
	err = user.FindOne(query)
	var ErrNotFound = errors.New("not found")

	if err != nil {
		if err.Error() == ErrNotFound.Error() {
			user = models.User{ID: data.Message.Author.ID.String(),
				Name:    data.Message.Author.Username,
				Address: addr.String(), Claimed: false}

			err = user.Save()
			if err != nil {
				fmt.Println("error: ", err.Error())
				bot.Reply(msg, s, "error while saving user")
				bot.React(msg, s, keys.ReactionWarning)
			} else {
				bot.React(msg, s, keys.ReactionDone)
				bot.Reply(msg, s, "Address successfully linked")
			}
		}
	} else {
		if user.Address == addr.String() {
			bot.React(msg, s, keys.ReactionDone)
			bot.Reply(msg, s, "Address already linked")
		} else {
			bot.React(msg, s, keys.ReactionDone)
			bot.Reply(msg, s, fmt.Sprintf("Your account alredy linked with %s address", user.Address))
		}

	}

}
