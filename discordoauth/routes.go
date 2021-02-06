package discordoauth

import (
	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
)

func Router(mux *mux.Router, conf *oauth2.Config) {
	mux.HandleFunc("/oauth2/login", AuthorizeLoginHandlerFn(conf)).Methods("GET")
	mux.HandleFunc("/oauth2/login/redirect", LoginRedirectHandlerFn(conf)).Methods("GET")
}
