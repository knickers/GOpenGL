package shapes

import (
	//"fmt"
	gl "github.com/chsc/gogl/gl21"
	"math"
)

func Box(x, y, z float32, solid bool) {
	mode := gl.LINE_LOOP
	if solid {
		mode = gl.QUADS
	}
	X := gl.Float(x) / 2
	Y := gl.Float(y) / 2
	Z := gl.Float(z) / 2
	gl.Begin(gl.Enum(mode))

	gl.Normal3f(0, 0, 1) // Positive Z Face
	gl.Vertex3f(-X, -Y, Z)
	gl.Vertex3f(X, -Y, Z)
	gl.Vertex3f(X, Y, Z)
	gl.Vertex3f(-X, Y, Z)

	gl.Normal3f(0, 0, -1) // Negative Z Face
	gl.Vertex3f(-X, -Y, -Z)
	gl.Vertex3f(-X, Y, -Z)
	gl.Vertex3f(X, Y, -Z)
	gl.Vertex3f(X, -Y, -Z)

	gl.Normal3f(0, 1, 0) // Positive Y Face
	gl.Vertex3f(-X, Y, -Z)
	gl.Vertex3f(-X, Y, Z)
	gl.Vertex3f(X, Y, Z)
	gl.Vertex3f(X, Y, -Z)

	gl.Normal3f(0, -1, 0) // Negative Y Face
	gl.Vertex3f(-X, -Y, -Z)
	gl.Vertex3f(X, -Y, -Z)
	gl.Vertex3f(X, -Y, Z)
	gl.Vertex3f(-X, -Y, Z)

	gl.Normal3f(1, 0, 0) // Positive X Face
	gl.Vertex3f(X, -Y, -Z)
	gl.Vertex3f(X, Y, -Z)
	gl.Vertex3f(X, Y, Z)
	gl.Vertex3f(X, -Y, Z)

	gl.Normal3f(-1, 0, 0) // Negative X Face
	gl.Vertex3f(-X, -Y, -Z)
	gl.Vertex3f(-X, -Y, Z)
	gl.Vertex3f(-X, Y, Z)
	gl.Vertex3f(-X, Y, -Z)

	gl.End()
}

func Cylinder(r, h gl.Float, slices int, hollow, solid bool) {
	res := 2 * math.Pi / float64(slices)
	mode := gl.LINE_LOOP
	if solid {
		mode = gl.QUADS
	}
	gl.Begin(gl.Enum(mode))
	for a := 0.0; a < 2*math.Pi; a += res {
		gl.Normal3f(gl.Float(math.Cos(a)), gl.Float(math.Sin(a)), 0)
		gl.Vertex3f(r*gl.Float(math.Cos(a)), r*gl.Float(math.Sin(a)), 0)
		gl.Vertex3f(r*gl.Float(math.Cos(a)), r*gl.Float(math.Sin(a)), h)
		a += res
		gl.Vertex3f(r*gl.Float(math.Cos(a)), r*gl.Float(math.Sin(a)), h)
		gl.Vertex3f(r*gl.Float(math.Cos(a)), r*gl.Float(math.Sin(a)), 0)
	}
	gl.End()
	if !hollow {
		// X Y plane
		if h < 0 {
			gl.Normal3f(0, 0, 1)
		} else {
			gl.Normal3f(0, 0, -1)
		}
		Circle(r, slices, solid)
		// Top (or bottom)
		if h < 0 {
			gl.Normal3f(0, 0, -1)
		} else {
			gl.Normal3f(0, 0, 1)
		}
		gl.Translatef(0, 0, h)
		Circle(r, slices, solid)
	}
}
