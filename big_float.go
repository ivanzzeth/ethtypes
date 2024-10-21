package ethtypes

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"

	"gorm.io/gorm/schema"
)

var bigFloatTestVal = BigFloat(big.Float{})
var _ schema.SerializerInterface = &bigFloatTestVal
var _ fmt.Stringer = bigFloatTestVal
var _ json.Marshaler = bigFloatTestVal
var _ json.Unmarshaler = &bigFloatTestVal

type BigFloat big.Float

func (bi BigFloat) GormDataType() string {
	return "numeric"
}

func (bi *BigFloat) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	switch value := dbValue.(type) {
	case string:
		bigFloat, ok := big.NewFloat(0).SetString(value)
		if !ok {
			return fmt.Errorf("invalid *big.Float: %v", value)
		}
		*bi = BigFloat(*bigFloat)
	default:
		return fmt.Errorf("unsupported data %#v", dbValue)
	}
	return nil
}

func (bi BigFloat) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	return bi.Unwrap().String(), nil
}

func (bi BigFloat) MarshalJSON() ([]byte, error) {
	str := bi.Unwrap().String()
	return json.Marshal(str)
}

func (bi *BigFloat) UnmarshalJSON(data []byte) error {
	str := ""
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	rawBi, ok := big.NewFloat(0).SetString(str)
	if !ok {
		return fmt.Errorf("invalid float string: %v", str)
	}

	bi.Set(rawBi)
	return nil
}

func (bi *BigFloat) Set(i *big.Float) {
	(*bi) = BigFloat(*i)
}

func (bi BigFloat) String() string {
	return bi.Unwrap().String()
}

func (bi *BigFloat) Unwrap() *big.Float {
	res := (*big.Float)(bi)
	return res
}
