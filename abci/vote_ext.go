package abci

import (
	"cosmossdk.io/log"
	"encoding/json"
	"fmt"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SusProposal defines a suspicious (sus) proposal that might be a scam
type SusProposal struct {
	HashedTitle string
	ScamPercent float64
}

// ScamProposalExtension defines the canonical vote extension structure for scam detection.
type ScamProposalExtension struct {
	Proposals []SusProposal
	Height    int64
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

		//proposals := make([]SusProposal, 0)
		// TODO: Convert this into a slice of MsgSubmitProposal
		// so we can detect more than one proposal
		var proposalMsg govtypes.MsgSubmitProposal
		for _, tx := range req.Txs {
			if err := h.cdc.Unmarshal(tx, &proposalMsg); err != nil {
				h.logger.Error(fmt.Sprintf("❌️ :: Transaction is not a gov proposal, %v", err))
				continue
			}

			//susProposal := SusProposal{
			//	HashedTitle: "",
			//	ScamPercent: 0,
			//}
			//proposals := append(proposals)
		}

		// Make an API call to OpenAI to compute the score for the proposal title and summary

		// produce a canonical vote extension
		// hash the title of the proposal
		// TODO hash the tile in the proposal with a nonce
		voteExtension := ScamProposalExtension{
			Proposals: []SusProposal{},
			Height:    req.Height,
		}

		bz, err := json.Marshal(voteExtension)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal vote extension: %w", err)
		}

		return &abci.ResponseExtendVote{VoteExtension: bz}, nil
	}
}

// VerifyVoteExtensionHandler handles the verification of the VoteExtensions provided by each validator.
// We are checking if the computed percent is the same for all validators
func (h *VoteExtHandler) VerifyVoteExtensionHandler() sdk.VerifyVoteExtensionHandler {
	return func(ctx sdk.Context, req *abci.RequestVerifyVoteExtension) (*abci.ResponseVerifyVoteExtension, error) {
		var voteExt ScamProposalExtension

		err := json.Unmarshal(req.VoteExtension, &voteExt)
		if err != nil {
			// NOTE: It is safe to return an error as the Cosmos SDK will capture all
			// errors, log them, and reject the proposal.
			return nil, fmt.Errorf("failed to unmarshal vote extension: %w", err)
		}

		if voteExt.Height != req.Height {
			return nil, fmt.Errorf("vote extension height does not match request height; expected: %d, got: %d", req.Height, voteExt.Height)
		}

		// TODO: Somehow verify the aggregate results of all validators and apply some

		return &abci.ResponseVerifyVoteExtension{Status: abci.ResponseVerifyVoteExtension_ACCEPT}, nil
	}
}
