package colour

import "image/color"

////////////////////////////////////////////////////////////////////////////////

// R8G8B8 represents a fully opaque 24-bit color.
type R8G8B8 struct {
	R, G, B uint8
}

// RGBA implements the standard color.Color interface.
func (c R8G8B8) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	return r, g, b, 0xFFFF
}

func r8g8b8Model(c color.Color) color.Color {
	_, ok := c.(R8G8B8)
	if ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return R8G8B8{
		R: uint8(r>>8),
		G: uint8(g>>8),
		B: uint8(b>>8),
	}
}
