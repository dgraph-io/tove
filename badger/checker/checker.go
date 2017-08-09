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

	switch stdout.stage() {
	case "": // Didn't start the first stage yet
	case "set-key":
		if stdout.finished() {
			AssertKeyValue(kv, k1, v1)
		} else {
			if Exists(kv, k1) {
				AssertKeyValue(kv, k1, v1)
			}
		}
	case "del-key":
		if stdout.finished() {
			Assert(!Exists(kv, k1))
		} else {
			if Exists(kv, k1) {
				AssertKeyValue(kv, k1, v1)
			}
		}
	default:
		Assert(false)
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
