module github.com/discord_login

go 1.15

require (
	github.com/AutonomyNetwork/autonomy-chain v1.0.2-0.20220311080500-9f07dc0f66bf
	github.com/andersfylling/disgord v0.26.2
	github.com/cosmos/cosmos-sdk v0.44.3
	github.com/ethereum/go-ethereum v1.9.25
	github.com/gorilla/mux v1.8.0
	github.com/rs/zerolog v1.23.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/tendermint/tendermint v0.34.14
	golang.org/x/oauth2 v0.0.0-20210514164344-f6687ab2804c
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
