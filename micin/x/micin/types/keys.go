package types

const (
	// ModuleName defines the module name
	ModuleName = "micin"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_micin"
)

var (
	ParamsKey = []byte("p_micin")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
