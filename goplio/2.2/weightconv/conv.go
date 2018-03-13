// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package weightconv

const conversion = 2.2046226218

// KToP converts a Kilograms temperature to Pounds.
func KToP(k Kilograms) Pounds { return Pounds(k * conversion) }

// PToK converts a Pounds temperature to Kilograms.
func PToK(p Pounds) Kilograms { return Kilograms(p / conversion) }

//!-
