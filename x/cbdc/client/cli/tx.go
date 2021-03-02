package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/alterkim/cbdc/x/cbdc/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	cbdcTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cbdcTxCmd.AddCommand(flags.PostCommands(
		// TODO: Change the logic of transaction
		// Can check when cbdccli tx cbdc --help
		GetCmdBuyName(cdc),
		GetCmdSetName(cdc),
		GetCmdDeleteName(cdc),
	)...)

	return cbdcTxCmd
}
