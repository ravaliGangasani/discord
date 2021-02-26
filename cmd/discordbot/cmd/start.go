package cmd

import (
	"fmt"
	config2 "github.com/discord_login/config"
	"github.com/discord_login/cosmos"
	"github.com/discord_login/discord_bot"
	"github.com/discord_login/discordoauth"
	"github.com/discord_login/logger"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
	"net/http"
)

func startOauthServer(oauthConfig *oauth2.Config, cfg *config2.Config) {
	router := mux.NewRouter()
	discordoauth.Router(router, oauthConfig, cfg)
	err := http.ListenAndServe(":8000", logger.RequestLogger(router))
	if err != nil{
		panic(err)
	}
}

func StartCmd() *cobra.Command {
	return &cobra.Command{
		Use: "start [config.json]",
		Short: "Starts the bot using the provided configuration file",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cfg, err := config2.Parse(args[0])
			if err != nil {
				return err
			}

			cosmosClient, err := cosmos.NewClient(cfg.ChainConfig)
			if err != nil {
				return err
			}

			discordBot, err := discord_bot.Create(cfg.BotConfig, cosmosClient)
			if err != nil{
				return err
			}

			fmt.Println("redirect url: ", cfg.OauthConfig.RedirectURL)
			oauthConfig := &oauth2.Config{
				ClientID:    cfg.OauthConfig.ClientID,
				ClientSecret: cfg.OauthConfig.ClientSecret,
				Endpoint:     oauth2.Endpoint{
					AuthURL: cfg.OauthConfig.EndPoint.AuthURL,
					TokenURL: cfg.OauthConfig.EndPoint.TokenURL,
					AuthStyle: oauth2.AuthStyleInParams,
				},
				RedirectURL:  cfg.OauthConfig.RedirectURL,
				Scopes:       []string{cfg.OauthConfig.Scope, "identify"},
			}
			go startOauthServer(oauthConfig, cfg)
			discordBot.Start()
			return nil
		},
	}
}
