package ethtypes

import (
	"context"
	"encoding"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/gorm/schema"
)

var hashTestVal = Hash(common.Hash{})
var _ schema.SerializerInterface = &hashTestVal
var _ fmt.Stringer = hashTestVal
var _ json.Marshaler = hashTestVal
var _ json.Unmarshaler = &hashTestVal
var _ encoding.TextMarshaler = hashTestVal
var _ encoding.TextUnmarshaler = &hashTestVal

type Hash common.Hash

func NewHash(b common.Hash) *Hash {
	bi := Hash(b)
	return &bi
}

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

func (hash Hash) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, hash.String())), nil
}

func (hash *Hash) UnmarshalJSON(data []byte) error {
	dataStr := strings.ReplaceAll(string(data), "\"", "")
	*hash = Hash(common.HexToHash(dataStr))

	return nil
}

func (bi Hash) MarshalText() (text []byte, err error) {
	return bi.MarshalJSON()
}

func (bi *Hash) UnmarshalText(text []byte) error {
	return bi.UnmarshalJSON(text)
}

func (hash Hash) String() string {
	return strings.ToLower(common.Hash(hash).Hex())
}

func (hash Hash) Unwrap() common.Hash {
	return common.Hash(hash)
}
