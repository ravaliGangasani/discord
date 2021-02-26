package discordoauth

import (
	"github.com/discord_login/config"
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
)

func Router(mux *mux.Router, conf *oauth2.Config, cfg *config.Config) {
	mux.HandleFunc("/oauth2/login", AuthorizeLoginHandlerFn(conf, cfg)).Methods("GET")
	mux.HandleFunc("/oauth2/login/redirect", LoginRedirectHandlerFn(conf, cfg)).Methods("GET")
}
