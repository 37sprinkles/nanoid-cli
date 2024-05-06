package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/jaevor/go-nanoid"
)

func parseFlags() (string, int) {
	var (
		charset string
		hex     bool
		alpha   bool
		numeric bool
		base64  bool
		base90  bool
	)

	flag.StringVar(&charset, "c", "", "Custom character set")
	flag.BoolVar(&hex, "hex", false, "Uses a hexadecimal character set")
	flag.BoolVar(&alpha, "alpha", false, "Uses an alphabetic character set")
	flag.BoolVar(&numeric, "numeric", false, "Uses a numeric character set")
	flag.BoolVar(&base64, "base64", false, "Uses a base64 character set")
	flag.BoolVar(&base90, "base90", false, "Uses a base90 character set (includes symbols)")

	flag.Usage = func() {
		fmt.Printf("Usage: nanoid [options] length\nOptions:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.Parse()

	if len(flag.Args()) != 1 {
		panic(errors.New("invalid args: missing length"))
	}

	length, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		panic(errors.New("invalid args: length must be an integer"))
	}

	if length == 0 {
		panic(errors.New("invalid args: length cannot be 0"))
	}

	if length <= 0 || length > 255 {
		panic(errors.New("invalid args: length must be greater that 0 and less that 256"))
	}

	if hex {
		if charset != "" {
			panic(errors.New("invalid args: multiple character sets specified"))
		}

		charset = "0123456789ABCDEF"
	}

	if alpha {
		if charset != "" {
			panic(errors.New("invalid args: multiple character sets specified"))
		}

		charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if numeric {
		if charset != "" {
			panic(errors.New("invalid args: multiple character sets specified"))
		}

		charset = "0123456789"
	}

	if base64 {
		if charset != "" {
			panic(errors.New("invalid args: multiple character sets specified"))
		}

		charset = "base64"
	}

	if base90 {
		if charset != "" {
			panic(errors.New("invalid args: multiple character sets specified"))
		}

		charset = "base90"
	}

	if charset == "" {
		panic(errors.New("invalid args: missing character set"))
	}

	return charset, length
}

func main() {
	charset, length := parseFlags()

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

	fmt.Println(generator())
}
