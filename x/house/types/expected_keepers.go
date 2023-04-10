package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SRKeeper defines the expected strategicreserve keeper.
type SRKeeper interface {
	InitiateOrderBookParticipation(ctx sdk.Context, addr sdk.AccAddress,
		bookUID string, liquidity, fee sdk.Int) (uint64, error)
	WithdrawOrderBookParticipation(ctx sdk.Context, depAddr,
		bookUID string, bpNumber uint64, mode WithdrawalMode, amount sdk.Int) (sdk.Int, error)
}
