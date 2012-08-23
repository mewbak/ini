package main

// This compares parsing performance of my original ini package
// with this one.
//
// The old one is faster with very large input files (>= 100kb)
// For the rest, the new one is faster. Since ini files will likely
// never reach that large a size, the new package is preferable.

import (
	"fmt"
	a "github.com/jteeuwen/gbbs/ini"
	b "github.com/jteeuwen/go-pkg-ini"
	"time"
)

const Limit = 100000
const File = "test.ini"

func main() {
	s := time.Now()

	for i := 0; i < Limit; i++ {
		testA()
	}

	fmt.Printf("%s\n", time.Now().Sub(s))

	s = time.Now()

	for i := 0; i < Limit; i++ {
		testB()
	}

	fmt.Printf("%s\n", time.Now().Sub(s))
}

func testA() {
	file := a.NewFile()
	err := file.Load(File)

	if err != nil {
		fmt.Printf("testB: %v\n", err)
	}
}

func testB() {
	cfg, err := b.Load(File)
	_ = cfg

	if err != nil {
		fmt.Printf("testB: %v\n", err)
	}
}
