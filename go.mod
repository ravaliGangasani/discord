module github.com/discord_login

go 1.15

require (
	github.com/andersfylling/disgord v0.26.2
	github.com/cosmos/cosmos-sdk v0.41.0
	github.com/ethereum/go-ethereum v1.9.25
	github.com/gorilla/mux v1.8.0
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/ravaliGangasani/autonomy v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.20.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/pflag v1.0.5
	github.com/tendermint/tendermint v0.34.3
	golang.org/x/oauth2 v0.0.0-20210201163806-010130855d6c
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

replace github.com/ravaliGangasani/autonomy => ../../github.com/ravaliGangasani/autonomy
