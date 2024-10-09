package ethtypes

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestHash(t *testing.T) {
	hash := Hash(common.HexToHash("0x1234"))
	js, err := hash.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("js: %v", string(js))

	decodedHash := Hash{}
	err = decodedHash.UnmarshalJSON(js)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("decoded: %v", decodedHash)
}
