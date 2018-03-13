// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/composit/exerciser/goplio/2.2/lenconv"
	"github.com/composit/exerciser/goplio/2.2/tempconv"
	"github.com/composit/exerciser/goplio/2.2/weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		fahr := tempconv.Fahrenheit(t)
		cels := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			fahr, tempconv.FToC(fahr), cels, tempconv.CToF(cels))

		feet := lenconv.Feet(t)
		meters := lenconv.Meters(t)
		fmt.Printf("%s = %s, %s = %s\n",
			feet, lenconv.FToM(feet), meters, lenconv.MToF(meters))

		pounds := weightconv.Pounds(t)
		kilos := weightconv.Kilograms(t)
		fmt.Printf("%s = %s, %s = %s\n",
			pounds, weightconv.PToK(pounds), kilos, weightconv.KToP(kilos))

	}
}

//!-
