package keeper

import (
	"context"
	"encoding/hex"
	"fmt"

	"nebulix/x/storage/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PostFile(goCtx context.Context, msg *types.MsgPostFile) (*types.MsgPostFileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// [TODO]: bad file msg checks

	file := types.File{
		Merkle:    hex.EncodeToString(msg.Merkle),
		Owner:     msg.Creator,
		Start:     ctx.BlockHeight(),
		FileSize:  msg.FileSize,
		Providers: msg.Providers, // [TODO]: ensure first challenge doesn't slash, just removes, to avoid framing providers
		Replicas:  msg.Replicas,
	}

	k.SetFile(ctx, file)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSignContract,
			sdk.NewAttribute(types.AttributeKeySigner, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyContract, hex.EncodeToString(msg.Merkle)),
			sdk.NewAttribute(types.AttributeKeyStart, fmt.Sprintf("%d", ctx.BlockHeight())),
		),
	)

	// paymentInfo, found := k.GetStoragePaymentInfo(ctx, msg.Creator)
	// if !found {
	// 	return nil, sdkerrors.ErrKeyNotFound.Wrapf("storage account does not exist")
	// }
	// if paymentInfo.End.Before(ctx.BlockTime()) {
	// 	return nil, sdkerrors.ErrUnauthorized.Wrapf("storage account is expired")
	// }

	// paymentInfo.SpaceUsed += msg.FileSize
	// if paymentInfo.SpaceUsed > paymentInfo.SpaceAvailable {
	// 	return nil, sdkerrors.ErrUnauthorized.Wrapf("storage account does not have enough space available %d > %d", paymentInfo.SpaceUsed, paymentInfo.SpaceAvailable)
	// }

	// k.SetStoragePaymentInfo(ctx, paymentInfo)

	return &types.MsgPostFileResponse{StartBlock: ctx.BlockHeight()}, nil
}
