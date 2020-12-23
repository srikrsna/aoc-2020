package common

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func ScanInput(fn func(scanner *bufio.Scanner) error) error {
	return Scan("input.txt", fn)
}

func Scan(file string, fn func(scanner *bufio.Scanner) error) error {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("input err: %w", err)
	}
	defer f.Close()

	return fn(bufio.NewScanner(f))
}

func Spew(v ...interface{}) {
	f, err := os.Create("debug.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	spew.Config.DisablePointerAddresses = true
	spew.Config.SpewKeys = true
	spew.Config.Indent = "\t"

	spew.Fdump(f, v...)
}
