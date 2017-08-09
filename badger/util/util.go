package util

import (
	"log"

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

func MustGet(kv *badger.KV, k []byte) badger.KVItem {
	var item badger.KVItem
	Must(kv.Get(k, &item))
	return item
}

func Assert(b bool) {
	if !b {
		log.Fatal("Assertion failed")
	}
}
