package discordoauth

import (
	"fmt"
	config2 "github.com/discord_login/config"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

func AuthorizeLoginHandlerFn(conf *oauth2.Config, cfg *config2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, conf.AuthCodeURL(cfg.State), http.StatusSeeOther)
	}
}

func LoginRedirectHandlerFn(conf *oauth2.Config, cfg *config2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		state := r.URL.Query().Get("state")

		if state != cfg.State {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("State does not match"))
			return
		}

		token, err := conf.Exchange(oauth2.NoContext, code)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}

		res, err := conf.Client(oauth2.NoContext, token).Get("https://discordapp.com/api/v6/users/@me")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error while accessing user details"))
			return
		}

		defer res.Body.Close()
		bz, err := ioutil.ReadAll(res.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error while reading data from response"))
			return
		}
		w.Write(bz)
		return
	}

}
