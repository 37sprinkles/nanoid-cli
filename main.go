package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jaevor/go-nanoid"
)

func printUsage() {
	fmt.Println("Usage: nanoid charset length [count]")
	fmt.Println("  charset (string): required, \"hex\", \"alpha\", \"numeric\", \"base64\", \"base90\" or a string of characters")
	fmt.Println("  length (byte): required, must be between 2 and 255")
	fmt.Println("  count (int32): optional, must be at least 1")
}

func parseFlags() (string, int, int, bool) {
	if len(flag.Args()) < 2 || len(flag.Args()) > 3 {
		return "", 0, 0, false
	}

	charset := flag.Arg(0)
	switch strings.ToLower(charset) {
	case "hex":
		charset = "0123456789ABCDEF"
	case "alpha":
		charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case "numeric":
		charset = "0123456789"
	}

	length, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		return "", 0, 0, false
	}

	if length <= 1 || length > 255 {
		return "", 0, 0, false
	}

	count := 1
	if len(flag.Args()) == 3 {
		count, err = strconv.Atoi(flag.Arg(2))

		if err != nil {
			return "", 0, 0, false
		}
	}

	if count <= 0 {
		return "", 0, 0, false
	}

	return charset, length, count, true
}

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		printUsage()
		os.Exit(0)
	}

	charset, length, count, ok := parseFlags()

	if !ok {
		printUsage()
		os.Exit(1)
	}

	var generator func() string
	var err error

	if charset == "base64" {
		generator, err = nanoid.Standard(length)
	} else if charset == "base90" {
		generator, err = nanoid.ASCII(length)
	} else {
		generator, err = nanoid.CustomASCII(charset, length)
	}

	if err != nil {
		panic(err)
	}

	for range count {
		fmt.Println(generator())
	}
}
