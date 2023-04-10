package keeper

import (
	"github.com/playerfury/fanx/x/orderbook/types"
)

var _ types.QueryServer = Keeper{}
