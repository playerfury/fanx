package keeper

import (
	"github.com/rotofury/xfury/x/mint/types"
)

var _ types.QueryServer = Keeper{}
