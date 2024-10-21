package ethtypes

import (
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/shopspring/decimal"
)

func TestBigDecimal(t *testing.T) {
	d := NewBigDecimal(decimal.NewFromFloat(12.4))
	res := d.Unwrap().Div(decimal.NewFromInt(4))

	assert.Equal(t, "3.1", res.String())
	assert.Equal(t, "12.4", d.Unwrap().String())
}
