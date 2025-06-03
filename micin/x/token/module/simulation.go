package token

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"micin/testutil/sample"
	tokensimulation "micin/x/token/simulation"
	"micin/x/token/types"
)

// avoid unused import issue
var (
	_ = tokensimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateCoin = "op_weight_msg_coin"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCoin int = 100

	opWeightMsgUpdateCoin = "op_weight_msg_coin"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateCoin int = 100

	opWeightMsgDeleteCoin = "op_weight_msg_coin"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteCoin int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	tokenGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		CoinList: []types.Coin{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		CoinCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&tokenGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateCoin int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateCoin, &weightMsgCreateCoin, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCoin = defaultWeightMsgCreateCoin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCoin,
		tokensimulation.SimulateMsgCreateCoin(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateCoin int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateCoin, &weightMsgUpdateCoin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateCoin = defaultWeightMsgUpdateCoin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateCoin,
		tokensimulation.SimulateMsgUpdateCoin(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteCoin int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteCoin, &weightMsgDeleteCoin, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteCoin = defaultWeightMsgDeleteCoin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteCoin,
		tokensimulation.SimulateMsgDeleteCoin(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateCoin,
			defaultWeightMsgCreateCoin,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokensimulation.SimulateMsgCreateCoin(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateCoin,
			defaultWeightMsgUpdateCoin,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokensimulation.SimulateMsgUpdateCoin(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteCoin,
			defaultWeightMsgDeleteCoin,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				tokensimulation.SimulateMsgDeleteCoin(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
