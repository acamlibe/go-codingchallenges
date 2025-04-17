package main

import (
	"flag"
	"fmt"
)

func main() {
	cPtr := flag.Bool("c", false, "list # of bytes")
	lPtr := flag.Bool("l", false, "list # of lines")
	wPtr := flag.Bool("w", false, "list # of words")
	mPtr := flag.Bool("m", false, "list # of characters")

	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		
	}
}