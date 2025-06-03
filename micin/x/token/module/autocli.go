package token

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "micin/api/micin/token"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "CoinAll",
					Use:       "list-coin",
					Short:     "List all coin",
				},
				{
					RpcMethod:      "Coin",
					Use:            "show-coin [id]",
					Short:          "Shows a coin by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateCoin",
					Use:            "create-coin [name] [supply] [owner]",
					Short:          "Create coin",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "supply"}, {ProtoField: "owner"}},
				},
				{
					RpcMethod:      "UpdateCoin",
					Use:            "update-coin [id] [name] [supply] [owner]",
					Short:          "Update coin",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "name"}, {ProtoField: "supply"}, {ProtoField: "owner"}},
				},
				{
					RpcMethod:      "DeleteCoin",
					Use:            "delete-coin [id]",
					Short:          "Delete coin",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
