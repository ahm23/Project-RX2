package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/store/types"
	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"nebulix/x/storage/keeper"
	module "nebulix/x/storage/module"
	"nebulix/x/storage/types"
)

// [TODO] pump this into deepseek; ask how to handle the expected keepers as well (bankKeeper & authKeeper)
type KeeperTestSuite struct {
	suite.Suite

	ctx          sdk.Context
	keeper       keeper.Keeper
	addressCodec address.Codec
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.reset()
}

func (suite *KeeperTestSuite) reset() {
	encCfg := moduletestutil.MakeTestEncodingConfig(module.AppModule{})
	addressCodec := addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix())
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

	storeService := runtime.NewKVStoreService(storeKey)
	ctx := testutil.DefaultContextWithDB(suite.T(), storeKey, storetypes.NewTransientStoreKey("transient_test")).Ctx

	authority := authtypes.NewModuleAddress(types.GovModuleName)

	k := keeper.NewKeeper(
		storeService,
		encCfg.Codec,
		addressCodec,
		authority,
		nil,
		nil,
	)

	// Initialize params
	if err := k.Params.Set(ctx, types.DefaultParams()); err != nil {
		suite.T().Fatalf("failed to set params: %v", err)
	}

	suite.ctx = ctx
	suite.keeper = k
	suite.addressCodec = addressCodec
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}
