package ethtypes

import (
	"encoding/json"
	"math/big"
	"testing"
)

func TestBigInt(t *testing.T) {
	rawBigInt := big.NewInt(12345)
	rawBigIntJson, err := json.Marshal(rawBigInt)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("rawBigIntJson: %v", string(rawBigIntJson))

	bigInt := BigInt(*big.NewInt(12345))
	bigIntJson, err := json.Marshal(bigInt)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("bigIntJson: %v", string(bigIntJson))

	decodedBigInt := BigInt{}
	err = json.Unmarshal(bigIntJson, &decodedBigInt)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("decoded bigInt: %v", decodedBigInt.String())
}
