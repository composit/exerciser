// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var (
		ss      = s
		decimal string
		buf     bytes.Buffer
	)

	if s[0] == '-' || s[0] == '+' {
		ss = s[1:]
		buf.WriteByte(s[0])
	}

	if dot := strings.LastIndex(ss, "."); dot >= 0 {
		decimal = ss[dot:]
		ss = ss[:dot]
	}

	n := len(ss)
	if n <= 3 {
		return s
	}

	offset := len(ss) % 3
	if offset == 0 {
		offset = 3
	}
	buf.WriteString(ss[:offset])
	for i := offset; i < len(ss); i += 3 {
		buf.WriteRune(',')
		buf.WriteString(ss[i : i+3])
	}

	buf.WriteString(decimal)

	return buf.String()
}

//!-
