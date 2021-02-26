package config

import (
	"encoding/json"
	"io/ioutil"
)

const STATE = "secret"
const SCOPEIDENTY = "identify"

type Config struct {
	OauthConfig OauthConfig  `json:"oauth_config"`
	BotConfig   *BotConfig   `json:"bot_config"`
	ChainConfig *ChainConfig `json:"chain_config"`
	State       string       `json:"state"`
}

type OauthConfig struct {
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RedirectURL  string   `json:"redirect_url"`
	EndPoint     EndPoint `json:"end_point"`
	Scope        string   `json:"scope"`
}

type EndPoint struct {
	AuthURL  string `json:"auth_url"`
	TokenURL string `json:"token_url"`
}

type BotConfig struct {
	Token  string `json:"token"`
	Prifix string `json:"prifix"`
}

type ChainConfig struct {
	NodeURI      string         `json:"node_uri"`
	Bech32Prefix string         `json:"bech_32_prefix"`
	ChainID      string         `json:"chain_id"`
	Fees         string         `json:"fees"`
	Account      *AccountConfig `json:"account"`
}

type AccountConfig struct {
	Mnemonic string `json:"mnemonic"`
	HDPath   string `json:"hd_path"`
}

func Parse(filePath string) (*Config, error) {
	bz, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(bz, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
