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
	v2 = []byte("value2")
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

	fmt.Println("start:update-key")
	kv = StartBadger()
	Must(kv.Set(k1, v2, 0))
	Must(kv.Close())
	fmt.Println("stop:update-key")

	fmt.Println("start:del-key")
	kv = StartBadger()
	Must(kv.Delete(k1))
	Must(kv.Close())
	fmt.Println("stop:del-key")

	//fmt.Println("start:ins-upd-del")
	//kv = StartBadger()
	//Must(kv.Set(k1, v1))
	//Must(kv.Set(k1, v2))
	//Must(kv.Delete(k1))
	//Must(kv.Close())
	//fmt.Println("stop:ins-upd-del")
}

func atomicUpdateWorkload() {
	const tmp = "tmp"
	Must(ioutil.WriteFile(tmp, []byte("world"), 0666))
	Must(os.Rename(tmp, "file1"))
}
