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
	kv := StartBadger()
	Must(kv.Set(k1, v1, 0))
	Must(kv.Close())
	fmt.Println("k1=v1")
}

func atomicUpdateWorkload() {
	const tmp = "tmp"
	Must(ioutil.WriteFile(tmp, []byte("world"), 0666))
	Must(os.Rename(tmp, "file1"))
}
