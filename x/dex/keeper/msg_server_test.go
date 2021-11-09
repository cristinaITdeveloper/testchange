package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/cristinaITdeveloper/testchange/testutil/keeper"
	"github.com/cristinaITdeveloper/testchange/x/dex/keeper"
	"github.com/cristinaITdeveloper/testchange/x/dex/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DexKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
