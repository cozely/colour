// Copyright 2019 Laurent Moussault <laurent.moussault@gmail.com>
// SPDX-License-Identifier: BSD-2-Clause

package colour

import "image/color"

// Standard color models for the colour types.
var (
	R8G8B8Model = color.ModelFunc(r8g8b8Model)
	RGBAModel = color.ModelFunc(func(c color.Color) color.Color{return RGBAof(c)})
)
