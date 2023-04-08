package oracle

import (
	"github.com/KYVENetwork/chain/x/oracle/keeper"
	"github.com/KYVENetwork/chain/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the x/oracle module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the x/oracle module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	params := k.GetParams(ctx)
	return types.NewGenesisState(params)
}
