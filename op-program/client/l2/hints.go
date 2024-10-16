package l2

import (
	"encoding/binary"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"

	preimage "github.com/ethereum-optimism/optimism/op-preimage"
)

const (
	HintL2BlockHeader      = "l2-block-header"
	HintL2Transactions     = "l2-transactions"
	HintL2Code             = "l2-code"
	HintL2StateNode        = "l2-state-node"
	HintL2Output           = "l2-output"
	HintL2AccountProof     = "l2-account-proof"
	HintL2ExecutionWitness = "l2-execution-witness"
)

type BlockHeaderHint common.Hash

var _ preimage.Hint = BlockHeaderHint{}

func (l BlockHeaderHint) Hint() string {
	return HintL2BlockHeader + " " + (common.Hash)(l).String()
}

type TransactionsHint common.Hash

var _ preimage.Hint = TransactionsHint{}

func (l TransactionsHint) Hint() string {
	return HintL2Transactions + " " + (common.Hash)(l).String()
}

type CodeHint common.Hash

var _ preimage.Hint = CodeHint{}

func (l CodeHint) Hint() string {
	return HintL2Code + " " + (common.Hash)(l).String()
}

type StateNodeHint common.Hash

var _ preimage.Hint = StateNodeHint{}

func (l StateNodeHint) Hint() string {
	return HintL2StateNode + " " + (common.Hash)(l).String()
}

type L2OutputHint common.Hash

var _ preimage.Hint = L2OutputHint{}

func (l L2OutputHint) Hint() string {
	return HintL2Output + " " + (common.Hash)(l).String()
}

type AccountProofHint struct {
	BlockNumber uint64
	Address     common.Hash
}

var _ preimage.Hint = AccountProofHint{}

func (l AccountProofHint) Hint() string {
	var blockNumBytes [4]byte

	binary.BigEndian.PutUint64(blockNumBytes[:], l.BlockNumber)

	return HintL2AccountProof + " " + hex.EncodeToString(blockNumBytes[:]) + " " + (common.Hash)(l.Address).String()
}

type ExecutionWitnessHint common.Hash

var _ preimage.Hint = ExecutionWitnessHint{}

func (l ExecutionWitnessHint) Hint() string {
	return HintL2ExecutionWitness + " " + (common.Hash)(l).String()
}
