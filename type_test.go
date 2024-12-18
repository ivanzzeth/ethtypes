package ethtypes

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestTypes(t *testing.T) {
	dsn := "host=192.168.31.83 user=postgres password=gavreqp51.sfg1 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	// make sure it works with real db.
	testTypes(t, db)

	liteDb, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	// make sure it works with sqlite
	testTypes(t, liteDb)
}

func testTypes(t *testing.T, db *gorm.DB) {
	t.Logf("test db %v", db.Dialector.Name())

	var err error
	type TestEvent struct {
		gorm.Model
		Bytes    Bytes
		Account  Address
		Accounts string
		Amount   BigInt
		Amounts  string
		Float    BigFloat
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
		Float: BigFloat(*big.NewFloat(2.3)),
	}).Error

	if err != nil {
		t.Fatal(err)
	}

	newTestEvent := &TestEvent{}
	err = db.First(newTestEvent).Error
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("newTestEvent: %v", newTestEvent)
}
