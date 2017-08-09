package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/dgraph-io/badger"
)

var (
	k1 = []byte("k1")
	v1 = []byte("value1")
)

func main() {

	must(os.Chdir("workload_dir"))
	atomicUpdateWorkload()
	badgerWorkload()
}

func badgerWorkload() {
	opt := &badger.DefaultOptions
	opt.Dir = "."
	opt.ValueDir = "."
	kv, err := badger.NewKV(opt)
	must(err)
	defer func() { must(kv.Close()) }()

	must(kv.Set(k1, v1, 0))
}

func atomicUpdateWorkload() {
	const tmp = "tmp"
	must(ioutil.WriteFile(tmp, []byte("world"), 0666))
	must(os.Rename(tmp, "file1"))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
