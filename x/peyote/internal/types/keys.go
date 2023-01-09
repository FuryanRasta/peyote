package types

const (
	// ModuleName is the name of this module
	ModuleName = "peyote"

	// StoreKey is the default store key for this module
	StoreKey = ModuleName

	// DefaultParamspace is the default param space for this module
	DefaultParamspace = ModuleName

	// BondsMintBurnAccount the root string for the peyote mint burn account address
	BondsMintBurnAccount = "peyote_mint_burn_account"

	// BatchesIntermediaryAccount the root string for the batches account address
	BatchesIntermediaryAccount = "batches_intermediary_account"

	// BondsReserveAccount the root string for the peyote reserve account address
	BondsReserveAccount = "peyote_reserve_account"

	// QuerierRoute is the querier route for this module's store.
	QuerierRoute = ModuleName

	// RouterKey is the message route for this module
	RouterKey = ModuleName
)

// Bonds and batches are stored as follow:
//
// - Bonds: 0x00<bond_token_bytes>
// - Batches: 0x01<bond_token_bytes>
// - Last batches: 0x02<bond_token_bytes>
var (
	BondsKeyPrefix       = []byte{0x00} // key for peyote
	BatchesKeyPrefix     = []byte{0x01} // key for batches
	LastBatchesKeyPrefix = []byte{0x02} // key for last batches
)

func GetBondKey(token string) []byte {
	return append(BondsKeyPrefix, []byte(token)...)
}

func GetBatchKey(token string) []byte {
	return append(BatchesKeyPrefix, []byte(token)...)
}

func GetLastBatchKey(token string) []byte {
	return append(LastBatchesKeyPrefix, []byte(token)...)
}
