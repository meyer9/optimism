package host

import (
	"context"
	"encoding/binary"

	"github.com/ethereum-optimism/optimism/op-program/host/prefetcher"
	"github.com/ethereum-optimism/optimism/op-program/host/types"
	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type L2ExperimentalClient struct {
	*L2Client
	client client.RPC
}

var _ prefetcher.L2Source = &L2ExperimentalClient{}

func NewL2ExperimentalClient(l2Client *L2Client, client client.RPC) *L2ExperimentalClient {
	return &L2ExperimentalClient{
		L2Client: l2Client,
		client:   client,
	}
}

// CodeByHash implements prefetcher.L2Source.
func (s *L2ExperimentalClient) CodeByHash(ctx context.Context, hash common.Hash) ([]byte, error) {
	panic("unsupported")
}

// NodeByHash implements prefetcher.L2Source.
func (s *L2ExperimentalClient) NodeByHash(ctx context.Context, hash common.Hash) ([]byte, error) {
	panic("unsupported")
}

func (s *L2ExperimentalClient) ExecutionWitness(ctx context.Context, blockNum uint64) (*types.ExecutionWitness, error) {
	var witness types.ExecutionWitness

	var blockNumBytes [8]byte
	binary.BigEndian.PutUint64(blockNumBytes[:], blockNum)

	err := s.client.CallContext(ctx, &witness, "debug_executionWitness", hexutil.Encode(blockNumBytes[:]), "true")
	if err != nil {
		return nil, err
	}
	return &witness, nil
}
