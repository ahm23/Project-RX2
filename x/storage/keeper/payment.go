package keeper

import (
	"nebulix/x/storage/types"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NOTE: I know I can pull this from the bankkeeper but this is just way more efficient and we should never be changing this anyways....
const (
	unblxUnit int64 = 1000000
)

func (k Keeper) GetStorageCost(ctx sdk.Context, gbs int64, days int64) math.Int {
	params, _ := k.Params.Get(ctx)
	pricePerGBDay := math.NewInt(params.PricePerGBDay)

	// [TBD]: verify this can't overflow in an attack attempt
	totalCost := pricePerGBDay.MulRaw(gbs).MulRaw(days)

	// [TODO]: oracle for price fetching
	nblxPrice, _ := math.LegacyNewDecFromStr("1000")
	// ^ temporary price for testing
	nblxCost := math.LegacyDec(totalCost).Quo(nblxPrice)
	unblxCost := math.NewInt(nblxCost.MulInt64(unblxUnit).Ceil().BigInt().Int64())

	return unblxCost
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
