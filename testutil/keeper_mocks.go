package testutil

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SentCoins struct {
	fromAddr sdk.AccAddress
	toModule string
	coins    sdk.Coins
}

// MockBankKeeper implements bank keeper interface for tests
type MockBankKeeper struct {
	balances map[string]sdk.Coins
	sent     []SentCoins
}

func (m *MockBankKeeper) SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error {
	// In real implementation, you'd update balances
	m.sent = append(m.sent, SentCoins{
		fromAddr: senderAddr,
		toModule: recipientModule,
		coins:    amt,
	})
	return nil
}

func (m *MockBankKeeper) SetBalance(addr string, coins sdk.Coins) {
	m.balances[addr] = coins
}
