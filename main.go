package main

import (
	config2 "github.com/discord_login/config"
	"github.com/discord_login/discord_bot"
	"github.com/discord_login/discordoauth"
	"github.com/discord_login/logger"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os"
)

func main() {
	var EndPoint = oauth2.Endpoint{
		AuthURL: "https://discord.com/api/oauth2/authorize",
		TokenURL: "https://discord.com/api/oauth2/token",
		AuthStyle: oauth2.AuthStyleInParams,
	}

	var conf = &oauth2.Config{
		ClientID:     "806764091020279848",
		ClientSecret: "4TPwUyAqc5XoArI4JN0HgaBY40Lg4I8h",
		RedirectURL:  "http://localhost:8000/oauth2/login/redirect",
		Endpoint:     EndPoint,
		Scopes:       []string{config2.SCOPEIDENTY},
	}


	router := mux.NewRouter()
	discordoauth.Router(router, conf)

	cfg := &config2.BotConfig{
		Token: os.Getenv("TOKEN"),
	}

	bot, err := discord_bot.Create(cfg)
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}

	bot.Start()
	err = http.ListenAndServe(":8000", logger.RequestLogger(router))
	if err != nil{
		log.Fatal(err.Error())
	}
}