package ethtypes

import (
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestAddress(t *testing.T) {
	rawAddr := common.HexToAddress("0x11")
	rawAddrJson, err := json.Marshal(rawAddr)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("rawAddrJson: %v", string(rawAddrJson))

	addr := Address(rawAddr)
	addrJson, err := json.Marshal(addr)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("addrJson: %v", string(addrJson))

	addr2 := Address{}
	err = json.Unmarshal(addrJson, &addr2)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("addr2: %v", addr2)
}
