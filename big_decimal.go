package ethtypes

import (
	"context"
	"encoding"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func init() {
	decimal.DivisionPrecision = 256
	decimal.PowPrecisionNegativeExponent = 256
}

var bigDecimalTestVal = BigDecimal(decimal.New(1, 0))

var _ schema.SerializerInterface = &bigDecimalTestVal
var _ fmt.Stringer = bigDecimalTestVal
var _ json.Marshaler = bigDecimalTestVal
var _ json.Unmarshaler = &bigDecimalTestVal
var _ encoding.TextMarshaler = bigDecimalTestVal
var _ encoding.TextUnmarshaler = &bigDecimalTestVal

type BigDecimal decimal.Decimal

func NewBigDecimal(decimal decimal.Decimal) *BigDecimal {
	bi := BigDecimal(decimal)
	return &bi
}

func (bi BigDecimal) GormDataType() string {
	return "numeric"
}

func (bi BigDecimal) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	// use field.Tag, field.TagSettings gets field's tags
	// checkout https://github.com/go-gorm/gorm/blob/master/schema/field.go for all options

	// returns different database type based on driver name
	switch db.Dialector.Name() {
	case "postgres":
		return "numeric"
	}

	return "text"
}

func (bi *BigDecimal) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	switch value := dbValue.(type) {
	case string:
		rawBi, err := decimal.NewFromString(value)
		if err != nil {
			return fmt.Errorf("invalid float string %v: %v", value, err)
		}
		*bi = BigDecimal(rawBi)
	default:
		return fmt.Errorf("unsupported data %#v", dbValue)
	}
	return nil
}

func (bi BigDecimal) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	return bi.Unwrap().String(), nil
}

func (bi BigDecimal) MarshalJSON() ([]byte, error) {
	str := bi.Unwrap().String()
	return json.Marshal(str)
}

func (bi *BigDecimal) UnmarshalJSON(data []byte) error {
	str := ""
	err := json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	rawBi, err := decimal.NewFromString(str)
	if err != nil {
		return fmt.Errorf("invalid float string %v: %v", str, err)
	}

	bi.Set(rawBi)
	return nil
}

func (bi BigDecimal) MarshalText() (text []byte, err error) {
	return bi.MarshalJSON()
}

func (bi *BigDecimal) UnmarshalText(text []byte) error {
	return bi.UnmarshalJSON(text)
}

func (bi BigDecimal) String() string {
	return bi.Unwrap().String()
}

func (bi *BigDecimal) Set(decimal decimal.Decimal) {
	*bi = BigDecimal(decimal)
}

func (bi BigDecimal) Unwrap() decimal.Decimal {
	return decimal.Decimal(bi)
}
