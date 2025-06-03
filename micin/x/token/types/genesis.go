package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		CoinList: []Coin{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in coin
	coinIdMap := make(map[uint64]bool)
	coinCount := gs.GetCoinCount()
	for _, elem := range gs.CoinList {
		if _, ok := coinIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for coin")
		}
		if elem.Id >= coinCount {
			return fmt.Errorf("coin id should be lower or equal than the last id")
		}
		coinIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
