package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/alterkim/nameservice/x/nameservice/keeper"
	"github.com/alterkim/nameservice/x/nameservice/types"
	// abci "github.com/tendermint/tendermint/abci/types"
)

// InitGenesis initialize default parameters (called on chain start)
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, k keeper.Keeper /* TODO: Define what keepers the module needs */, data types.GenesisState) {
	for _, record := range data.WhoisRecords {
		keeper.SetWhois(ctx, record.Value, record)
	}
}

// ExportGenesis writes the current store values (called after stopping the chain)
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) (data types.GenesisState) {
	var records []types.Whois
	iterator := k.GetNamesIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		name := string(iterator.Key())
		whois, _ := k.GetWhois(ctx, name)
		records = append(records, whois)

	}
	return types.GenesisState{WhoisRecords: records}
}
