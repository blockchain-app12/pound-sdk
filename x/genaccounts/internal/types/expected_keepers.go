package types

import (
	sdk "github.com/osiz-blockchainapp/pound-sdk/types"
	authexported "github.com/osiz-blockchainapp/pound-sdk/x/auth/exported"
)

// AccountKeeper defines the expected account keeper (noalias)
type AccountKeeper interface {
	NewAccount(sdk.Context, authexported.Account) authexported.Account
	SetAccount(sdk.Context, authexported.Account)
	IterateAccounts(ctx sdk.Context, process func(authexported.Account) (stop bool))
}
