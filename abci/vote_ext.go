package abci

import (
	"cosmossdk.io/log"
	"fmt"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ScamProposalExtension defines the canonical vote extension structure for scam detection.
type ScamProposalExtension struct {
	ProposalID  uint64
	ScamPercent uint64
}

func NewVoteExtensionHandler(lg log.Logger, cdc codec.Codec) *VoteExtHandler {
	return &VoteExtHandler{
		logger: lg,
		cdc:    cdc,
	}
}

func (h *VoteExtHandler) ExtendVoteHandler() sdk.ExtendVoteHandler {
	return func(ctx sdk.Context, req *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
		h.logger.Info(fmt.Sprintf("Extending votes at block height : %v", req.Height))

		// TODO: Convert this into a slice of MsgSubmitProposal
		// so we can detect more than one proposal
		var proposalMsg govtypes.MsgSubmitProposal
		for _, tx := range req.Txs {
			if err := h.cdc.Unmarshal(tx, &proposalMsg); err != nil {
				h.logger.Error(fmt.Sprintf("❌️ :: Transaction is not a gov proposal, %v", err))
			}
		}

		// TODO: API call with the description and title of the proposal
		// result := detector.Detect(
		//	proposalMsg.Description,
		// 	proposalMsg.Title,
		//)

		// produce a canonical vote extension
		// TODO:
		//_ := ScamProposalExtension{
		//	ProposalID:  1,
		//	ScamPercent: 100,
		//}

		return &abci.ResponseExtendVote{VoteExtension: []byte{}}, nil
	}
}

func (h *VoteExtHandler) VerifyVoteExtensionHandler() sdk.VerifyVoteExtensionHandler {
	return func(ctx sdk.Context, req *abci.RequestVerifyVoteExtension) (*abci.ResponseVerifyVoteExtension, error) {

		return &abci.ResponseVerifyVoteExtension{Status: abci.ResponseVerifyVoteExtension_ACCEPT}, nil
	}
}
