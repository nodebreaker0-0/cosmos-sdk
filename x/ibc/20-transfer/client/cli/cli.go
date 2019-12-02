package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
)

// GetQueryCmd returns the query commands for IBC fungible token transfer
func GetQueryCmd(cdc *codec.Codec, queryRoute string) *cobra.Command {
	ics20TransferQueryCmd := &cobra.Command{
		Use:   "transfer",
		Short: "IBC fungible token transfer query subcommands",
	}

	ics20TransferQueryCmd.AddCommand(client.GetCommands(
		GetCmdQueryNextSequence(cdc, queryRoute),
	)...)

	return ics20TransferQueryCmd
}

// GetTxCmd returns the transaction commands for IBC fungible token transfer
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	ics20TransferTxCmd := &cobra.Command{
		Use:   "transfer",
		Short: "IBC fungible token transfer transaction subcommands",
	}

	ics20TransferTxCmd.AddCommand(client.PostCommands(
		GetTransferTxCmd(cdc),
		GetMsgRecvPacketCmd(cdc),
	)...)

	return ics20TransferTxCmd
}
