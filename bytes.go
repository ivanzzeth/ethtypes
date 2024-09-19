package ethtypes

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm/schema"
)

var bytesTestVal = Bytes([]byte{})
var _ schema.SerializerInterface = &bytesTestVal

type Bytes []byte

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

func (b Bytes) String() string {
	return b.Hex()
}

func (b Bytes) Unwrap() []byte {
	return []byte(b)
}
