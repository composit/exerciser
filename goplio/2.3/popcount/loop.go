// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcount

// PopCount returns the population count (number of set bits) of x.
func PopCountLoop(x uint64) int {
	result := 0
	for i := uint64(0); i <= 7; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}

//!-
