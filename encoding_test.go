package ethtypes

import (
	"encoding"
	"encoding/json"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/magiconair/properties/assert"
)

func TestTextEncoding(t *testing.T) {
	testcases := []struct {
		encoder textEncoder
	}{
		{
			encoder: NewAddress(common.HexToAddress("0x1111")),
		},
		{
			encoder: NewBigFloat(big.NewFloat(123.456)),
		},
		{
			encoder: NewBigInt(big.NewInt(123)),
		},
		{
			encoder: NewBytes(common.Hex2Bytes("0x1111")),
		},
		{
			encoder: NewHash(common.HexToHash("0x1111")),
		},
	}

	for _, tt := range testcases {
		testTextEncoder(t, tt.encoder)
	}
}

func TestJsonEncoding(t *testing.T) {
	testcases := []struct {
		encoder jsonEncoder
	}{
		{
			encoder: NewAddress(common.HexToAddress("0x1111")),
		},
		{
			encoder: NewBigFloat(big.NewFloat(123.456)),
		},
		{
			encoder: NewBigInt(big.NewInt(123)),
		},
		{
			encoder: NewBytes(common.Hex2Bytes("0x1111")),
		},
		{
			encoder: NewHash(common.HexToHash("0x1111")),
		},
	}

	for _, tt := range testcases {
		testJsonEncoder(t, tt.encoder)
	}
}

type textEncoder interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
}

type jsonEncoder interface {
	json.Marshaler
	json.Unmarshaler
}

func testTextEncoder(t *testing.T, encoder textEncoder) {
	text, err := encoder.MarshalText()
	if err != nil {
		t.Fatal(err)
	}

	newEncoder := reflect.New(reflect.TypeOf(encoder).Elem()).Interface().(textEncoder)
	err = newEncoder.UnmarshalText(text)
	if err != nil {
		t.Fatal(err)
	}

	newText, err := newEncoder.MarshalText()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, text, newText)
}

func testJsonEncoder(t *testing.T, encoder jsonEncoder) {
	text, err := encoder.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	newEncoder := reflect.New(reflect.TypeOf(encoder).Elem()).Interface().(jsonEncoder)
	err = newEncoder.UnmarshalJSON(text)
	if err != nil {
		t.Fatal(err)
	}

	newText, err := newEncoder.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, text, newText)
}
