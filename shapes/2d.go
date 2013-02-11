package shapes

import (
	//"fmt"
	gl "github.com/chsc/gogl/gl21"
	"math"
)

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

// The rectangle is drawn with point one being at the top right
// and point two being at the bottom left.
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

// The rectangle is drawn with point one being at the top right
// and point two being at the bottom left.
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
	Ellipse(r, r, slices, solid)
}

func PartialCircle(r gl.Float, begRad, endRad float64, slices int, solid bool) {
	PartialEllipse(r, r, begRad, endRad, slices, solid)
}

func Ellipse(rX, rY gl.Float, slices int, solid bool) {
	PartialEllipse(rX, rY, 0, 2*math.Pi, slices, solid)
}

func PartialEllipse(rX, rY gl.Float, beginRad, endRad float64, slices int, solid bool) {
	mode := gl.LINE_LOOP
	if solid {
		mode = gl.POLYGON
	}
	res := (endRad - beginRad) / float64(slices)
	gl.Begin(gl.Enum(mode))
	gl.Vertex2f(0, 0)
	for a := beginRad; a <= endRad; a += res {
		gl.Vertex2f(rX*gl.Float(math.Cos(a)), rY*gl.Float(math.Sin(a)))
	}
	gl.Vertex2f(rX*gl.Float(math.Cos(endRad)), rY*gl.Float(math.Sin(endRad)))
	gl.End()
}
