package ethtypes

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestTypes(t *testing.T) {
	dsn := "host=172.16.80.124 user=postgres password=gavreqp51.sfg1 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	type TestEvent struct {
		gorm.Model
		Bytes    Bytes
		Account  Address
		Accounts string
		Amount   BigInt
		Amounts  string
	}

	tables := []interface{}{
		&TestEvent{},
	}

	db.Migrator().DropTable(tables...)
	db.AutoMigrate(tables...)

	b := common.Hex2Bytes("123456")

	t.Log("bytes: ", b)
	err = db.Save(&TestEvent{
		Bytes:   b,
		Account: Address(common.HexToAddress("0x1111")),
		Accounts: ToString([]common.Address{
			common.HexToAddress("0x2222"),
			common.HexToAddress("0x3333"),
		}),
		Amount: BigInt(*big.NewInt(1)),
		Amounts: ToString([]*big.Int{
			big.NewInt(2),
			big.NewInt(3),
		}),
	}).Error

	if err != nil {
		t.Fatal(err)
	}
}
