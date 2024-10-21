package ethtypes

import (
	"context"
	"encoding"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gorm.io/gorm/schema"
)

var bytesTestVal = Bytes([]byte{})
var _ schema.SerializerInterface = &bytesTestVal
var _ fmt.Stringer = bytesTestVal
var _ json.Marshaler = bytesTestVal
var _ json.Unmarshaler = &bytesTestVal
var _ encoding.TextMarshaler = bytesTestVal
var _ encoding.TextUnmarshaler = &bytesTestVal

type Bytes []byte

func NewBytes(b []byte) *Bytes {
	bi := Bytes(b)
	return &bi
}

func (b *Bytes) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	switch value := dbValue.(type) {
	case string:
		// TODO: Validation
		if !strings.HasPrefix(value, "0x") {
			value = "0x" + value
		}
		*b = Bytes(common.Hex2Bytes(value))
	default:
		return fmt.Errorf("unsupported data %#v", dbValue)
	}
	return nil
}

func (b Bytes) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	return b.Hex(), nil
}

func (b Bytes) Hex() string {
	return fmt.Sprintf("0x%x", []byte(b))
}

func (b Bytes) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, b.String())), nil
}

func (b *Bytes) UnmarshalJSON(data []byte) error {
	dataStr := strings.ReplaceAll(string(data), "\"", "")
	byts, err := hexutil.Decode(dataStr)
	if err != nil {
		return err
	}

	b.Set(byts)
	return nil
}

func (bi Bytes) MarshalText() (text []byte, err error) {
	return bi.MarshalJSON()
}

func (bi *Bytes) UnmarshalText(text []byte) error {
	return bi.UnmarshalJSON(text)
}

func (b *Bytes) Set(data []byte) {
	*b = Bytes(data)
}

func (b Bytes) String() string {
	return b.Hex()
}

func (b Bytes) Unwrap() []byte {
	return []byte(b)
}
