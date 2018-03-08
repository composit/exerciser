package main

import (
	"fmt"
)

func main() {
	fmt.Printf("1 degree C to K: %g\n", CToK(1))
	fmt.Printf("1 degree K to C: %g\n", KToC(1))
	fmt.Printf("1 degree F to K: %g\n", FToK(1))
	fmt.Printf("1 degree K to F: %g\n", KToF(1))
}
