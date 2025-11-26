package keeper

import (
	"nebulix/x/storage/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NOTE: I know I can pull this from the bankkeeper but this is just way more efficient and we should never be changing this anyways....
const (
	unblxUnit int64 = 1000000
)

func (k Keeper) GetStorageCost(ctx sdk.Context, gbs int64, hours int64) sdk.Int {
	basePricePerGBHour := sdk.NewDec(k.Params.Get(ctx).PricePerGBHour)

	var finalPricePerGbHour sdk.Dec

	switch {
	case gbs >= 20_000:
		finalPricePerGbHour = basePricePerGBHour.Mul(sdk.MustNewDecFromStr("12.5").QuoInt64(15))
	case gbs >= 5_000:
		finalPricePerGbHour = basePricePerGBHour.Mul(sdk.NewDec(14).QuoInt64(15))
	default:
		finalPricePerGbHour = basePricePerGBHour
	}

	totalCost := finalPricePerGbHour.MulInt64(gbs).MulInt64(hours)

	nblxPrice := sdk.NewDec(3.5)
	nblxCost := totalCost.Quo(nblxPrice)
	unblxCost := nblxCost.MulInt64(unblxUnit)

	return unblxCost.TruncateInt()
}

// [TODO]
func (k Keeper) GetStoragePaymentInfo(
	ctx sdk.Context,
	address string,
) (val types.StoragePaymentInfo, found bool) {
	// store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoragePaymentInfoKeyPrefix))

	// b := store.Get(types.StoragePaymentInfoKey(
	// 	address,
	// ))
	// if b == nil {
	// 	return val, false
	// }

	// k.cdc.MustUnmarshal(b, &val)
	// k.FixStoragePaymentInfo(ctx, val)
	return val, true
}

// [TODO]
func (k Keeper) SetStoragePaymentInfo(ctx sdk.Context, payInfo types.StoragePaymentInfo) {
	// store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoragePaymentInfoKeyPrefix))
	// b := k.cdc.MustMarshal(&payInfo)
	// store.Set(types.StoragePaymentInfoKey(
	// 	payInfo.Address,
	// ), b)
}
