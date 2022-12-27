package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

// CipherAction ...
type CipherAction int

const (
	// Encode ...
	Encode CipherAction = iota
	// Decode ...
	Decode
)

var (
	encodeFlag = false
	decodeFlag = false
)

// cipher takes in the text to be ciphered along with the direction that
// is being taken; -1 means encoding, +1 means decoding.
// https://go.dev/play/p/0rWUGaMLjn
func cipher(text string, direction CipherAction) string {
	// shift -> number of letters to move to right or left
	// offset -> size of the alphabet, in this case the plain ASCII
	shift, offset := rune(3), rune(26)

	// string->rune conversion
	runes := []rune(text)

	for index, char := range runes {
		// Iterate over all runes, and perform substitution
		// wherever possible. If the letter is not in the range
		// [1 .. 25], the offset defined above is added or
		// subtracted.
		switch direction {
		case Encode:
			if char >= 'a'+shift && char <= 'z' ||
				char >= 'A'+shift && char <= 'Z' {
				char = char - shift
			} else if char >= 'a' && char < 'a'+shift ||
				char >= 'A' && char < 'A'+shift {
				char = char - shift + offset
			}
		case Decode:
			if char >= 'a' && char <= 'z'-shift ||
				char >= 'A' && char <= 'Z'-shift {
				char = char + shift
			} else if char > 'z'-shift && char <= 'z' ||
				char > 'Z'-shift && char <= 'Z' {
				char = char + shift - offset
			}
		}

		// Above `if`s handle both upper and lower case ASCII
		// characters; anything else is returned as is (includes
		// numbers, punctuation and space).
		runes[index] = char
	}

	return string(runes)
}

// encode and decode provide the API for encoding and decoding text using
// the Caesar Cipher algorithm.
func encode(text string) string {
	return cipher(text, Encode)
}

func decode(text string) string {
	return cipher(text, Decode)
}

func init() {
	flag.BoolVar(&encodeFlag, "encode", false, "Encode the text from the standard input")
	flag.BoolVar(&decodeFlag, "decode", false, "Decode the text from the standard input")

	flag.Parse()
}

func main() {
	if encodeFlag && decodeFlag {
		fmt.Fprintf(os.Stderr, "error: conflicting options: encode / decode\n")
		os.Exit(1)
	}

	if !encodeFlag && !decodeFlag {
		os.Exit(0)
	}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := s.Text()
		if encodeFlag {
			fmt.Print(encode(text))
		} else if decodeFlag {
			fmt.Print(decode(text))
		}
	}
}
