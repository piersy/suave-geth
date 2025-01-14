package vm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	suave "github.com/ethereum/go-ethereum/suave/core"
)

type SuaveExecutionBackend struct {
	ConfidentialStoreBackend suave.ConfidentialStoreBackend
	MempoolBackend           suave.MempoolBackend
	OffchainEthBackend       suave.OffchainEthBackend
	confidentialInputs       []byte
	callerStack              []*common.Address
}

func NewRuntimeSuaveExecutionBackend(evm *EVM, caller common.Address) *SuaveExecutionBackend {
	if !evm.Config.IsOffchain {
		return nil
	}

	return &SuaveExecutionBackend{
		ConfidentialStoreBackend: evm.suaveExecutionBackend.ConfidentialStoreBackend,
		MempoolBackend:           evm.suaveExecutionBackend.MempoolBackend,
		OffchainEthBackend:       evm.suaveExecutionBackend.OffchainEthBackend,
		confidentialInputs:       evm.suaveExecutionBackend.confidentialInputs,
		callerStack:              append(evm.suaveExecutionBackend.callerStack, &caller),
	}
}

// Implements PrecompiledContract for Offchain smart contracts
type SuavePrecompiledContractWrapper struct {
	addr     common.Address
	backend  *SuaveExecutionBackend
	contract SuavePrecompiledContract
}

func NewSuavePrecompiledContractWrapper(addr common.Address, backend *SuaveExecutionBackend, contract SuavePrecompiledContract) *SuavePrecompiledContractWrapper {
	return &SuavePrecompiledContractWrapper{addr: addr, backend: backend, contract: contract}
}

func (p *SuavePrecompiledContractWrapper) RequiredGas(input []byte) uint64 {
	return p.contract.RequiredGas(input)
}

func (p *SuavePrecompiledContractWrapper) Run(input []byte) ([]byte, error) {
	stub := &BackendStub{
		impl: &backendImpl{
			backend: p.backend,
		},
	}

	switch p.addr {
	case isOffchainAddress:
		return (&isOffchainPrecompile{}).RunOffchain(p.backend, input)

	case confidentialInputsAddress:
		return (&confidentialInputsPrecompile{}).RunOffchain(p.backend, input)

	case confStoreStoreAddress:
		return stub.confidentialStoreStore(input)

	case confStoreRetrieveAddress:
		return stub.confidentialStoreRetrieve(input)

	case newBidAddress:
		return stub.newBid(input)

	case fetchBidsAddress:
		return stub.fetchBids(input)

	case extractHintAddress:
		return stub.extractHint(input)

	case simulateBundleAddress:
		return stub.simulateBundle(input)

	case buildEthBlockAddress:
		return stub.buildEthBlock(input)

	case submitEthBlockBidToRelayAddress:
		return stub.submitEthBlockBidToRelay(input)
	}

	return nil, fmt.Errorf("precompile %s not found", p.addr)
}
