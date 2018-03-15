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

func buildBigFloat() {
	const (
		xmin, ymin, xmax, ymax float64 = -2, -2, +2, +2
		width, height          float64 = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	for py := float64(0); py < height; py++ {
		var yr, yq, ym, yy = &big.Float{}, &big.Float{}, &big.Float{}, &big.Float{}
		yr.Sub(big.NewFloat(ymax), big.NewFloat(ymin))
		yq.Quo(big.NewFloat(py), big.NewFloat(height))
		ym.Mul(yq, yr)
		yy.Add(ym, big.NewFloat(ymin))
		for px := float64(0); px < width; px++ {
			var xr, xq, xm, xx = &big.Float{}, &big.Float{}, &big.Float{}, &big.Float{}
			xr.Sub(big.NewFloat(xmax), big.NewFloat(xmin))
			xq.Quo(big.NewFloat(px), big.NewFloat(width))
			xm.Mul(xq, xr)
			xx.Add(xm, big.NewFloat(xmin))
			x, _ := xx.Float64()
			y, _ := yy.Float64()

			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(int(px), int(py), mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}
