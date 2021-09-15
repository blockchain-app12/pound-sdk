// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/osiz-blockchainapp/pound-sdk/x/crisis/types
package crisis

import (
	"github.com/osiz-blockchainapp/pound-sdk/x/crisis/internal/keeper"
	"github.com/osiz-blockchainapp/pound-sdk/x/crisis/internal/types"
)

const (
	DefaultCodespace  = types.DefaultCodespace
	CodeInvalidInput  = types.CodeInvalidInput
	ModuleName        = types.ModuleName
	DefaultParamspace = types.DefaultParamspace

	EventTypeInvariant   = types.EventTypeInvariant
	AttributeValueCrisis = types.AttributeValueCrisis
	AttributeKeyRoute    = types.AttributeKeyRoute
)

var (
	// functions aliases
	RegisterCodec         = types.RegisterCodec
	ErrNilSender          = types.ErrNilSender
	ErrUnknownInvariant   = types.ErrUnknownInvariant
	NewGenesisState       = types.NewGenesisState
	DefaultGenesisState   = types.DefaultGenesisState
	NewMsgVerifyInvariant = types.NewMsgVerifyInvariant
	ParamKeyTable         = types.ParamKeyTable
	NewInvarRoute         = types.NewInvarRoute
	NewKeeper             = keeper.NewKeeper

	// variable aliases
	ModuleCdc                = types.ModuleCdc
	ParamStoreKeyConstantFee = types.ParamStoreKeyConstantFee
)

type (
	GenesisState       = types.GenesisState
	MsgVerifyInvariant = types.MsgVerifyInvariant
	InvarRoute         = types.InvarRoute
	Keeper             = keeper.Keeper
)
