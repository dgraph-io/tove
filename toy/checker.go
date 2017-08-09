package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	must(os.Chdir(os.Args[1]))

	buf, err := ioutil.ReadFile("file1")
	must(err)

	str := strings.TrimSpace(string(buf))
	isHello := str == "hello"
	isWorld := str == "world"

	if !isHello && !isWorld {
		must(errors.New("not hello or world"))
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
