package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"nebulix/x/storage/types"
)

const (
	gb_bytes int64 = 1_000_000_000
)

func validateBuy(days int64, bytes int64, denom string) (duration time.Duration, gb int64, err error) {
	if denom != "ujkl" {
		err = sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "cannot pay with anything other than ujkl")
		return
	}

	duration = time.Duration(days) * time.Hour * 24
	if duration < time.Hour*24 {
		err = fmt.Errorf("duration can't be less than 1 day")
		return
	}

	gb = bytes / gb_bytes
	if gb <= 0 {
		err = fmt.Errorf("cannot buy less than a gb")
		return
	}

	return
}

func (k msgServer) BuyStorage(goCtx context.Context, msg *types.MsgBuyStorage) (*types.MsgBuyStorageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.Params.Get(ctx)

	if msg.ForAddress == "" {
		msg.ForAddress = msg.Creator
	}

	forAddr, err := sdk.AccAddressFromBech32(msg.ForAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "for address is not a proper bech32")
	}

	duration, gbs, err := validateBuy(msg.Duration, msg.Bytes, msg.PaymentDenom)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to validate buy request")
	}

	hours := sdk.NewDec(duration.Milliseconds()).Quo(sdk.NewDec(60 * 60 * 1000))
	storageCost := k.GetStorageCost(ctx, gbs, hours.TruncateInt().Int64())

	toPay := sdk.NewCoin(msg.PaymentDenom, storageCost)

	// TODO: process payment & generate storage credits

	return &types.MsgBuyStorageResponse{}, nil
}

// TODO: upgrade storage duration
// TODO: upgrade storage quantity
