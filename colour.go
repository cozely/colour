// Copyright 2018-2019 Laurent Moussault <laurent.moussault@gmail.com>
// SPDX-License-Identifier: BSD-2-Clause

package colour

import (
	"image/color"
	"math"
)

////////////////////////////////////////////////////////////////////////////////

// Colour can convert itself to alpha-premultipled RGBA as 32 bit floats, in
// both linear and standard (sRGB) color spaces.
type Colour interface {
	color.Color

	// Linear returns the color in alpha-premultiplied linear color space. Each
	// value ranges within [0, 1].
	Linear() (r, g, b, a float32)
}

////////////////////////////////////////////////////////////////////////////////

func linearOf(c float32) float32 {
	if c <= 0.04045 {
		return c / 12.92
	}
	return float32(math.Pow(float64(c+0.055)/(1+0.055), 2.4))
}

func standardOf(c float32) float32 {
	if c <= 0.0031308 {
		return 12.92 * c
	}
	return (1+0.055)*float32(math.Pow(float64(c), 1/2.4)) - 0.055
}

////////////////////////////////////////////////////////////////////////////////

// // ColorsFrom returns a new Palette created from the file at the specified
// // path.
// func ColorsFrom(r io.Reader) ([]Colour, error) {
// 	var pal = []Colour{}

// 	cf, _, err := image.DecodeConfig(r)
// 	if err != nil {
// 		return pal, errors.New("unable to decode file for palette")
// 	}

// 	p, ok := cf.ColorModel.(color.Palette)
// 	if !ok {
// 		return pal, errors.New("image file not paletted for palette")
// 	}

// 	for i := range p {
// 		r, g, b, al := p[i].RGBA()
// 		if i > 255 {
// 			return pal, errors.New("too many colors for palette")
// 		}
// 		c := SRGBA{
// 			R: float32(r) / float32(0xFFFF),
// 			G: float32(g) / float32(0xFFFF),
// 			B: float32(b) / float32(0xFFFF),
// 			A: float32(al) / float32(0xFFFF),
// 		}
// 		//TODO: append name
// 		pal = append(pal, c)
// 	}

// 	internal.Debug.Printf("Loaded color palette (%d entries) from %s", len(p), path)

// 	return pal, nil
// }
