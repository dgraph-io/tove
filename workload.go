package main

import (
	"log"

	"github.com/dgraph-io/badger"
)

var (
	k1 = []byte("k1")
	v1 = []byte("value1")
)

func main() {

	opt := &badger.DefaultOptions
	opt.Dir = "workload_dir"
	opt.ValueDir = opt.Dir

	kv, err := badger.NewKV(opt)
	must(err)
	defer func() { must(kv.Close()) }()

	must(kv.Set(k1, v1, 0))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
