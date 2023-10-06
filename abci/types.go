package abci

import (
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

// ProposalHandler defines the ProposalHandler to be used for
// PrepareProposal, ProcessProposal and PreBlocker
type ProposalHandler struct {
	logger        log.Logger
	valStore      baseapp.ValidatorStore
	cdc           codec.Codec
	govKeeper     govkeeper.Keeper
	stakingKeeper stakingkeeper.Keeper
}

type VoteExtHandler struct {
	logger       log.Logger
	currentBlock int64
	cdc          codec.Codec
}

type InjectedVoteExt struct {
	VoteExtSigner []byte
	Bids          [][]byte
}

type InjectedVotes struct {
	Votes []InjectedVoteExt
}

type AppVoteExtension struct {
	Height int64
	Bids   [][]byte
}

type SpecialTransaction struct {
	Height int
	Bids   [][]byte
}
