package eth

import (
	"context"
	"crypto/md5"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"testing"

	ptypes "github.com/Chaintable/pipeline/types"
	"github.com/Chaintable/pipeline/util"
	"github.com/XinFinOrg/XDPoSChain/common"
	"github.com/XinFinOrg/XDPoSChain/consensus/ethash"
	"github.com/XinFinOrg/XDPoSChain/core"
	"github.com/XinFinOrg/XDPoSChain/core/rawdb"
	"github.com/XinFinOrg/XDPoSChain/core/types"
	"github.com/XinFinOrg/XDPoSChain/core/vm"
	"github.com/XinFinOrg/XDPoSChain/params"
	"github.com/XinFinOrg/XDPoSChain/rpc"
)

func genesisDebankAPI(t *testing.T, genesis *core.Genesis) (*DebankAPI, *types.Block) {
	t.Helper()
	db := rawdb.NewMemoryDatabase()
	block := genesis.MustCommit(db)
	chain, err := core.NewBlockChain(db, nil, genesis.Config, ethash.NewFaker(), vm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { chain.Stop(); db.Close() })
	eth := &Ethereum{blockchain: chain, chainConfig: genesis.Config}
	eth.ApiBackend = &EthAPIBackend{eth: eth}
	return NewDebankAPI(eth), block
}

func TestDebankMainnetGenesis(t *testing.T) {
	genesis := core.DefaultGenesisBlock()
	api, block := genesisDebankAPI(t, genesis)
	if block.Hash() != params.MainnetGenesisHash {
		t.Fatalf("unexpected mainnet genesis: %s", block.Hash())
	}
	var previous *ptypes.BlockFile
	for _, query := range []rpc.BlockNumberOrHash{
		rpc.BlockNumberOrHashWithNumber(0),
		rpc.BlockNumberOrHashWithHash(block.Hash(), true),
	} {
		out, err := api.DebankBlock(context.Background(), query)
		if err != nil {
			t.Fatal(err)
		}
		bf := out.BlockFile
		if len(bf.Txs) != 8 || len(bf.Traces) != 8 || len(bf.StorageContracts) != 4 {
			t.Fatalf("incomplete genesis: txs=%d traces=%d storage=%d", len(bf.Txs), len(bf.Traces), len(bf.StorageContracts))
		}
		if out.ValidationHash != 675247 || bf.Validation().ValidationHash != out.ValidationHash {
			t.Fatalf("unexpected validation: %+v", bf.Validation())
		}
		kinds := map[string]int{}
		for i, tx := range bf.Txs {
			if len(tx.ID) != 66 || strings.HasPrefix(tx.ID, "0xgenesis") {
				t.Fatalf("invalid tx ID: %q", tx.ID)
			}
			kind := tx.ID[2:4]
			kinds[kind]++
			if tx.ID != "0x"+kind+strings.Repeat("0", 22)+tx.To[2:] || tx.TransactionIndex != int64(i) {
				t.Fatalf("invalid tx identity/order: %+v", tx)
			}
			trace := bf.Traces[i]
			if trace.TxID != tx.ID || trace.ID != fmt.Sprintf("%x", md5.Sum([]byte(tx.ID+"0"))) {
				t.Fatalf("invalid trace binding at %d", i)
			}
			account := genesis.Alloc[common.HexToAddress(tx.To)]
			switch kind {
			case "01":
				if account.Balance == nil || (*big.Int)(tx.Value).Cmp(account.Balance) != 0 || trace.CallCreateType != "call" {
					t.Fatalf("invalid alloc balance for %s", tx.To)
				}
			case "02":
				if len(account.Code) == 0 || !reflect.DeepEqual([]byte(tx.Input), account.Code) || !reflect.DeepEqual([]byte(trace.Output), account.Code) || trace.CallCreateType != "create" {
					t.Fatalf("invalid alloc code for %s", tx.To)
				}
			case "03":
				if tx.To != "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee" || len(tx.Input) != 0 || (*big.Int)(tx.Value).Sign() != 0 {
					t.Fatalf("invalid native placeholder: %+v", tx)
				}
			default:
				t.Fatalf("unexpected genesis kind: %s", kind)
			}
		}
		if !reflect.DeepEqual(kinds, map[string]int{"01": 2, "02": 5, "03": 1}) {
			t.Fatalf("unexpected genesis records: %v", kinds)
		}
		if out.Header.Hash != block.Hash() || out.Header.StateRoot != block.Root() {
			t.Fatal("genesis header changed")
		}
		var diff ptypes.BlockStorageDiff
		if err := util.DecodeFromRlp(out.StateDiff, &diff); err != nil {
			t.Fatal(err)
		}
		if diff.Hash != block.Root() || diff.ParentHash != types.EmptyRootHash || len(diff.NewAccounts) != 8 || len(diff.NewCodes) != 5 {
			t.Fatalf("unexpected genesis state diff: %+v", diff)
		}
		if previous != nil && (!reflect.DeepEqual(previous.Txs, bf.Txs) || !reflect.DeepEqual(previous.Traces, bf.Traces) || !reflect.DeepEqual(previous.StorageContracts, bf.StorageContracts)) {
			t.Fatal("genesis records depend on lookup or map iteration order")
		}
		previous = bf
	}
}

func TestDebankUnknownGenesis(t *testing.T) {
	api, _ := genesisDebankAPI(t, &core.Genesis{Config: params.TestChainConfig, Alloc: types.GenesisAlloc{common.Address{1}: {Balance: big.NewInt(1)}}})
	if _, err := api.DebankBlock(context.Background(), rpc.BlockNumberOrHashWithNumber(0)); err == nil {
		t.Fatal("unsupported genesis must not return an empty successful result")
	}
}
