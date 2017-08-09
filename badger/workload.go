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

	opt := &badger.DefaultOptions
	opt.Dir = "."
	opt.ValueDir = "."

	kv, err := badger.NewKV(opt)
	must(err)

	must(kv.Set(k1, v1, 0))

	must(kv.Close())

	badAtomicUpdate()
}

func badAtomicUpdate() {
	const tmp = "tmp"
	must(ioutil.WriteFile(tmp, []byte("world"), 0666))
	must(os.Rename(tmp, "file1"))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
