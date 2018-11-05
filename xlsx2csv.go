package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "hoge")
	}
	flag.Parse()
}

func main() {
	xlsx2csv()
	os.Exit(0)
}

func xlsx2csv() {
}
