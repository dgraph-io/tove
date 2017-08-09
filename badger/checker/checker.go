package main

import (
	"io/ioutil"
	"os"
	"strings"

	. "github.com/dgraph-io/tove/badger/util"
)

var (
	k1 = []byte("k1")
	v1 = []byte("value1")
	v2 = []byte("value2")
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
	kv := StartBadger()
	defer func() { Must(kv.Close()) }()

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
	case "stop:big":
		// FIXME: for now, don't do any checks. Just make sure Badger starts
		// up with no errors.
		/*
			for i := 0; i < KeyCount; i++ {
				const j = Versions - 1
				Assert(KeyHasValue(
					kv,
					ConstructKey(uint16(i)),
					ConstructValue(uint16(i), uint16(j)),
				))
			}
		*/
	default:
	}
}

func checkBadgerConsistency(stdout []string) {
	kv := StartBadger()
	defer func() { Must(kv.Close()) }()

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
		if Exists(kv, k1) {
			Assert(KeyHasValue(kv, k1, v1))
		}
	case "stop:set-key":
		Assert(KeyHasValue(kv, k1, v1))
	case "start:update-key":
		Assert(Exists(kv, k1))
		Assert(KeyHasValue(kv, k1, v1) || KeyHasValue(kv, k1, v2))
	case "stop:update-key":
		Assert(KeyHasValue(kv, k1, v2))
	case "start:del-key":
		if Exists(kv, k1) {
			Assert(KeyHasValue(kv, k1, v2))
		}
	case "stop:del-key":
		Assert(!Exists(kv, k1))
	case "start:ins-upd-del":
		if Exists(kv, k1) {
			Assert(KeyHasValue(kv, k1, v1) || KeyHasValue(kv, k1, v2))
		}
	case "stop:ins-upd-del":
		Assert(!Exists(kv, k1))
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
