package main

import (
	"fmt"
	"io/ioutil"
	"os"

	. "github.com/dgraph-io/tove/badger/util"
)

var (
	k1 = []byte("k1")
	v1 = []byte("value1")
)

func main() {

	Must(os.Chdir("workload_dir"))
	//atomicUpdateWorkload()
	badgerWorkload()
}

func badgerWorkload() {
	fmt.Println("start:set-key")
	kv := StartBadger()
	Must(kv.Set(k1, v1, 0))
	Must(kv.Close())
	fmt.Println("stop:set-key")

	//kv = StartBadger()
	//Must(kv.Delete(k1))
	//Must(kv.Close())
	//fmt.Println("del k1")
}

func atomicUpdateWorkload() {
	const tmp = "tmp"
	Must(ioutil.WriteFile(tmp, []byte("world"), 0666))
	Must(os.Rename(tmp, "file1"))
}
