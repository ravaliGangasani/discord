package discord_bot

import (
	"context"
	"fmt"
	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
	"github.com/discord_login/config"
	"github.com/discord_login/cosmos"
	"github.com/ethereum/go-ethereum/log"
)

type Bot struct {
	cfg     *config.BotConfig
	discord *disgord.Client
	cosmosClient *cosmos.Client
}

func Create(conf *config.BotConfig, client *cosmos.Client) (*Bot, error) {
	if conf.Prifix == "" {
		conf.Prifix = "!"
	}

	discordClient := disgord.New(disgord.Config{
		ProjectName: "discord_bot",
		BotToken:    conf.Token,
		RejectEvents: []string{
			disgord.EvtTypingStart,

			// These require special privilege
			// https://discord.com/developers/docs/topics/gateway#privileged-intents
			disgord.EvtPresenceUpdate,
			disgord.EvtGuildMemberAdd,
			disgord.EvtGuildMemberUpdate,
			disgord.EvtGuildMemberRemove,
		},
		DMIntents: disgord.IntentDirectMessages |
			disgord.IntentDirectMessageReactions |
			disgord.IntentDirectMessageTyping,
		Presence: &disgord.UpdateStatusPayload{
			Game: &disgord.Activity{
				Name: "Welcome users!",
			},
		},
	})

	return &Bot{
		cfg:     conf,
		discord: discordClient,
		cosmosClient: client,
	}, nil
}

func (bot *Bot) Start() {

	defer bot.discord.Gateway().StayConnectedUntilInterrupted()

	fmt.Println("starting  bot")
	filter, _ := std.NewMsgFilter(context.Background(), bot.discord)
	filter.SetPrefix(bot.cfg.Prifix)
	handler := bot.discord.Gateway().WithMiddleware(
		filter.HasPrefix,
		filter.NotByBot,
		filter.StripPrefix)
	handler.MessageCreate(
		bot.HandleHelp,
		bot.HandleQueryBalance,
		bot.HandleFaucet,
		)

	fmt.Println("listening for messages...")
}

// Reply sends a Discord message as a reply to the given msg
func (bot *Bot) Reply(msg *disgord.Message, s disgord.Session, message string) {

	//_, _, err := msg.Author.SendMsg(context.Background(), s, &disgord.Message{
	//	Type:    disgord.MessageTypeDefault,
	//	Content: message,
	//})

	_, err := s.Channel(msg.ChannelID).WithContext(context.Background()).CreateMessage(&disgord.CreateMessageParams{
		Content: message,
	})
	//_, err = msg1.Send(context.Background(), s)
	if err != nil {
		log.Error("failed to reply messages: ", err.Error())
	}
}

// React allows to react with the provided emoji to the given message
func (bot *Bot) React(msg *disgord.Message, s disgord.Session, emoji interface{}, flags ...disgord.Flag) {
	err := msg.React(context.Background(), s, emoji, flags...)
	if err != nil {
		log.Error("failed to reply messages: ", err.Error())
	}
}
