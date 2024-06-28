//kage:unit pixels

//go:build ignore

package main

var Radius float

func Fragment(_ vec4, pos vec2, _ vec4) vec4 {
	zpos := pos
	r := Radius

	center := vec2(r, r)
	dist := distance(zpos, center)
	if dist > r {
		return vec4(0)
	}
	return vec4(0.4, 0.7, 0.9, 1)
}
