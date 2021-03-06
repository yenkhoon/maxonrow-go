package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
	maintenanceCmd "github.com/maxonrow/maxonrow-go/x/maintenance/client/cli"
)

type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:   "maintenance",
		Short: "Querying commands for the maintenance module",
	}

	queryCmd.AddCommand(client.GetCommands(
		maintenanceCmd.GetCmdGetProposal(mc.storeKey, mc.cdc),
	)...)

	return queryCmd
}

func (mc ModuleClient) GetTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:   "maintenance",
		Short: "Maintenance transactions subcommands",
	}

	txCmd.AddCommand(client.PostCommands(
		maintenanceCmd.GetCmdSubmitProposal(mc.cdc),
		maintenanceCmd.GetCmdCastAction(mc.cdc),
	)...)

	return txCmd
}
