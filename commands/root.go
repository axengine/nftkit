package commands

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "nftkit",
	Short: "nftkit : kit of ERC721 && ERC1155 transfer",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		if cmd.Name() == VersionCmd.Name() {
			return nil
		}

		return nil
	},
}
