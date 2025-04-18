package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

type Counters struct {
	Bytes uint64
	Chars uint64
	Words uint64
	Lines uint64
}

type Options struct {
	PrintBytes bool
	PrintChars bool
	PrintWords bool
	PrintLines bool
}

func main() {
	cPtr := flag.Bool("c", false, "list # of bytes")
	lPtr := flag.Bool("l", false, "list # of lines")
	wPtr := flag.Bool("w", false, "list # of words")
	mPtr := flag.Bool("m", false, "list # of characters")

	flag.Parse()

	if !(*cPtr || *lPtr || *wPtr || *mPtr) {
		*cPtr = true
		*lPtr = true
		*wPtr = true
	}

	options := Options{
		PrintBytes: *cPtr, 
		PrintChars: *mPtr,
		PrintWords: *wPtr, 
		PrintLines: *lPtr,
	}

	args := flag.Args()

	var err error

	if len(args) == 0 {
		err = handleStdIn(&options)
	} else {
		err = handleFiles(args, &options)
	}

	if err != nil {
		log.Fatalln(err)
	}
}

func handleStdIn(options *Options) error {
	counters, err := count(os.Stdin)

	if err != nil {
		return fmt.Errorf("Failed to count stdin: %w", err)
	}

	print("", counters, options)

	return nil
}

func handleFiles(files []string, options *Options) error {
	totalCounters := Counters{}

	for _, v := range files {
		f, err := os.Open(v)

		if err != nil {
			return fmt.Errorf("Failed to open file: %w", err)
		}

		defer f.Close()

		var r io.Reader = f

		counters, err := count(r)

		if err != nil {
			return fmt.Errorf("Failed to count file: %w", err)
		}

		totalCounters.Lines += counters.Lines
		totalCounters.Words += counters.Words
		totalCounters.Chars += counters.Chars
		totalCounters.Bytes += counters.Bytes

		print(v, counters, options)
	}

	if len(files) > 1 {
		print("total", &totalCounters, options)
	}

	return nil
}

func count(r io.Reader) (*Counters, error) {
	br := bufio.NewReader(r)

	inWord := false

	counters := Counters{}

	for {
		char, size, err := br.ReadRune()

		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, fmt.Errorf("Failed to read character from file: %w", err)
		}

		counters.Bytes += uint64(size)
		counters.Chars++

		if char == '\n' {
			counters.Lines++
		}

		if unicode.IsSpace(char) {
			inWord = false			
		} else if !inWord {
			counters.Words++
			inWord = true
		}
	}

	return &counters, nil
}

func print(desc string, counters *Counters, options *Options) {
	fmt.Print("\t")

	if options.PrintLines {
		fmt.Printf("%d\t", counters.Lines)
	}

	if options.PrintWords {
		fmt.Printf("%d\t", counters.Words)
	}

	if options.PrintChars {
		fmt.Printf("%d\t", counters.Chars)
	}

	if options.PrintBytes {
		fmt.Printf("%d\t", counters.Bytes)
	}

	fmt.Print(desc + "\n")
}