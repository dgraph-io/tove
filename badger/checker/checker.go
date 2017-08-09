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
	stdout := lines(strings.Split(string(buf), "\n"))

	//checkAtomicUpdateConstency(stdout)
	checkBadgerConsistency(stdout)
}

type lines []string

func (l lines) contains(s string) bool {
	for _, line := range l {
		if line == s {
			return true
		}
	}
	return false
}

func (l lines) finished() bool {
	for i := len(l) - 1; i >= 0; i-- {
		if strings.HasPrefix(l[i], "start") {
			return false
		}
	}
	return true
}

func (l lines) stage() string {
	for i := len(l) - 1; i >= 0; i-- {
		str := l[i]
		if strings.HasPrefix(str, "start:") {
			return strings.TrimPrefix(str, "start:")
		}
	}
	return ""
}

func checkBadgerConsistency(stdout lines) {
	kv := StartBadger()
	defer func() { Must(kv.Close()) }()

	if len(stdout) == 0 {
		return
	}

	switch stdout[len(stdout)-1] {
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
	default:
	}
}

func checkAtomicUpdateConstency(stdout lines) {
	buf, err := ioutil.ReadFile("file1")
	Must(err)
	str := strings.TrimSpace(string(buf))
	isHello := str == "hello"
	isWorld := str == "world"
	Assert(isHello || isWorld)
}
