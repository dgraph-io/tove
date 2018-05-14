package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	. "github.com/dgraph-io/tove/badger/util"
)

func main() {

	Must(os.Chdir(os.Args[1]))
	buf, err := ioutil.ReadFile(os.Args[2])
	Must(err)
	stdout := strings.Split(string(buf), "\n")

	//checkAtomicUpdateConstency(stdout)
	//checkBadgerConsistency(stdout)
	checkBadgerBigWorkloadConsistency(stdout)
}

func checkBadgerBigWorkloadConsistency(stdout []string) {
	db := StartBadger()
	defer func() { Must(db.Close()) }()

	if len(stdout) == 0 {
		return
	}

	var lastMsg string
	for i := len(stdout) - 1; i >= 0; i-- {
		if strings.HasPrefix(stdout[i], "start:") ||
			strings.HasPrefix(stdout[i], "stop:") {
			lastMsg = stdout[i]
			break
		}
	}

	switch lastMsg {
	case "start:big":
		for i := 0; i < KeyCount; i++ {
			k := ConstructKey(uint16(i))
			if Exists(db, k) {
				hasValidValue := false
				for j := 0; j < Versions; j++ {
					value := MustGet(db, k)
					v := ConstructValue(uint16(i), uint16(j))
					if reflect.DeepEqual(value, v) {
						hasValidValue = true
						break
					}
				}
				Assert(hasValidValue)
			}
		}
	case "stop:big":
		for i := 0; i < KeyCount; i++ {
			const j = Versions - 1
			Assert(KeyHasValue(
				db,
				ConstructKey(uint16(i)),
				ConstructValue(uint16(i), uint16(j)),
			))
		}
	}
}

func checkBadgerConsistency(stdout []string) {

	var (
		k1 = []byte("k1")
		v1 = []byte("value1")
		v2 = []byte("value2")
	)

	db := StartBadger()
	defer func() { Must(db.Close()) }()

	if len(stdout) == 0 {
		return
	}

	var lastMsg string
	for i := len(stdout) - 1; i >= 0; i-- {
		if strings.HasPrefix(stdout[i], "start:") ||
			strings.HasPrefix(stdout[i], "stop:") {
			lastMsg = stdout[i]
			break
		}
	}

	switch lastMsg {
	case "start:set-key":
		if Exists(db, k1) {
			Assert(KeyHasValue(db, k1, v1))
		}
	case "stop:set-key":
		Assert(KeyHasValue(db, k1, v1))
	case "start:update-key":
		Assert(Exists(db, k1))
		Assert(KeyHasValue(db, k1, v1) || KeyHasValue(db, k1, v2))
	case "stop:update-key":
		Assert(KeyHasValue(db, k1, v2))
	case "start:del-key":
		if Exists(db, k1) {
			Assert(KeyHasValue(db, k1, v2))
		}
	case "stop:del-key":
		Assert(!Exists(db, k1))
	case "start:ins-upd-del":
		if Exists(db, k1) {
			Assert(KeyHasValue(db, k1, v1) || KeyHasValue(db, k1, v2))
		}
	case "stop:ins-upd-del":
		Assert(!Exists(db, k1))
	default:
		Assert(false)
	}
}

func checkAtomicUpdateConstency(stdout []string) {
	buf, err := ioutil.ReadFile("file1")
	Must(err)
	str := strings.TrimSpace(string(buf))
	isHello := str == "hello"
	isWorld := str == "world"
	Assert(isHello || isWorld)
}
