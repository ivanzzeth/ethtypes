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

var bigFloatTestVal = BigFloat(big.Float{})
var _ schema.SerializerInterface = &bigFloatTestVal
var _ fmt.Stringer = bigFloatTestVal
var _ json.Marshaler = bigFloatTestVal
var _ json.Unmarshaler = &bigFloatTestVal
var _ encoding.TextMarshaler = bigFloatTestVal
var _ encoding.TextUnmarshaler = &bigFloatTestVal

type BigFloat big.Float

func NewBigFloat(b *big.Float) *BigFloat {
	bi := BigFloat(*b)
	return &bi
}

func (bi BigFloat) GormDataType() string {
	return "numeric"
}

func (bi BigFloat) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	// use field.Tag, field.TagSettings gets field's tags
	// checkout https://github.com/go-gorm/gorm/blob/master/schema/field.go for all options

	// returns different database type based on driver name
	switch db.Dialector.Name() {
	case "postgres":
		return "numeric"
	}

	return "text"
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

func (bi BigFloat) MarshalText() (text []byte, err error) {
	return bi.MarshalJSON()
}

func (bi *BigFloat) UnmarshalText(text []byte) error {
	return bi.UnmarshalJSON(text)
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
