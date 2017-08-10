package main

import (
	"log"
	"os"

	"golang.org/x/sys/unix"
)

var value []byte

func init() {
	value = make([]byte, 1<<20)
	for i := range value {
		value[i] = byte(1664525*i + 1013904223)
	}
}

func main() {
	flags := os.O_RDWR | os.O_CREATE | os.O_EXCL | unix.O_DSYNC
	f, err := os.OpenFile("file", flags, 0644)
	must(err)
	dirfd, err := os.Open(".")
	must(err)
	must(dirfd.Sync())
	must(dirfd.Close())
	_, err = f.Write(value)
	must(err)
	must(f.Close())
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
