package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Bin struct {
	Id        string `json:"id"`
	Root      []byte `json:"root"`
	Provider  string `json:"provider"` // [TODO]: use sdk.AccAddress
	FileCount uint64 `json:"file_count"`
	// Status      BinStatus `json:"status"`  // [TBD]: do I want to challenge entire bins instead of files?
	CreatedAt   int64    `json:"created_at"` // block height
	LastUpdated int64    `json:"last_updated"`
	Deposit     sdk.Coin `json:"deposit"`
}

type BinStatus int32

const (
	BinStatusActive BinStatus = iota
	BinStatusInactive
	BinStatusChallenged
	BinStatusSlashed
)

// Key for Bin Root store
func GetBinRootKey(binId uint16) []byte {
	key := make([]byte, 2)
	// [NOTE]: BigEndian for consistent byte ordering... learned the hard way...
	binary.BigEndian.PutUint16(key, binId)
	return key
}
