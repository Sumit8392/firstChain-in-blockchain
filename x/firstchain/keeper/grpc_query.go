package keeper

import (
	"github.com/example/firstChain/x/firstchain/types"
)

var _ types.QueryServer = Keeper{}
