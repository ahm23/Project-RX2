// x/storage/keeper/keys.go
package keeper

// Prefix keys for different types of data
// These are arbitrary bytes you choose - just make them unique!
var (
	BinKey       = []byte{0x01} // You choose: 0x01
	FileKey      = []byte{0x02} // You choose: 0x02
	ProviderKey  = []byte{0x03} // You choose: 0x03
	ChallengeKey = []byte{0x04} // You choose: 0x04
	ParamsKey    = []byte{0x05} // You choose: 0x05
)

// Key creation functions
func KeyBin(binID string) []byte {
	// Format: 0x01 + binID
	return append(BinKey, []byte(binID)...)
}

func KeyFile(fileID string) []byte {
	// Format: 0x02 + fileID
	return append(FileKey, []byte(fileID)...)
}

func KeyProvider(providerAddr string) []byte {
	// Format: 0x03 + providerAddress
	return append(ProviderKey, []byte(providerAddr)...)
}

func KeyChallenge(challengeID string) []byte {
	// Format: 0x04 + challengeID
	return append(ChallengeKey, []byte(challengeID)...)
}

// Secondary indexes for efficient queries
func KeyFilesByBin(binID string) []byte {
	// Format: "files_bin_" + binID
	return []byte("files_bin_" + binID)
}

func KeyBinsByProvider(providerAddr string) []byte {
	// Format: "bins_provider_" + providerAddr
	return []byte("bins_provider_" + providerAddr)
}

func KeyActiveChallengesByBin(binID string) []byte {
	// Format: "challenges_bin_" + binID
	return []byte("challenges_bin_" + binID)
}

// Helper to extract IDs from keys (useful for iteration)
func ExtractBinIDFromKey(key []byte) string {
	// Remove the prefix (0x01) to get the binID
	return string(key[len(BinKey):])
}

func ExtractFileIDFromKey(key []byte) string {
	// Remove the prefix (0x02) to get the fileID
	return string(key[len(FileKey):])
}
