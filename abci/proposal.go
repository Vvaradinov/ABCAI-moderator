package abci

import (
	"cosmossdk.io/log"
	"fmt"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

func NewPrepareProposalHandler(
	lg log.Logger,
	txCg client.TxConfig,
	cdc codec.Codec,
	runProv bool,
) *PrepareProposalHandler {
	return &PrepareProposalHandler{
		logger:      lg,
		txConfig:    txCg,
		cdc:         cdc,
		runProvider: runProv,
	}
}

func (h *PrepareProposalHandler) PrepareProposalHandler() sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestPrepareProposal) (resp *abci.ResponsePrepareProposal, err error) {
		// The transactions in the proposal
		transactions := req.Txs

		//if req.Height >= ctx.ConsensusParams().Abci.VoteExtensionsEnableHeight {
		// TODO: Here you have to check the Vote Extensions from the previous block
		//}

		// TODO: Convert this into a slice of MsgSubmitProposal
		// so we can detect more than one proposal
		var proposalMsg govtypes.MsgSubmitProposal
		for _, tx := range transactions {
			if err := h.cdc.Unmarshal(tx, &proposalMsg); err != nil {
				h.logger.Error(fmt.Sprintf("❌️ :: Transaction is not a gov proposal, %v", err))
			}
		}

		// TODO: API call with the description and title of the proposal
		// detector.Detect(
		//	proposalMsg.Description,
		// 	proposalMsg.Title,
		//)

		return resp, nil
	}
}

func (h *ProcessProposalHandler) ProcessProposalHandler() sdk.ProcessProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestProcessProposal) (resp *abci.ResponseProcessProposal, err error) {
		resp.Status = 1 // Accepts the proposal
		return resp, nil
	}
}
