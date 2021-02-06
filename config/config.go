package config

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
)

const AuthRedirectUrl = "https://discord.com/api/oauth2/authorize?client_id=806764091020279848&redirect_uri=http%3A%2F%2Flocalhost%3A8000%2Foauth2%2Flogin%2Fredirect&response_type=code&scope=identify"
const STATE = "secret"
const SCOPEIDENTY = "identify"

type Config struct {
	BotConfig   *BotConfig   `json:"bot_config"`
	//ChainConfig ChainConfig `json:"chain_config"`
}

type BotConfig struct {
	Token  string `json:"token"`
	Prifix string `json:"prifix"`
}

type ChainConfig struct {
	NodeURI      string `json:"node_uri"`
	Bech32Prefix string `json:"bech_32_prefix"`
	ChainID      string `json:"chain_id"`
}

func Parse(filePath string) (*Config, error){
	bz, err := ioutil.ReadFile(filePath)
	if err !=nil {
		return nil, err
	}

	var config Config
	err = toml.Unmarshal(bz, &config)
	if err != nil{
		return nil, err
	}
	return &config, nil
}