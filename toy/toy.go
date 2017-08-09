package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	const file1 = "file1"
	const tmp = "tmp"
	must(ioutil.WriteFile(tmp, []byte("world"), 0666))
	must(os.Rename(tmp, file1))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
