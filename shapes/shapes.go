package shapes

import (
	//"fmt"
	"GOpenGL/text"
	gl "github.com/chsc/gogl/gl21"
	"math"
)

type Point struct {
	X, Y, Z gl.Float
}

var Origin = Point{0, 0, 0}

func Axis(size gl.Float) {
	Line(Origin, Point{size, 0, 0}) // X
	Line(Origin, Point{0, size, 0}) // Y
	Line(Origin, Point{0, 0, size}) // Z

	gl.Color3d()
	text.Text(size+1, 0, 0, "X")
	text.Text(0, size+1, 0, "Y")
	text.Text(0, 0, size+1, "Z")
}
