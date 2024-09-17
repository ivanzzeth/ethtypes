package ethtypes

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm/schema"
)

var addrTestVal = Address(common.Address{})
var _ schema.SerializerInterface = &addrTestVal

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

func (addr Address) String() string {
	return strings.ToLower(common.Address(addr).Hex())
}

func (addr Address) Unwrap() common.Address {
	return common.Address(addr)
}
