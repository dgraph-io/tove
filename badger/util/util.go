package util

import (
	"encoding/binary"
	"log"
	"reflect"
	"github.com/dgraph-io/badger"
)

func StartBadger() *badger.DB {
	opt := badger.DefaultOptions
	opt.MaxTableSize = 2 << 10
	opt.NumLevelZeroTables = 1
	opt.NumLevelZeroTablesStall = 2
	opt.NumMemtables = 1
	opt.LevelOneSize = 2 << 20
	opt.ValueLogFileSize = 5 << 20
	opt.SyncWrites = true
	// Allow truncation because ALICE would add garbage data to the value log.
	opt.Truncate = true

	opt.Dir = "."
	opt.ValueDir = "."
	db, err := badger.Open(opt)
	Must(err)
	return db
}

func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func MustBool(b bool, err error) bool {
	Must(err)
	return b
}

func Exists(db *badger.DB, k []byte) bool {
	err := db.View(func(txn *badger.Txn) error {
		_, err := txn.Get(k)
		return err
	})

	if err != nil && err != badger.ErrKeyNotFound {
		log.Fatal(err)
	}

	return err == nil
}

func MustGet(db *badger.DB, k []byte) []byte {
	var value []byte
	Must(db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(k)
		if err != nil {
			return err
		}

		value, err = item.ValueCopy(value)
		if err != nil {
			return err
		}

		return nil
	}))

	return value
}

func MustSet(db *badger.DB, k []byte, v []byte) {
	Must(db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(k), []byte(v))
	}))
}

func MustDelete(db *badger.DB, k []byte) {
	Must(db.Update(func(txn *badger.Txn) error {
		return txn.Delete(k)
	}))
}

func Assert(b bool) {
	if !b {
		log.Fatal("Assertion failed")
	}
}

func KeyHasValue(db *badger.DB, k, v []byte) bool {
	Assert(Exists(db, k))
	value := MustGet(db, k)
	return reflect.DeepEqual(v, value)
}

var (
	BigK = make([]byte, 64)
	BigV = make([]byte, 1<<20)
)

const (
	KeyCount = 3
	Versions = 2
)

var (
	// Use a hand rolled RNG to avoid rand related system calls.
	x   uint32
	rnd = func() uint32 {
		x = uint32(1664525*uint64(x) + 1013904223)
		return x
	}
)

func init() {
	for i := range BigK {
		BigK[i] = byte(rnd())
	}
	for i := range BigV {
		BigV[i] = byte(rnd())
	}
}

func ConstructValue(i, j uint16) []byte {
	binary.BigEndian.PutUint16(BigV, i)
	binary.BigEndian.PutUint16(BigV[2:], j)
	return BigV
}

func ConstructKey(i uint16) []byte {
	binary.BigEndian.PutUint16(BigK, i)
	return BigK
}
