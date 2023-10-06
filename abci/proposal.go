package abci

import (
	"cosmossdk.io/log"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

// ScamProposalTx defines the custom transaction identifying the scam proposal by its ID.
type ScamProposalTx struct {
	ProposalID uint64
	IsScam     bool
}

// NewProposalHandler creates a new instance of the handler to be used.
func NewProposalHandler(
	lg log.Logger,
	valStore baseapp.ValidatorStore,
	cdc codec.Codec,
	govKeeper govkeeper.Keeper,
	stakingKeeper stakingkeeper.Keeper,
) *ProposalHandler {
	return &ProposalHandler{
		logger:        lg,
		valStore:      valStore,
		cdc:           cdc,
		govKeeper:     govKeeper,
		stakingKeeper: stakingKeeper,
	}
}

// PrepareProposalHandler is the handler to be used for PrepareProposal.
func (h *ProposalHandler) PrepareProposalHandler() sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
		err := baseapp.ValidateVoteExtensions(ctx, h.valStore, req.Height, ctx.ChainID(), req.LocalLastCommit)
		if err != nil {
			return nil, err
		}

		//_ := req.Txs

		if req.Height >= ctx.ConsensusParams().Abci.VoteExtensionsEnableHeight {
		}

		// TODO: API call with the description and title of the proposal
		// detector.Detect(
		//	proposalMsg.Description,
		// 	proposalMsg.Title,
		//)

		return nil, nil
	}
}

// ProcessProposalHandler is the handler to be used for ProcessProposal.
func (h *ProposalHandler) ProcessProposalHandler() sdk.ProcessProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestProcessProposal) (resp *abci.ResponseProcessProposal, err error) {
		resp.Status = 1 // Accepts the proposal
		return resp, nil
	}
}

// computeScamIdentificationResults aggregates the scam identification results from each validator.
func computeScamIdentificationResults(ctx sdk.Context, ci abci.ExtendedCommitInfo) (bool, error) {
	// Get all the votes from the commit info
	//votes := ci.Votes

	//for i, vote := range votes {
	//	vote.VoteExtension
	//}

	// Compute the average of all vote percentages
	// If the average is greater than the threshold, return true
	// Otherwise, return false
	return false, nil

}
