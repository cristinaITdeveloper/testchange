package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/cristinaITdeveloper/testchange/testutil/keeper"
	"github.com/cristinaITdeveloper/testchange/x/dex/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestSellOrderBookQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSellOrderBook(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetSellOrderBookRequest
		response *types.QueryGetSellOrderBookResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetSellOrderBookRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetSellOrderBookResponse{SellOrderBook: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetSellOrderBookRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetSellOrderBookResponse{SellOrderBook: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetSellOrderBookRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.SellOrderBook(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}

func TestSellOrderBookQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNSellOrderBook(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllSellOrderBookRequest {
		return &types.QueryAllSellOrderBookRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SellOrderBookAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SellOrderBook), step)
			require.Subset(t, msgs, resp.SellOrderBook)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.SellOrderBookAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.SellOrderBook), step)
			require.Subset(t, msgs, resp.SellOrderBook)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.SellOrderBookAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.SellOrderBookAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
