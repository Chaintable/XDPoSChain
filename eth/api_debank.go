package eth

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	ptracer "github.com/Chaintable/pipeline/tracer"
	ptypes "github.com/Chaintable/pipeline/types"
	"github.com/Chaintable/pipeline/util"
	"github.com/XinFinOrg/XDPoSChain/XDCx/tradingstate"
	"github.com/XinFinOrg/XDPoSChain/common"
	"github.com/XinFinOrg/XDPoSChain/common/hexutil"
	"github.com/XinFinOrg/XDPoSChain/consensus/misc"
	"github.com/XinFinOrg/XDPoSChain/core"
	"github.com/XinFinOrg/XDPoSChain/core/state"
	"github.com/XinFinOrg/XDPoSChain/core/types"
	"github.com/XinFinOrg/XDPoSChain/core/vm"
	"github.com/XinFinOrg/XDPoSChain/log"
	"github.com/XinFinOrg/XDPoSChain/params"
	"github.com/XinFinOrg/XDPoSChain/rpc"
)

type DebankOutPutRaw struct {
	BlockFile      *ptypes.BlockFile        `json:"block_file"`
	Header         *ptypes.Header           `json:"header"`
	StateDiff      *ptypes.BlockStorageDiff `json:"state_diff"`
	ValidationHash int64                    `json:"validation_hash"`
}

type DebankAPI struct {
	eth *Ethereum
}

func NewDebankAPI(eth *Ethereum) *DebankAPI {
	return &DebankAPI{
		eth: eth,
	}
}

func (api *DebankAPI) DebankBlockRaw(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*DebankOutPutRaw, error) {
	output, err := api.DebankBlock(ctx, blockNrOrHash)
	if err != nil {
		return nil, err
	}
	var stateDiff ptypes.BlockStorageDiff
	err = util.DecodeFromRlp(output.StateDiff, &stateDiff)
	if err != nil {
		return nil, err
	}
	return &DebankOutPutRaw{
		BlockFile:      output.BlockFile,
		Header:         output.Header,
		StateDiff:      &stateDiff,
		ValidationHash: output.ValidationHash,
	}, nil
}

func (api *DebankAPI) DebankBlock(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (*ptypes.DebankOutPut, error) {
	block, err := api.eth.ApiBackend.BlockByNumberOrHash(ctx, blockNrOrHash)
	if err != nil {
		return nil, err
	}
	if block == nil {
		return nil, fmt.Errorf("block not found")
	}
	if block.NumberU64() == 0 {
		genesis, err := getGenesisState(block.Hash())
		if err != nil {
			return nil, fmt.Errorf("failed to get genesis state: %w", err)
		}
		header := util.BuildPilelineBlockHeader(block)
		// Convert GenesisAlloc to ptypes.GenesisAlloc
		palloc := make(ptypes.GenesisAlloc, len(genesis))
		for addr, acc := range genesis {
			palloc[addr] = ptypes.GenesisAccount{
				Balance: acc.Balance,
				Code:    acc.Code,
				Storage: acc.Storage,
				Nonce:   acc.Nonce,
			}
		}
		blockDiff := ptracer.GenesisAllocToStateDiff(palloc)
		blockDiff.Hash = header.StateRoot
		blockFile := &ptypes.BlockFile{
			Block:            util.BuildPipelineBlock(block),
			Txs:              make([]ptypes.Transaction, 0),
			Events:           make([]ptypes.Event, 0),
			Traces:           make([]ptypes.Trace, 0),
			ErrorEvents:      make([]ptypes.Event, 0),
			ErrorTraces:      make([]ptypes.Trace, 0),
			StorageContracts: make([]string, 0),
		}
		for addr, account := range genesis {
			if len(account.Storage) > 0 {
				blockFile.StorageContracts = append(blockFile.StorageContracts, strings.ToLower(addr.Hex()))
			}
		}
		var stateDiffBytes []byte
		if blockDiff != nil {
			stateDiffBytes, err = util.EncodeToRlp(blockDiff)
			if err != nil {
				log.Error("Failed to encode state diff", "err", err)
				stateDiffBytes = []byte{}
			}
		} else {
			stateDiffBytes = []byte{}
		}

		return &ptypes.DebankOutPut{
			BlockFile:      blockFile,
			Header:         header,
			StateDiff:      hexutil.Bytes(stateDiffBytes),
			ValidationHash: blockFile.Validation().ValidationHash,
		}, nil
	}

	// Prepare base state
	parent := api.eth.blockchain.GetBlock(block.ParentHash(), block.NumberU64()-1)
	if parent == nil {
		return nil, fmt.Errorf("parent block not found")
	}
	statedb, err := api.eth.blockchain.StateAt(parent.Root())
	if err != nil {
		return nil, err
	}
	XDCxState := &tradingstate.TradingStateDB{}
	parentState := statedb.Copy()

	rpcTracer := ptracer.RPCTracer{}
	vmConfig := vm.Config{
		Tracer: &rpcTracer,
	}

	rpcTracer.OnBlockStart(block)

	chainConfig := api.eth.blockchain.Config()

	// Mutate the block and state according to any hard-fork specs
	if chainConfig.DAOForkSupport && chainConfig.DAOForkBlock != nil && chainConfig.DAOForkBlock.Cmp(block.Number()) == 0 {
		misc.ApplyDAOHardFork(statedb)
	}

	var (
		txs     = block.Transactions()
		header  = block.Header()
		gp      = new(core.GasPool).AddGas(block.GasLimit())
		usedGas = new(uint64)
	)

	// Set hooks for log tracing
	statedb.OnLog = rpcTracer.OnLog

	// Recompute transactions up to the target index.
	feeCapacity := state.GetTRC21FeeCapacityFromState(statedb)
	if common.TIPSigning.Cmp(block.Header().Number) == 0 {
		statedb.DeleteAddress(common.BlockSignersBinary)
	}

	var receipts = make(types.Receipts, 0)
	for i, tx := range txs {
		statedb.SetTxContext(tx.Hash(), i)
		if core.IsSkipEvmTransaction(api.eth.chainConfig, block.Number(), tx) {
			rpcTracer.OnSkipEvmTxStart(tx, *tx.From())
		} else {
			rpcTracer.OnTxStart(tx, *tx.From())
		}

		// Apply the transaction
		receipt, gas, err, tokenFeeUsed := core.ApplyTransaction(api.eth.chainConfig, feeCapacity, api.eth.blockchain, nil, gp, statedb, XDCxState, block.Header(), tx, usedGas, vmConfig)
		if err != nil {
			return nil, fmt.Errorf("could not apply tx %d [%v]: %w", i, tx.Hash().Hex(), err)
		}
		if tokenFeeUsed {
			fee := common.GetGasFee(block.Header().Number.Uint64(), gas)
			feeCapacity[*tx.To()] = new(big.Int).Sub(feeCapacity[*tx.To()], fee)
		}

		receipt.SetEffectiveGasPrice(tx, header.BaseFee)
		receipts = append(receipts, receipt)
		rpcTracer.OnTxEnd(receipt, nil)
	}

	// Finalize the block, applying any consensus engine specific extras (e.g. block rewards)
	api.eth.engine.Finalize(api.eth.blockchain, header, statedb, parentState, block.Transactions(), block.Uncles(), receipts)

	root, destructs, accounts, storages, codes, err := statedb.StateDiff(chainConfig.IsEIP158(block.Number()))
	if err != nil {
		return nil, fmt.Errorf("could not get state diff: %w", err)
	}

	if root != block.Header().Root {
		return nil, fmt.Errorf("state root mismatch: expected %x, got %x", block.Header().Root, root)
	}

	parentRoot := parent.Root()

	res := rpcTracer.GetOutPut(parentRoot, root, destructs, accounts, storages, codes)

	return res, nil
}

func getGenesisState(blockhash common.Hash) (alloc types.GenesisAlloc, err error) {
	// Genesis allocation is missing and there are several possibilities:
	// the node is legacy which doesn't persist the genesis allocation or
	// the persisted allocation is just lost.
	// - supported networks(mainnet, testnets), recover with defined allocations
	// - private network, can't recover
	var genesis *core.Genesis
	switch blockhash {
	case params.MainnetGenesisHash:
		genesis = core.DefaultGenesisBlock()
	case params.TestnetGenesisHash:
		genesis = core.DefaultTestnetGenesisBlock()
	}
	if genesis != nil {
		return genesis.Alloc, nil
	}

	return nil, nil
}
