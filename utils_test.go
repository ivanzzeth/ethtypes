package ethtypes

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/magiconair/properties/assert"
)

func TestUtils(t *testing.T) {
	assert.Equal(t, ToString("1"), "1")
	assert.Equal(t, ToString([]string{"1"}), `["1"]`)
	assert.Equal(t, ToString([]common.Address{
		common.HexToAddress("0x2222"),
		common.HexToAddress("0x3333"),
	}), `["0x0000000000000000000000000000000000002222","0x0000000000000000000000000000000000003333"]`)

	assert.Equal(t, ToString([]*big.Int{
		big.NewInt(2),
		big.NewInt(3),
	}), `["2","3"]`)
}
