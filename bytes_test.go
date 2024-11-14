package ethtypes

import (
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/magiconair/properties/assert"
)

func TestBytesJson(t *testing.T) {
	b1 := Bytes([]byte{1, 2, 3, 4})
	js, err := b1.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("js: %v", string(js))

	decodedB1 := Bytes{}
	err = decodedB1.UnmarshalJSON(js)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("decoded: %v", decodedB1)

}

func TestBytes(t *testing.T) {
	b1 := Bytes([]byte{1, 2, 3, 4})
	b1Str := b1.String()
	v, _ := hexutil.Decode(b1Str)
	b2 := Bytes(v)
	b2Str := b2.String()

	assert.Equal(t, b2Str, b1Str)
}
