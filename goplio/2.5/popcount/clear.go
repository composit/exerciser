// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcount

// PopCount returns the population count (number of set bits) of x.
func PopCountClear(x uint64) int {
	var result uint64
	for i := uint64(1); i < 64; i++ {
		x = x & (x - 1)
		result++
		if x == 0 {
			break
		}
	}
	return int(result)
}

//!-
