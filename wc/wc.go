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

	args := flag.Args()

	fmt.Println(args)

	if len(args) == 0 {

		return
	}

	var totalLines, totalWords, totalChars, totalBytes uint64

	for _, v := range args {
		f, err := os.Open(v)

		if err != nil {
			log.Fatalln(err)
		}

		var r io.Reader = f

		lines, words, chars, bytes, err := count(r)

		totalLines += lines
		totalWords += words
		totalChars += chars
		totalBytes += bytes

		fmt.Print("\t")

		if *lPtr {
			fmt.Printf("%d ", lines)
		}

		if *wPtr {
			fmt.Printf("%d ", words)
		}

		if *mPtr {
			fmt.Printf("%d ", chars)
		}

		if *cPtr {
			fmt.Printf("%d ", bytes)
		}

		fmt.Print(v + "\n")

		f.Close()
	}

	if len(args) > 1 {
		fmt.Print("\t")

		if *lPtr {
			fmt.Printf("%d ", totalLines)
		}

		if *wPtr {
			fmt.Printf("%d ", totalWords)
		}

		if *mPtr {
			fmt.Printf("%d ", totalChars)
		}

		if *cPtr {
			fmt.Printf("%d ", totalBytes)
		}

		fmt.Print("total")
	}
}

func count(r io.Reader) (lines, words, chars, bytes uint64, err error) {
	br := bufio.NewReader(r)

	inWord := false

	for {
		char, size, err := br.ReadRune()

		if err != nil {
			if err == io.EOF {
				break
			}

			return 0, 0, 0, 0, fmt.Errorf("Failed to read character from file: %w", err)
		}

		bytes += uint64(size)
		chars++

		if char == '\n' {
			lines++
		}

		if unicode.IsSpace(char) {
			inWord = false			
		} else if !inWord {
			words++
			inWord = true
		}
	}

	return
}
