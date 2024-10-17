package ethtypes

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestSlice(t *testing.T) {
	t.Logf("testAddress======")
	addr1 := common.HexToAddress("0x11")
	addr2 := common.HexToAddress("0x22")
	addrs := []*Address{
		NewAddress(addr1),
		NewAddress(addr2),
	}
	testSlice(t, addrs)

	t.Logf("testBigInt======")
	bigInts := []*BigInt{
		NewBigInt(big.NewInt(1)),
		NewBigInt(big.NewInt(2)),
		NewBigInt(big.NewInt(3)),
	}
	testSlice(t, bigInts)

	t.Logf("testBytes======")
	bytes := []*Bytes{
		NewBytes(common.Hex2Bytes("1111")),
		NewBytes(common.Hex2Bytes("2222")),
		NewBytes(common.Hex2Bytes("3333")),
	}
	testSlice(t, bytes)

	t.Logf("testHash======")
	hashes := []*Hash{
		NewHash(common.HexToHash("0x1111")),
		NewHash(common.HexToHash("0x2222")),
		NewHash(common.HexToHash("0x3333")),
	}
	testSlice(t, hashes)
}

func testSlice[E JsonObj](t *testing.T, data []E) {
	slice := NewGormSlice(data)
	js, err := slice.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("slice js: %v", string(js))

	newSlice := NewGormSlice[E](nil)
	err = newSlice.UnmarshalJSON(js)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("newSlice: %v", newSlice.Unwrap())
}
