// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package lenconv performs Meters and Feet conversions.
package lenconv

import "fmt"

type Meters float64
type Feet float64

func (m Meters) String() string { return fmt.Sprintf("%g meters", m) }
func (f Feet) String() string   { return fmt.Sprintf("%g feet", f) }

//!-
