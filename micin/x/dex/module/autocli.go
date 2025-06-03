package dex

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "micin/api/micin/dex"
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
					RpcMethod: "PoolAll",
					Use:       "list-pool",
					Short:     "List all pool",
				},
				{
					RpcMethod:      "Pool",
					Use:            "show-pool [id]",
					Short:          "Shows a pool by id",
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
					RpcMethod:      "CreatePool",
					Use:            "create-pool [tokenA] [tokenB] [reserveA] [reserveB]",
					Short:          "Create pool",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "tokenA"}, {ProtoField: "tokenB"}, {ProtoField: "reserveA"}, {ProtoField: "reserveB"}},
				},
				{
					RpcMethod:      "UpdatePool",
					Use:            "update-pool [id] [tokenA] [tokenB] [reserveA] [reserveB]",
					Short:          "Update pool",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "tokenA"}, {ProtoField: "tokenB"}, {ProtoField: "reserveA"}, {ProtoField: "reserveB"}},
				},
				{
					RpcMethod:      "DeletePool",
					Use:            "delete-pool [id]",
					Short:          "Delete pool",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "RemoveLiquidity",
					Use:            "remove-liquidity [sender] [token-a] [token-b] [share]",
					Short:          "Send a remove-liquidity tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "sender"}, {ProtoField: "tokenA"}, {ProtoField: "tokenB"}, {ProtoField: "share"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
