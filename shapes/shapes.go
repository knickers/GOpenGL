package shapes

import (
	//"fmt"
	"GOpenGL/text"
	gl "github.com/chsc/gogl/gl21"
	//"math"
)

type Point struct {
	X, Y, Z gl.Float
}

var Origin = Point{0, 0, 0}

func Axis(size gl.Float) {
	Line(Origin, Point{size, 0, 0}) // X
	Line(Origin, Point{0, size, 0}) // Y
	Line(Origin, Point{0, 0, size}) // Z

	gl.Color3d(0, 0, 0)
	text.String(float64(size+1), 0, 0, "X")
	text.String(0, float64(size+1), 0, "Y")
	text.String(0, 0, float64(size+1), "Z")
}
