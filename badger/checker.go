package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/dgraph-io/badger"
)

var (
	k1 = []byte("k1")
	v1 = []byte("value1")
)

func main() {

	must(os.Chdir(os.Args[1]))

	opt := &badger.DefaultOptions
	opt.Dir = "."
	opt.ValueDir = "."

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

	checkAtomicUpdate()
}

func checkAtomicUpdate() {
	buf, err := ioutil.ReadFile("file1")
	must(err)
	str := strings.TrimSpace(string(buf))
	isHello := str == "hello"
	isWorld := str == "world"
	if !isHello && !isWorld {
		must(errors.New("not hello or world"))
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
