// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package weightconv performs Kilograms and Pounds conversions.
package weightconv

import "fmt"

type Kilograms float64
type Pounds float64

func (k Kilograms) String() string { return fmt.Sprintf("%g kilograms", k) }
func (p Pounds) String() string    { return fmt.Sprintf("%g pounds", p) }

//!-
