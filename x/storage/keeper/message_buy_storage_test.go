package keeper

import (
	"nebulix/x/storage/keeper"
	"nebulix/x/storage/types"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func (suite *KeeperTestSuite) TestMsgBuyStorage() {
	ms := keeper.NewMsgServerImpl(suite.keeper)

	params := types.DefaultParams()
	require.NoError(suite.T(), suite.keeper.Params.Set(suite.ctx, params))

	// _, err := suite.addressCodec.BytesToString(suite.keeper.GetAuthority())
	// suite.Require().NoError(err)

	testAddresses, err := testutil.CreateTestAddresses("cosmos", 2)
	suite.Require().NoError(err)

	testAccount := testAddresses[0]
	depoAccount := testAddresses[1]

	coins := sdk.NewCoins(sdk.NewCoin("ujkl", sdk.NewInt(100000000000))) // Send some coins to their account
	testAcc, _ := sdk.AccAddressFromBech32(testAccount)
	err = suite.bankKeeper.SendCoinsFromModuleToAccount(suite.ctx, types.ModuleName, testAcc, coins)
	suite.Require().NoError(err)

	suite.storageKeeper.SetParams(suite.ctx, types.Params{
		DepositAccount:         depoAccount,
		ProofWindow:            50,
		ChunkSize:              1024,
		PriceFeed:              "jklprice",
		MissesToBurn:           3,
		MaxContractAgeInBlocks: 100,
		PricePerTbPerMonth:     15,
		CollateralPrice:        2,
		CheckWindow:            11,
		ReferralCommission:     25,
		PolRatio:               40,
	})

	// default params
	testCases := []struct {
		name      string
		input     *types.MsgBuyStorage
		expErr    bool
		expErrMsg string
	}{
		{
			name: "invalid creator address",
			input: &types.MsgBuyStorage{
				Creator:  "not_an_address",
				Receiver: "nblxj3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Duration: 30,
				Bytes:    2048,
			},
			expErr:    true,
			expErrMsg: sdkerrors.ErrInvalidAddress.Error(),
		},
		{
			name: "invalid receiver address",
			input: &types.MsgBuyStorage{
				Creator:  "nblxj3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Receiver: "not_an_address",
				Duration: 30,
				Bytes:    2048,
			},
			expErr:    true,
			expErrMsg: sdkerrors.ErrInvalidAddress.Error(),
		},
		{
			name: "invalid duration (0 days)",
			input: &types.MsgBuyStorage{
				Creator:  "nblx1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Receiver: "nblx1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Duration: 0,
				Bytes:    2048,
			},
			expErr:    true,
			expErrMsg: sdkerrors.ErrInvalidRequest.Error(),
		},
		{
			name: "invalid duration (negative days)",
			input: &types.MsgBuyStorage{
				Creator:  "nblx1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Receiver: "nblx1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Duration: -1,
				Bytes:    2048,
			},
			expErr:    true,
			expErrMsg: sdkerrors.ErrInvalidRequest.Error(),
		},
		{
			name: "valid address",
			input: &types.MsgBuyStorage{
				Creator:  "nblx1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Receiver: "nblx1j3p63s42w7ywaczlju626st55mzu5z399f5n6n",
				Duration: 30,
				Bytes:    2048,
			},
		},
	}

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			_, err := ms.UpdateParams(suite.ctx, tc.input)

			if tc.expErr {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expErrMsg)
			} else {
				suite.Require().NoError(err)
			}
		})
	}
}

func TestMsgBuyStorage_ValidateBasic(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
