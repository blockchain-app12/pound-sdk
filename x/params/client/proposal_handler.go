package client

import (
	govclient "github.com/osiz-blockchainapp/pound-sdk/x/gov/client"
	"github.com/osiz-blockchainapp/pound-sdk/x/params/client/cli"
	"github.com/osiz-blockchainapp/pound-sdk/x/params/client/rest"
)

// param change proposal handler
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
