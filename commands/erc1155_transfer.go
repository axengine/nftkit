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
	amount int64
)

func NewERC1155TransferCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "1155",
		Short: "ERC1155转账",
		Long:  "./ntfkit 1155 --amount=1 --to=0x4f4BB24c92a8913FD0462E32D4c6599f7aF7B03d --key=...",
		RunE:  ERC1155TransferFrom,
		PreRun: func(cmd *cobra.Command, args []string) {

		},
	}
	cmd.Flags().StringVar(&key, "key", "", "private key")
	cmd.Flags().StringVar(&to, "to", "", "to address")
	cmd.Flags().Int64Var(&amount, "amount", 0, "amount")
	return cmd
}

func ERC1155TransferFrom(cmd *cobra.Command, args []string) error {
	if key == "" || to == "" || amount == 0 {
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
	bal, err := cli.ORC1155BalanceOf("0x19d61BBbF4682196EF2c448d70C9f2424D421C82", from, big.NewInt(1), nil)
	if err != nil {
		return err
	}
	if big.NewInt(amount).Cmp(bal) > 0 {
		return errors.New("余额不足")
	}
	result, err := cli.ORC1155SafeTransferFrom(key,
		"0x19d61BBbF4682196EF2c448d70C9f2424D421C82",
		from,
		to,
		big.NewInt(1),
		big.NewInt(amount), nil)

	fmt.Printf("交易hash：%s err:%v", result, err)
	return nil
}
