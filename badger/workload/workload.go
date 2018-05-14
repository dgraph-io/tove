package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	. "github.com/dgraph-io/tove/badger/util"
)

func main() {

	Must(os.Chdir("workload_dir"))
	//atomicUpdateWorkload()
	//badgerWorkload()
	badgerBigWorkload()
}

func badgerBigWorkload() {
	fmt.Println("start:big")
	db := StartBadger()
	for j := 0; j < Versions; j++ {
		for i := 0; i < KeyCount; i++ {
			key := ConstructKey(uint16(i))
			val := ConstructValue(uint16(i), uint16(j))
			MustSet(db, key, val)
		}
	}
	time.Sleep(100 * time.Millisecond)
	Must(db.Close())
	fmt.Println("stop:big")
}

func badgerWorkload() {

	var (
		k1 = []byte("k1")
		v1 = []byte("value1")
		v2 = []byte("value2")
	)

	fmt.Println("start:set-key")
	db := StartBadger()
	MustSet(db, k1, v1)
	Must(db.Close())
	fmt.Println("stop:set-key")

	fmt.Println("start:update-key")
	db = StartBadger()
	MustSet(db, k1, v2)
	Must(db.Close())
	fmt.Println("stop:update-key")

	fmt.Println("start:del-key")
	db = StartBadger()
	MustDelete(db, k1)
	Must(db.Close())
	fmt.Println("stop:del-key")

	fmt.Println("start:ins-upd-del")
	db = StartBadger()
	MustSet(db, k1, v1)
	MustSet(db, k1, v2)
	MustDelete(db, k1)
	Must(db.Close())
	fmt.Println("stop:ins-upd-del")
}

func atomicUpdateWorkload() {
	const tmp = "tmp"
	Must(ioutil.WriteFile(tmp, []byte("world"), 0666))
	Must(os.Rename(tmp, "file1"))
}
