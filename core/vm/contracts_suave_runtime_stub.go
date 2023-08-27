// Code generated by suave/gen. DO NOT EDIT.
// Hash: 0479cfd5872468c477531906bed5fc154c66974e33d2e7b38e91f526ec8d4182
package vm

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/suave/artifacts"
	"github.com/mitchellh/mapstructure"
)

type BackendImpl interface {
	buildEthBlock(blockArgs types.BuildBlockArgs, bid [16]byte, namespace string) ([]byte, []byte, error)
	confidentialStoreRetrieve(bidId [16]byte, key string) ([]byte, error)
	confidentialStoreStore(bidId [16]byte, key string, data []byte) error
	extractHint(bundleData []byte) ([]byte, error)
	fetchBids(cond uint64, namespace string) ([]types.Bid, error)
	newBid(decryptionCondition uint64, allowedPeekers []common.Address, BidType string) (types.Bid, error)
	simulateBundle(bundleData []byte) (uint64, error)
	submitEthBlockBidToRelay(relayUrl string, builderBid []byte) ([]byte, error)
}

type BackendStub struct {
	impl BackendImpl
}

func (b *BackendStub) buildEthBlock(input []byte) ([]byte, error) {
	var err error

	unpacked, err := artifacts.SuaveAbi.Methods["buildEthBlock"].Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	var blockArgs types.BuildBlockArgs
	if err := mapstructure.Decode(unpacked[0], &blockArgs); err != nil {
		return nil, err
	}

	var (
		res0 []byte
		res1 []byte
	)

	if res0, res1, err = b.impl.buildEthBlock(blockArgs, unpacked[1].([16]byte), unpacked[2].(string)); err != nil {
		return nil, err
	}

	packedRes, err := artifacts.SuaveAbi.Methods["buildEthBlock"].Outputs.Pack(res0, res1)
	if err != nil {
		return nil, err
	}
	return packedRes, nil
}

func (b *BackendStub) confidentialStoreRetrieve(input []byte) ([]byte, error) {
	var err error

	unpacked, err := artifacts.SuaveAbi.Methods["confidentialStoreRetrieve"].Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	var (
		res0 []byte
	)

	if res0, err = b.impl.confidentialStoreRetrieve(unpacked[0].([16]byte), unpacked[1].(string)); err != nil {
		return nil, err
	}

	return res0, nil
}

func (b *BackendStub) confidentialStoreStore(input []byte) ([]byte, error) {
	var err error

	unpacked, err := artifacts.SuaveAbi.Methods["confidentialStoreStore"].Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	if err = b.impl.confidentialStoreStore(unpacked[0].([16]byte), unpacked[1].(string), unpacked[2].([]byte)); err != nil {
		return nil, err
	}

	return nil, nil
}

func (b *BackendStub) extractHint(input []byte) ([]byte, error) {
	var err error

	unpacked, err := artifacts.SuaveAbi.Methods["extractHint"].Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	var (
		res0 []byte
	)

	if res0, err = b.impl.extractHint(unpacked[0].([]byte)); err != nil {
		return nil, err
	}

	return res0, nil
}

func (b *BackendStub) fetchBids(input []byte) ([]byte, error) {
	var err error

	unpacked, err := artifacts.SuaveAbi.Methods["fetchBids"].Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	var (
		res0 []types.Bid
	)

	if res0, err = b.impl.fetchBids(unpacked[0].(uint64), unpacked[1].(string)); err != nil {
		return nil, err
	}

	packedRes, err := artifacts.SuaveAbi.Methods["fetchBids"].Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}
	return packedRes, nil
}

func (b *BackendStub) newBid(input []byte) ([]byte, error) {
	var err error

	unpacked, err := artifacts.SuaveAbi.Methods["newBid"].Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	var (
		res0 types.Bid
	)

	if res0, err = b.impl.newBid(unpacked[0].(uint64), unpacked[1].([]common.Address), unpacked[2].(string)); err != nil {
		return nil, err
	}

	packedRes, err := artifacts.SuaveAbi.Methods["newBid"].Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}
	return packedRes, nil
}

func (b *BackendStub) simulateBundle(input []byte) ([]byte, error) {
	var err error

	var (
		res0 uint64
	)

	if res0, err = b.impl.simulateBundle(input); err != nil {
		return nil, err
	}

	packedRes, err := artifacts.SuaveAbi.Methods["simulateBundle"].Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}
	return packedRes, nil
}

func (b *BackendStub) submitEthBlockBidToRelay(input []byte) ([]byte, error) {
	var err error

	unpacked, err := artifacts.SuaveAbi.Methods["submitEthBlockBidToRelay"].Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	var (
		res0 []byte
	)

	if res0, err = b.impl.submitEthBlockBidToRelay(unpacked[0].(string), unpacked[1].([]byte)); err != nil {
		return nil, err
	}

	return res0, nil
}