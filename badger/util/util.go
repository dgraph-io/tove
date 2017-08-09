package util

import (
	"encoding/binary"
	"log"
	"reflect"
	"time"

	"github.com/dgraph-io/badger"
)

func StartBadger() *badger.KV {
	opt := &badger.DefaultOptions
	opt.ValueGCRunInterval = time.Second
	opt.MaxTableSize = 512
	opt.NumLevelZeroTables = 1
	opt.NumLevelZeroTablesStall = 2
	opt.NumMemtables = 1
	opt.LevelOneSize = 2 << 20
	opt.ValueLogFileSize = 5 << 20
	opt.Dir = "."
	opt.ValueDir = "."
	kv, err := badger.NewKV(opt)
	Must(err)
	return kv
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

func Exists(kv *badger.KV, k []byte) bool {
	ok, err := kv.Exists(k)
	Must(err)
	return ok
}

func MustGet(kv *badger.KV, k []byte) *badger.KVItem {
	item := new(badger.KVItem)
	Must(kv.Get(k, item))
	return item
}

func Assert(b bool) {
	if !b {
		log.Fatal("Assertion failed")
	}
}

func KeyHasValue(kv *badger.KV, k, v []byte) bool {
	Assert(Exists(kv, k))
	item := MustGet(kv, k)
	Assert(reflect.DeepEqual(k, item.Key()))
	return reflect.DeepEqual(v, item.Value())
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
