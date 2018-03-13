// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package lenconv

const conversion = 3.28

// MToF converts a Meters temperature to Feet.
func MToF(m Meters) Feet { return Feet(m / conversion) }

// FToM converts a Feet temperature to Meters.
func FToM(f Feet) Meters { return Meters(f * conversion) }

//!-
