package keeper

import (
	"context"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
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

	if msg.Recipient == "" {
		msg.Recipient = msg.Creator
	}

	recipient, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, sdkerrors.Wrap(err, "for address is not a proper bech32")
	}

	_, found := k.GetStoragePaymentInfo(ctx, recipient.String())
	if found {
		return nil, sdkerrors.Wrap(err, "account has an active storage subscription, consider upgrading storage")
	}

	duration, gbs, err := validateBuy(msg.Duration, msg.Bytes, msg.PaymentDenom)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to validate buy request")
	}

	hours := sdk.NewDec(duration.Milliseconds()).Quo(sdk.NewDec(60 * 60 * 1000))
	storageCost := k.GetStorageCost(ctx, gbs, hours.TruncateInt().Int64())

	toPay := sdk.NewCoin(msg.PaymentDenom, storageCost)

	// TODO: process payment & generate storage credits
	accExists := k.authKeeper.HasAccount(ctx, recipient)
	if !accExists {
		defer telemetry.IncrCounter(1, "new", "account") // TBD: do I want/need this?
		k.authKeeper.SetAccount(ctx, k.authKeeper.NewAccountWithAddress(ctx, recipient))
	}

	payInfo := types.StoragePaymentInfo{
		Start:          ctx.BlockTime(),
		End:            ctx.BlockTime().Add(duration),
		SpaceAvailable: bytes,
		SpaceUsed:      spaceUsed,
		Address:        forAddress.String(),
	}

	return &types.MsgBuyStorageResponse{}, nil
}

// TODO: upgrade storage duration
// TODO: upgrade storage quantity
