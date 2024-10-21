package ethtypes

import (
	"context"
	"encoding"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var bigIntTestVal = BigInt(big.Int{})
var _ schema.SerializerInterface = &bigIntTestVal
var _ fmt.Stringer = bigIntTestVal
var _ json.Marshaler = bigIntTestVal
var _ json.Unmarshaler = &bigIntTestVal
var _ encoding.TextMarshaler = bigIntTestVal
var _ encoding.TextUnmarshaler = &bigIntTestVal

type BigInt big.Int

func (bi BigInt) GormDataType() string {
	return "numeric"
}

func (bi BigInt) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	// use field.Tag, field.TagSettings gets field's tags
	// checkout https://github.com/go-gorm/gorm/blob/master/schema/field.go for all options

	// returns different database type based on driver name
	switch db.Dialector.Name() {
	case "postgres":
		return "numeric"
	}

	return "text"
}

func (bi *BigInt) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	switch value := dbValue.(type) {
	case string:
		bigInt, ok := big.NewInt(0).SetString(value, 10)
		if !ok {
			return fmt.Errorf("invalid *big.Int: %v", value)
		}
		*bi = BigInt(*bigInt)
	default:
		return fmt.Errorf("unsupported data %#v", dbValue)
	}
	return nil
}

func (bi BigInt) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	return bi.Unwrap().String(), nil
}

func (bi BigInt) MarshalJSON() ([]byte, error) {
	return bi.Unwrap().MarshalJSON()
}

func (bi *BigInt) UnmarshalJSON(data []byte) error {
	rawBi := bi.Unwrap()
	err := rawBi.UnmarshalJSON(data)
	if err != nil {
		return err
	}

	bi.Set(rawBi)
	return nil
}

func (bi BigInt) MarshalText() (text []byte, err error) {
	return bi.MarshalJSON()
}

func (bi *BigInt) UnmarshalText(text []byte) error {
	return bi.UnmarshalJSON(text)
}

func (bi *BigInt) Set(i *big.Int) {
	(*bi) = BigInt(*i)
}

func (bi BigInt) String() string {
	return bi.Unwrap().String()
}

func (bi *BigInt) Unwrap() *big.Int {
	res := (*big.Int)(bi)
	return res
}
