package main

import (
	"log"
	"os"
	"reflect"

	"github.com/dgraph-io/badger"
)

var (
	k1 = []byte("k1")
	v1 = []byte("value1")
)

func main() {

	if len(os.Args) < 3 {
		log.Fatal("bad number of args")
	}

	opt := &badger.DefaultOptions
	opt.Dir = os.Args[1]
	opt.ValueDir = opt.Dir

	kv, err := badger.NewKV(opt)
	must(err)
	defer func() { must(kv.Close()) }()

	ok, err := kv.Exists(k1)
	must(err)
	if ok {
		var item badger.KVItem
		must(kv.Get(k1, &item))
		if !reflect.DeepEqual(item.Value(), v1) {
			log.Fatal("value not set")
		}
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
