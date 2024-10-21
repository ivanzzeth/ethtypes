package ethtypes

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"encoding"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm/schema"
)

var addrTestVal = Address(common.Address{})
var _ schema.SerializerInterface = &addrTestVal
var _ fmt.Stringer = addrTestVal
var _ json.Marshaler = addrTestVal
var _ json.Unmarshaler = &addrTestVal
var _ encoding.TextMarshaler = addrTestVal
var _ encoding.TextUnmarshaler = &addrTestVal

type Address common.Address

func (addr *Address) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	switch value := dbValue.(type) {
	case string:
		if !common.IsHexAddress(value) {
			return fmt.Errorf("invalid address: %v", value)
		}
		*addr = Address(common.HexToAddress(value))
	default:
		return fmt.Errorf("unsupported data %#v", dbValue)
	}
	return nil
}

func (addr Address) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	return strings.ToLower(common.Address(addr).Hex()), nil
}

func (addr Address) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, addr.String())), nil
}

func (addr *Address) UnmarshalJSON(data []byte) error {
	dataStr := strings.ReplaceAll(string(data), "\"", "")
	if !common.IsHexAddress(dataStr) {
		return fmt.Errorf("invalid address: %v", dataStr)
	}
	*addr = Address(common.HexToAddress(dataStr))

	return nil
}

func (addr Address) MarshalText() (text []byte, err error) {
	return addr.MarshalJSON()
}

func (addr *Address) UnmarshalText(text []byte) error {
	return addr.UnmarshalJSON(text)
}

func (addr Address) String() string {
	return strings.ToLower(common.Address(addr).Hex())
}

func (addr Address) Unwrap() common.Address {
	return common.Address(addr)
}

type AddressArray []Address

func (arr AddressArray) String() string {
	res := ""
	for _, v := range arr {
		res += v.String()
	}

	return res
}
