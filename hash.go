package ethtypes

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm/schema"
)

var hashTestVal = Hash(common.Hash{})
var _ schema.SerializerInterface = &hashTestVal

type Hash common.Hash

func (hash *Hash) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	switch value := dbValue.(type) {
	case string:
		// TODO: validation
		*hash = Hash(common.HexToHash(value))
	default:
		return fmt.Errorf("unsupported data %#v", dbValue)
	}
	return nil
}

func (hash Hash) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	return strings.ToLower(common.Hash(hash).Hex()), nil
}

func (hash Hash) String() string {
	return strings.ToLower(common.Hash(hash).Hex())
}

func (hash Hash) Unwrap() common.Hash {
	return common.Hash(hash)
}
