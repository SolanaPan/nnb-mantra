package module

import (
	"github.com/SolanaPan/nnb/v1/x/kyc/keeper"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

type AppModule struct {
	AppModuleBasic

	keeper keeper.Keeper
}

// IsAppModule implements module.AppModule.
func (a AppModule) IsAppModule() {
	panic("unimplemented")
}

// IsOnePerModuleType implements module.AppModule.
func (a AppModule) IsOnePerModuleType() {
	panic("unimplemented")
}

// Name implements module.AppModule.
func (a AppModule) Name() string {
	panic("unimplemented")
}

// RegisterGRPCGatewayRoutes implements module.AppModule.
func (a AppModule) RegisterGRPCGatewayRoutes(client.Context, *runtime.ServeMux) {
	panic("unimplemented")
}

// RegisterInterfaces implements module.AppModule.
func (a AppModule) RegisterInterfaces(types.InterfaceRegistry) {
	panic("unimplemented")
}

// RegisterLegacyAminoCodec implements module.AppModule.
func (a AppModule) RegisterLegacyAminoCodec(*codec.LegacyAmino) {
	panic("unimplemented")
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
	}
}

// Implement AppModuleBasic, AppModule, etc. as in other modules
