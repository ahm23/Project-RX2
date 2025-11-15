package keeper

import (
	"context"

	"nebulix/x/storage/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) InitProvider(ctx context.Context, msg *types.MsgInitProvider) (*types.MsgInitProviderResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgInitProviderResponse{}, nil
}
