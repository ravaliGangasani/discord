package keys

import (
	"os"
	"path"
)

var (
	HomeDir, _ = os.UserHomeDir()
	DataDir    = path.Join(HomeDir, ".discordbot")
)

const (
	AppName = "discordbot"

	ReactionWarning = "⚠️"
	ReactionDone    = "✅"
	ReactionTime    = "⌛"

	LogCommand       = "command"
	LogUser          = "user"
	LogExpirationEnd = "expiration_end"
)
