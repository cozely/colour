// Copyright 2018-2019 Laurent Moussault <laurent.moussault@gmail.com>
// SPDX-License-Identifier: BSD-2-Clause

package colour

import (
	"image/color"
)

////////////////////////////////////////////////////////////////////////////////

// RGBA represents a color in alpha-premultiplied linear color space. Each
// value ranges within [0, 1], and can be used directly by GPU shaders.
//
// An alpha-premultiplied color component c has been scaled by alpha (a), so has
// valid values 0 <= c <= a.
//
// Note that additive blending can also be achieved when alpha is set to 0 while
// the color components are non-null.
type RGBA struct {
	R float32
	G float32
	B float32
	A float32
}

////////////////////////////////////////////////////////////////////////////////

// RGBAof converts any color to alpha-premultiplied, linear color space.
func RGBAof(c color.Color) RGBA {
	switch c := c.(type) {
	case RGBA:
		return c
	case Colour:
		r, g, b, a := c.Linear()
		return RGBA{r, g, b, a}
	default:
		r, g, b, a := c.RGBA()
		return RGBA{
			R: linearOf(float32(r) / float32(0xFFFF)),
			G: linearOf(float32(g) / float32(0xFFFF)),
			B: linearOf(float32(b) / float32(0xFFFF)),
			A: float32(a) / float32(0xFFFF),
		}
	}
}

////////////////////////////////////////////////////////////////////////////////

// Linear implements the Colour interface.
func (c RGBA) Linear() (r, g, b, a float32) {
	return c.R, c.G, c.B, c.A
}

////////////////////////////////////////////////////////////////////////////////

// RGBA returns the colour in sRGB space, each component scaled by 0xFFFF
// (implements the standard color.Color interface).
func (c RGBA) RGBA() (r, g, b, a uint32) {
	r = uint32(standardOf(c.R) * 0xFFFF)
	g = uint32(standardOf(c.G) * 0xFFFF)
	b = uint32(standardOf(c.B) * 0xFFFF)
	a = uint32(c.A * 0xFFFF)
	return r, g, b, a
}

func rgbaModel(c color.Color) color.Color {
	switch c := c.(type) {
	case RGBA:
		return c
	case Colour:
		r, g, b, a := c.Linear()
		return RGBA{r, g, b, a}
	default:
		r, g, b, a := c.RGBA()
		return RGBA{
			R: linearOf(float32(r) / float32(0xFFFF)),
			G: linearOf(float32(g) / float32(0xFFFF)),
			B: linearOf(float32(b) / float32(0xFFFF)),
			A: float32(a) / float32(0xFFFF),
		}
	}
}
