package main

import (
	"image/color"
	"math/cmplx"
)

func main() {
	//buildFloat64()
	//buildComplex64()
	//buildComplex128()
	//buildBigFloat()
	buildBigRat()
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
