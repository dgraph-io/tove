package util

import (
	"log"
	"reflect"

	"github.com/dgraph-io/badger"
)

func StartBadger() *badger.KV {
	opt := &badger.DefaultOptions
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

func AssertKeyValue(kv *badger.KV, k, v []byte) {
	Assert(Exists(kv, k))
	item := MustGet(kv, k)
	Assert(reflect.DeepEqual(v, item.Value()))
	Assert(reflect.DeepEqual(k, item.Key()))
}
