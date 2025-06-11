package types

const (
	// ModuleName defines the module name
	ModuleName = "micin"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_micin"

	// RouterKey is used to route messages
	RouterKey = ModuleName

	// QuerierRoute is used for queries
	QuerierRoute = ModuleName
)

var (
	ParamsKey = []byte("p_micin")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
