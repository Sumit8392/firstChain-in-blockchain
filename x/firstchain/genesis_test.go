package firstchain_test

import (
	"testing"

	keepertest "github.com/example/firstChain/testutil/keeper"
	"github.com/example/firstChain/x/firstchain"
	"github.com/example/firstChain/x/firstchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FirstchainKeeper(t)
	firstchain.InitGenesis(ctx, *k, genesisState)
	got := firstchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
