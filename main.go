package main

import (
	"ntfkit/commands"
)

func main() {
	rootCmd := commands.RootCmd
	rootCmd.AddCommand(commands.VersionCmd,
		commands.NewERC721TransferCmd(),
		commands.NewERC1155TransferCmd(),
	)
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
