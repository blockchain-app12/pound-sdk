package rest

import (
	"github.com/gorilla/mux"

	"github.com/osiz-blockchainapp/pound-sdk/client/context"
)

// RegisterRoutes registers minting module REST handlers on the provided router.
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
}