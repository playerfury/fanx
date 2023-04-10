package keeper

import (
	"github.com/rotofury/xfury/x/dvm/types"
)

var _ types.QueryServer = Keeper{}
