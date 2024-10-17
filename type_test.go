package ethtypes

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func TestTypes(t *testing.T) {
	db := GetTestDb()

	var err error

	type TestEvent struct {
		gorm.Model
		Bytes      Bytes
		Account    Address
		Accounts   string
		Amount     *BigInt
		Amounts    string
		Values     *GormSlice[*BigInt]
		JsonValues datatypes.JSONSlice[BigInt]
	}

	tables := []interface{}{
		&TestEvent{},
	}

	db.Migrator().DropTable(tables...)
	db.AutoMigrate(tables...)

	b := common.Hex2Bytes("123456")

	t.Log("bytes: ", b)

	slice := NewGormSlice([]*BigInt{
		NewBigInt(big.NewInt(5)),
		NewBigInt(big.NewInt(6)),
		NewBigInt(big.NewInt(7)),
	})
	err = db.Save(&TestEvent{
		Bytes:   b,
		Account: Address(common.HexToAddress("0x1111")),
		Accounts: ToString([]common.Address{
			common.HexToAddress("0x2222"),
			common.HexToAddress("0x3333"),
		}),
		Amount: NewBigInt(big.NewInt(1)),
		Amounts: ToString([]*big.Int{
			big.NewInt(2),
			big.NewInt(3),
		}),
		Values: slice,
		JsonValues: datatypes.NewJSONSlice([]BigInt{
			BigInt(*big.NewInt(8)),
			BigInt(*big.NewInt(9)),
		}),
	}).Error
	if err != nil {
		t.Fatal(err)
	}

	t.Log("11111")
	newTestEvent := &TestEvent{}
	err = db.First(newTestEvent).Error
	if err != nil {
		t.Fatal(err)
	}

	t.Log("22222")

	t.Logf("newTestEvent: %v", newTestEvent)
}
