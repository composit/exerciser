// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/png"
	"math/big"
	"os"
)

func buildBigRat() {
	const (
		xmin, ymin, xmax, ymax int64 = -2, -2, +2, +2
		width, height          int64 = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	for py := int64(0); py < height; py++ {
		yr, yq, ym, yy := &big.Rat{}, &big.Rat{}, &big.Rat{}, &big.Rat{}
		yr.Sub(big.NewRat(ymax, 1), big.NewRat(ymin, 1))
		yq.Quo(big.NewRat(py, 1), big.NewRat(height, 1))
		ym.Mul(yq, yr)
		yy.Add(ym, big.NewRat(ymin, 1))
		for px := int64(0); px < width; px++ {
			xr, xq, xm, xx := &big.Rat{}, &big.Rat{}, &big.Rat{}, &big.Rat{}
			xr.Sub(big.NewRat(xmax, 1), big.NewRat(xmin, 1))
			xq.Quo(big.NewRat(px, 1), big.NewRat(width, 1))
			xm.Mul(xq, xr)
			xx.Add(xm, big.NewRat(xmin, 1))
			x, _ := xx.Float64()
			y, _ := yy.Float64()

			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(int(px), int(py), mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}
