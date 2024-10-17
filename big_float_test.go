package ethtypes

import (
	"encoding/json"
	"math/big"
	"testing"
)

func TestBigFloat(t *testing.T) {
	rawBigFloat := big.NewFloat(123.456)
	rawBigFloatJson, err := json.Marshal(rawBigFloat)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("rawBigFloatJson: %v", string(rawBigFloatJson))

	bigFloat := BigFloat(*big.NewFloat(123.456))
	bigFloatJson, err := json.Marshal(bigFloat)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("bigFloatJson: %v", string(bigFloatJson))

	decodedBigFloat := BigFloat{}
	err = json.Unmarshal(bigFloatJson, &decodedBigFloat)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("decoded bigFloat: %v", decodedBigFloat.String())
}
