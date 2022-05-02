package cosmos

import (
	"github.com/AutonomyNetwork/autonomy-chain/app"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

	"github.com/discord_login/config"
)

// Client represents a Cosmos client that should be used to create and send transactions to the chain
type Client struct {
	cliCtx client.Context

	privKey cryptotypes.PrivKey

	fees string
}

// NewClient allows to build a new Client instance
func NewClient(chainCfg *config.ChainConfig) (*Client, error) {
	// Get the private keys
	algo := hd.Secp256k1
	derivedPriv, err := algo.Derive()(chainCfg.Account.Mnemonic, "", chainCfg.Account.HDPath)
	if err != nil {
		return nil, err
	}
	privKey := algo.Generate()(derivedPriv)

	// Build the RPC client
	rpcClient, err := rpchttp.New(chainCfg.NodeURI, "/websocket")
	if err != nil {
		return nil, err
	}

	// Build the config
	prefix := chainCfg.Bech32Prefix
	sdkCfg := sdk.GetConfig()
	sdkCfg.SetBech32PrefixForAccount(prefix, prefix+sdk.PrefixPublic)
	sdkCfg.SetBech32PrefixForValidator(
		prefix+sdk.PrefixValidator+sdk.PrefixOperator,
		prefix+sdk.PrefixValidator+sdk.PrefixOperator+sdk.PrefixPublic,
	)
	sdkCfg.SetBech32PrefixForConsensusNode(
		prefix+sdk.PrefixValidator+sdk.PrefixConsensus,
		prefix+sdk.PrefixValidator+sdk.PrefixConsensus+sdk.PrefixPublic,
	)
	sdkCfg.Seal()

	// Build the context
	encodingConfig := app.MakeEncodingConfig()
	cliCtx := client.Context{}.
		WithCodec(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithAccountRetriever(types.AccountRetriever{}).
		WithBroadcastMode(flags.BroadcastSync).
		WithClient(rpcClient).
		WithChainID(chainCfg.ChainID).
		WithFromAddress(sdk.AccAddress(privKey.PubKey().Address()))

	return &Client{
		cliCtx:  cliCtx,
		privKey: privKey,
		fees:    chainCfg.Fees,
	}, nil
}

// AccAddress returns the address of the account that is going to be used to sign the transactions
func (client *Client) AccAddress() string {
	return sdk.AccAddress(client.privKey.PubKey().Address()).String()
}
