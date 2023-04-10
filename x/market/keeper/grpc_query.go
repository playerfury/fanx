package keeper

import (
	"github.com/rotofury/xfury/x/market/types"
)

var _ types.QueryServer = Keeper{}
