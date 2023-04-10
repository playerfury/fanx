package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/playerfury/fanx/app/params"
	"github.com/playerfury/fanx/x/strategicreserve/types"
)

// DepositDenom returns deposit coin denomination
func (k Keeper) DepositDenom(ctx sdk.Context) (res string) {
	return params.DefaultBondDenom
}

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramstore.GetParamSet(ctx, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
