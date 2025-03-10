package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateRollapp{}, "rollapp/CreateRollapp", nil)
	cdc.RegisterConcrete(&MsgUpdateRollappInformation{}, "rollapp/UpdateRollappInformation", nil)
	cdc.RegisterConcrete(&MsgTransferOwnership{}, "rollapp/TransferDymNameOwnership", nil)
	cdc.RegisterConcrete(&MsgUpdateState{}, "rollapp/UpdateState", nil)
	cdc.RegisterConcrete(&MsgAddApp{}, "rollapp/AddApp", nil)
	cdc.RegisterConcrete(&MsgUpdateApp{}, "rollapp/UpdateApp", nil)
	cdc.RegisterConcrete(&MsgRemoveApp{}, "rollapp/RemoveApp", nil)
	cdc.RegisterConcrete(&MsgRollappFraudProposal{}, "rollapp/RollappFraudProposal", nil)
	cdc.RegisterConcrete(&MsgMarkObsoleteRollapps{}, "rollapp/MarkObsoleteRollapps", nil)
	cdc.RegisterConcrete(&MsgForceGenesisInfoChange{}, "rollapp/ForceGenesisInfoChange", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateRollapp{},
		&MsgUpdateRollappInformation{},
		&MsgTransferOwnership{},
		&MsgUpdateState{},
		&MsgAddApp{},
		&MsgUpdateApp{},
		&MsgRemoveApp{},
		&MsgRollappFraudProposal{},
		&MsgMarkObsoleteRollapps{},
		&MsgForceGenesisInfoChange{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(Amino)
)

func init() {
	RegisterCodec(Amino)
	// Register all Amino interfaces and concrete types on the authz Amino codec so that this can later be
	// used to properly serialize MsgGrant and MsgExec instances
	sdk.RegisterLegacyAminoCodec(Amino)
	RegisterCodec(authzcodec.Amino)

	Amino.Seal()
}
