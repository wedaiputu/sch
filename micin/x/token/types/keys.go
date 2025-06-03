package types

const (
	// ModuleName defines the module name
	ModuleName = "token"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_token"
)

var (
	ParamsKey = []byte("p_token")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	CoinKey      = "Coin/value/"
	CoinCountKey = "Coin/count/"
)
