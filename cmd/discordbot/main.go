package main

import (
	"github.com/discord_login/cmd/discordbot/cmd"
	"os"
)

func main() {
	executor := cmd.RootCmd()
	if err := executor.Execute(); err != nil {
		os.Exit(1)
	}
}