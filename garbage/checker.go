package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
)

var value []byte

func init() {
	value = make([]byte, 1<<20)
	for i := range value {
		value[i] = byte(1664525*i + 1013904223)
	}
}

func main() {

	must(os.Chdir(os.Args[1]))

	_, err := os.Stat("file")
	if err != nil {
		return
	}

	buf, err := ioutil.ReadFile("file")
	must(err)

	if !bytes.HasPrefix(value, buf) {
		log.Fatal("didn't have prefix")
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
