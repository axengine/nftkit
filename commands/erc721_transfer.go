package commands

import (
	"fmt"
	"github.com/axengine/ethcli"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"math/big"
)

var (
	key string
	to  string
)

var (
	tokenId int64
)

func NewERC721TransferCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "721",
		Short: "ERC721转账",
		Long:  "./ntfkit 721 --id=8256 --to=0x4f4BB24c92a8913FD0462E32D4c6599f7aF7B03d --key=...",
		RunE:  ERC721TransferFrom,
		PreRun: func(cmd *cobra.Command, args []string) {

		},
	}
	cmd.Flags().StringVar(&key, "key", "", "private key")
	cmd.Flags().StringVar(&to, "to", "", "to address")
	cmd.Flags().Int64Var(&tokenId, "id", 0, "tokenId")
	return cmd
}

func ERC721TransferFrom(cmd *cobra.Command, args []string) error {
	if key == "" || to == "" || tokenId == 0 {
		return errors.New("invalid param")
	}

	to = common.HexToAddress(to).Hex()
	zero := common.Address{}
	if to == zero.Hex() {
		return errors.New("invalid to")
	}

	priKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return err
	}
	from := crypto.PubkeyToAddress(priKey.PublicKey).Hex()

	cli, _ := ethcli.New("https://zerorpc.singularity.gold")
	owner, err := cli.ORC721OwnerOf("0x97BC55F017585d04339fc85a156E2846137fd21f", big.NewInt(tokenId), nil)
	if err != nil {
		return err
	}
	if owner != from {
		return errors.New("tokenId不是自己的")
	}
	result, err := cli.ORC721TransferFrom("0x97BC55F017585d04339fc85a156E2846137fd21f",
		key,
		from,
		to,
		big.NewInt(int64(tokenId)))

	fmt.Printf("交易hash：%s err:%v", result, err)
	return nil
}
