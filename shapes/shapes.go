package shapes

import (
	//"fmt"
	gl "github.com/chsc/gogl/gl21"
	"math"
)

type Point struct {
	X, Y, Z gl.Float
}

var Origin = Point{0, 0, 0}

//****************************************************************************//
//******************************** 2D Shapes *********************************//
//****************************************************************************//
func Line(one, two Point) {
	gl.Begin(gl.LINES)
	gl.Vertex3f(one.X, one.Y, one.Z)
	gl.Vertex3f(two.X, two.Y, two.Z)
	gl.End()
}

func Triangle(one, two, three Point, solid bool) {
	mode := gl.LINE_LOOP
	if solid {
		mode = gl.TRIANGLES
	}
	gl.Begin(gl.Enum(mode))
	gl.Vertex3f(one.X, one.Y, one.Z)
	gl.Vertex3f(two.X, two.Y, two.Z)
	gl.Vertex3f(three.X, three.Y, three.Z)
	gl.End()
}

func Rectangle(one, two Point, solid bool) {
	mode := gl.LINE_LOOP
	if solid {
		mode = gl.QUADS
	}
	gl.Begin(gl.Enum(mode))          // Clockwise from top right
	gl.Vertex3f(one.X, one.Y, one.Z) // Q1
	gl.Vertex3f(one.X, two.Y, one.Z) // Q4
	gl.Vertex3f(two.X, two.Y, two.Z) // Q3
	gl.Vertex3f(two.X, one.Y, two.Z) // Q2
	gl.End()
}

func FadeRectangle(one, two Point, horizontal bool) {
	gl.Begin(gl.QUADS)
	if horizontal { // Clockwise from top right
		gl.Vertex3f(one.X, one.Y, one.Z) // Q1
		gl.Vertex3f(one.X, two.Y, one.Z) // Q4
		gl.Color3f(1, 1, 1)
		gl.Vertex3f(two.X, two.Y, two.Z) // Q3
		gl.Vertex3f(two.X, one.Y, two.Z) // Q2
	} else { // Counter clockwise from top right
		gl.Vertex3f(one.X, one.Y, one.Z) // Q1
		gl.Vertex3f(two.X, one.Y, one.Z) // Q2
		gl.Color3f(1, 1, 1)
		gl.Vertex3f(two.X, two.Y, two.Z) // Q3
		gl.Vertex3f(one.X, two.Y, two.Z) // Q4
	}
	gl.End()
}

func Circle(r gl.Float, slices int, solid bool) {
	mode := gl.LINE_LOOP
	if solid {
		mode = gl.POLYGON
	}
	res := 2 * math.Pi / float64(slices)
	gl.Begin(gl.Enum(mode))
	gl.Vertex2f(0, 0)
	for a := 0.0; a < 2*math.Pi; a += res {
		gl.Vertex2f(r*gl.Float(math.Cos(a)), r*gl.Float(math.Sin(a)))
	}
	gl.End()
}

func Ellipse(rX, rY gl.Float, slices int, solid bool) {
	mode := gl.LINE_LOOP
	if solid {
		mode = gl.POLYGON
	}
	res := 2 * math.Pi / float64(slices)
	gl.Begin(gl.Enum(mode))
	gl.Vertex2f(0, 0)
	for a := 0.0; a < 2*math.Pi; a += res {
		gl.Vertex2f(rX*gl.Float(math.Cos(a)), rY*gl.Float(math.Sin(a)))
	}
	gl.End()
}

//****************************************************************************//
//******************************** 3D Shapes *********************************//
//****************************************************************************//
func Axis(size gl.Float) {
	Line(Origin, Point{size, 0, 0}) // X
	Line(Origin, Point{0, size, 0}) // Y
	Line(Origin, Point{0, 0, size}) // Z
}

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
