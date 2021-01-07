package mytest

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/my-test/x/mytest/types"
	"github.com/my-test/x/mytest/keeper"
)

func handleMsgCreateCommit(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateCommit) (*sdk.Result, error) {
	var commit = types.Commit{
		Creator: msg.Creator,
		ID:      msg.ID,
    SolutionHash: msg.SolutionHash,
    SolutionScavengerHash: msg.SolutionScavengerHash,
	}
	k.CreateCommit(ctx, commit)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
