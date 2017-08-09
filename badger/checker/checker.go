package main

import (
	"io/ioutil"
	"os"
	"reflect"
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
	stdout := records(strings.Split(string(buf), "\n"))

	//checkAtomicUpdateConstency(stdout)
	checkBadgerConsistency(stdout)
}

type records []string

func (r records) contains(s string) bool {
	for _, rec := range r {
		if rec == s {
			return true
		}
	}
	return false
}

func checkBadgerConsistency(stdout records) {
	kv := StartBadger()
	defer func() { Must(kv.Close()) }()

	if stdout.contains("k1=v1") {
		Assert(Exists(kv, k1))
	}
	if Exists(kv, k1) {
		item := MustGet(kv, k1)
		Assert(reflect.DeepEqual(item.Value(), v1))
	}
}

func checkAtomicUpdateConstency(stdout records) {
	buf, err := ioutil.ReadFile("file1")
	Must(err)
	str := strings.TrimSpace(string(buf))
	isHello := str == "hello"
	isWorld := str == "world"
	Assert(isHello || isWorld)
}
