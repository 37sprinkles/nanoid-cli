package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jaevor/go-nanoid"
)

func parseFlags() (string, int, int) {
	printUsage := func() {
		fmt.Println("Usage: nanoid charset length [count]")
		fmt.Println("  charset (string): required, \"hex\", \"alpha\", \"numeric\", \"base64\", \"base90\" or a string of characters")
		fmt.Println("  length (byte): required, must be between 2 and 255")
		fmt.Println("  count (int32): optional, must be at least 1")
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.Parse()

	if len(flag.Args()) < 2 || len(flag.Args()) > 3 {
		printUsage()
	}

	charset := strings.ToLower(flag.Arg(0))
	switch charset {
	case "hex":
		charset = "0123456789ABCDEF"
	case "alpha":
		charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case "numeric":
		charset = "0123456789"
	}

	length, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		printUsage()
	}

	if length <= 1 || length > 255 {
		printUsage()
	}

	count := 1
	if len(flag.Args()) == 3 {
		count, err = strconv.Atoi(flag.Arg(2))

		if err != nil {
			printUsage()
		}
	}

	if count <= 0 {
		printUsage()
	}

	return charset, length, count
}

func main() {
	charset, length, count := parseFlags()

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
