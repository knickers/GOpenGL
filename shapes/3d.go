package shapes

import (
	"encoding/json"
	"fmt"
	gl "github.com/chsc/gogl/gl21"
	"io/ioutil"
	"linAlg"
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
	Rectangle(Point{X, Y, Z}, Point{-X, -Y, Z}, true)

	gl.Normal3f(0, 0, -1) // Negative Z Face
	Rectangle(Point{X, Y, -Z}, Point{-X, -Y, -Z}, true)

	gl.Normal3f(0, 1, 0) // Positive Y Face
	gl.Rotatef(90, 1, 0, 0)
	Rectangle(Point{X, Z, Y}, Point{-X, -Z, Y}, true)
	gl.Rotatef(-90, 1, 0, 0)

	gl.Normal3f(0, -1, 0) // Negative Y Face
	gl.Rotatef(-90, 1, 0, 0)
	Rectangle(Point{X, Z, Y}, Point{-X, -Z, Y}, true)
	gl.Rotatef(90, 1, 0, 0)

	gl.Normal3f(1, 0, 0) // Positive X Face
	gl.Rotatef(90, 0, 1, 0)
	Rectangle(Point{-Z, -Y, X}, Point{Z, Y, X}, true)
	gl.Rotatef(-90, 0, 1, 0)

	gl.Normal3f(-1, 0, 0) // Negative X Face
	gl.Rotatef(-90, 0, 1, 0)
	Rectangle(Point{-Z, -Y, X}, Point{Z, Y, X}, true)
	gl.Rotatef(90, 0, 1, 0)

	gl.End()
}

func Cylinder(r, h gl.Float, slices int, hollow, solid bool) {
	ExtrudedTaperedEllipse(r, r, r, r, h, slices, hollow, solid)
}

func TaperedCylinder(rBot, rTop, h gl.Float, slices int, hollow, solid bool) {
	ExtrudedTaperedEllipse(rBot, rBot, rTop, rTop, h, slices, hollow, solid)
}

func ExtrudedEllipse(rX, rY, h gl.Float, slices int, hollow, solid bool) {
	ExtrudedTaperedEllipse(rX, rY, rX, rY, h, slices, hollow, solid)
}

func ExtrudedTaperedEllipse(rXbot, rYbot, rXtop, rYtop, h gl.Float, slices int, hollow, solid bool) {
	ExtrudedTaperedPartialEllipse(rXbot, rYbot, rXtop, rYtop, h, 0, 2*math.Pi, slices, hollow, solid)
}

func ExtrudedTaperedPartialEllipse(rXbot, rYbot, rXtop, rYtop, h gl.Float, beginRad, endRad float64, slices int, hollow, solid bool) {
	res := (endRad - beginRad) / float64(slices)
	mode := gl.LINE_LOOP
	if solid {
		mode = gl.QUADS
	}
	gl.Begin(gl.Enum(mode))
	for a := beginRad; a < endRad; a += res {
		dx := math.Cos(a)
		dy := math.Sin(a)
		gl.Normal3f(gl.Float(dx), gl.Float(dy), 0)
		gl.Vertex3f(rXbot*gl.Float(dx), rYbot*gl.Float(dy), 0)
		gl.Vertex3f(rXtop*gl.Float(dx), rYtop*gl.Float(dy), h)
		dx = math.Cos(a + res)
		dy = math.Sin(a + res)
		gl.Normal3f(gl.Float(dx), gl.Float(dy), 0)
		gl.Vertex3f(rXtop*gl.Float(dx), rYtop*gl.Float(dy), h)
		gl.Vertex3f(rXbot*gl.Float(dx), rYbot*gl.Float(dy), 0)
	}
	gl.End()
	if !hollow {
		// X Y plane normal
		if h < 0 {
			gl.Normal3f(0, 0, 1)
		} else {
			gl.Normal3f(0, 0, -1)
		}
		PartialEllipse(rXbot, rYbot, beginRad, endRad, slices, solid)
		// Top (or bottom) normal
		if h < 0 {
			gl.Normal3f(0, 0, -1)
		} else {
			gl.Normal3f(0, 0, 1)
		}
		gl.Translatef(0, 0, h)
		PartialEllipse(rXtop, rYtop, beginRad, endRad, slices, solid)
	}
}

func CurvedCylinder(outR, inR, angle float64, stacks, slices int, solid bool) {
	ExtrudedCurvedEllipse(outR, inR, inR, angle, stacks, slices, solid)
}

func ExtrudedCurvedEllipse(r, w, h, radians float64, stacks, slices int, solid bool) {
	a := (math.Abs(radians)) / float64(slices)
	b := 2 * math.Pi / float64(stacks)
	twoPi := 2 * math.Pi
	if radians < 0 {
		gl.Rotatef(gl.Float(linAlg.RtoD(radians)), 0, 0, 1)
	}
	for i := 0.0; i < float64(math.Abs(radians)); i += a {
		for j := 0.0; i < twoPi; j += b {
			gl.Begin(gl.QUADS)

			gl.Normal3f(
				gl.Float(math.Cos(j)*math.Cos(i)),
				gl.Float(math.Cos(j)*math.Sin(i)),
				gl.Float(math.Sin(j)))
			gl.Vertex3f(
				gl.Float(r*math.Cos(i)+w*math.Cos(j)*math.Cos(i)),
				gl.Float(r*math.Sin(i)+w*math.Cos(j)*math.Sin(i)),
				gl.Float(h*math.Sin(j)))

			gl.Normal3f(
				gl.Float(math.Cos(j+b)*math.Cos(i)),
				gl.Float(math.Cos(j+b)*math.Sin(i)),
				gl.Float(math.Sin(j+b)))
			gl.Vertex3f(
				gl.Float(r*math.Cos(i)+w*math.Cos(j+b)*math.Cos(i)),
				gl.Float(r*math.Sin(i)+w*math.Cos(j+b)*math.Sin(i)),
				gl.Float(h*math.Sin(j+b)))

			gl.Normal3f(
				gl.Float(math.Cos(j+b)*math.Cos(i+a)),
				gl.Float(math.Cos(j+b)*math.Sin(i+a)),
				gl.Float(math.Sin(j+b)))
			gl.Vertex3f(
				gl.Float(r*math.Cos(i+a)+w*math.Cos(j+b)*math.Cos(i+a)),
				gl.Float(r*math.Sin(i+a)+w*math.Cos(j+b)*math.Sin(i+a)),
				gl.Float(h*math.Sin(j+b)))

			gl.Normal3f(
				gl.Float(math.Cos(j)*math.Cos(i+a)),
				gl.Float(math.Cos(j)*math.Sin(i+a)),
				gl.Float(math.Sin(j)))
			gl.Vertex3f(
				gl.Float(r*math.Cos(i+a)+w*math.Cos(j)*math.Cos(i+a)),
				gl.Float(r*math.Sin(i+a)+w*math.Cos(j)*math.Sin(i+a)),
				gl.Float(h*math.Sin(j)))

			gl.End()
		}
	}
	if solid {
		gl.PushMatrix()
		gl.Translatef(gl.Float(r), 0, 0)
		gl.Rotatef(90, 1, 0, 0)
		Ellipse(gl.Float(w), gl.Float(h), stacks, solid)
		gl.PopMatrix()

		gl.PushMatrix()
		gl.Rotatef(gl.Float(linAlg.RtoD(math.Abs(radians))), 0, 0, 1)
		gl.Translatef(gl.Float(r), 0, 0)
		gl.Rotatef(90, 1, 0, 0)
		Ellipse(gl.Float(w), gl.Float(h), stacks, solid)
		gl.PopMatrix()
	}
	if radians < 0 {
		gl.Rotatef(gl.Float(linAlg.RtoD(radians)), 0, 0, 1)
	}
}

func RotatePolygonFile(file string, radians float64, slices int, smoothVert, smoothHor bool) error {
	fp, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("ReadFile:", err)
		return err
	}

	var outline []linAlg.Point
	err = json.Unmarshal(fp, &outline)
	if err != nil {
		fmt.Println("Unmarshal:", err)
		return err
	}

	return RotatePolygonPoints(outline, radians, slices, smoothVert, smoothHor)
}

func RotatePolygonPoints(outline []linAlg.Point, radians float64, slices int, smoothVert, smoothHor bool) error {
	k := radians / float64(slices)
	var q1, q2, q3, q4, n1, n2, n3, n4 linAlg.Point

	gl.Begin(gl.QUADS)
	for i := 0.0; i < radians; i += k {
		for j, p1 := range outline {
			p2 := outline[j+1]
			q1 = linAlg.Point{p2.X * math.Cos(i+k), p2.Y * math.Sin(i+k), p2.Y}
			q2 = linAlg.Point{p2.X * math.Cos(i), p2.Y * math.Sin(i), p2.Y}
			q3 = linAlg.Point{p1.X * math.Cos(i), p1.Y * math.Sin(i), p1.Y}
			q4 = linAlg.Point{p1.X * math.Cos(i+k), p1.Y * math.Sin(i+k), p1.Y}
			n1 = linAlg.Point(linAlg.FindTriangleNormal(q1, q4, q3))
			n2, n3, n4 = n1, n1, n1

			if smoothHor {
				n3.X = p1.X - p2.X
				n3.Y = p1.Y - p2.Y
				theta := -math.Atan2(n3.X, n3.Y)

				n2.X = math.Cos(theta) * math.Cos(i)
				n2.Y = math.Cos(theta) * math.Sin(i)
				n2.Z = math.Sin(theta)
				n3.X = n2.X
				n3.Y = n2.Y
				n3.Z = n2.Z

				n1.X = math.Cos(theta) * math.Cos(i+k)
				n1.Y = math.Cos(theta) * math.Sin(i+k)
				n1.Z = math.Sin(theta)
				n4.X = n1.X
				n4.Y = n1.Y
				n4.Z = n1.Z
			}
			if smoothVert {
				if j < len(outline)-1 {
					n2.Z = (n2.Z + math.Sin(p1.Y-p2.Y)) / 2
					n1.Z = (n1.Z + math.Sin(p1.X-p2.X)) / 2
					n1.Z = math.Sin(math.Atan2(n2.Z, n1.Z))
					n2.Z = n1.Z
				}
				if j > 0 {
					p3 := outline[j-1]
					n3.Z = (n3.Z + math.Sin(p3.Y-p1.Y)) / 2
					n4.Z = (n4.Z + math.Sin(p3.Y-p1.Y)) / 2
					n1.Z = math.Sin(math.Atan2(n3.Z, n4.Z))
					n2.Z = n1.Z
				}
			}
			gl.Normal3f(gl.Float(n1.X), gl.Float(n1.Y), gl.Float(n1.Z))
			gl.Vertex3f(gl.Float(q1.X), gl.Float(q1.Y), gl.Float(q1.Z))
			gl.Normal3f(gl.Float(n2.X), gl.Float(n2.Y), gl.Float(n2.Z))
			gl.Vertex3f(gl.Float(q2.X), gl.Float(q2.Y), gl.Float(q2.Z))
			gl.Normal3f(gl.Float(n3.X), gl.Float(n3.Y), gl.Float(n3.Z))
			gl.Vertex3f(gl.Float(q3.X), gl.Float(q3.Y), gl.Float(q3.Z))
			gl.Normal3f(gl.Float(n4.X), gl.Float(n4.Y), gl.Float(n4.Z))
			gl.Vertex3f(gl.Float(q4.X), gl.Float(q4.Y), gl.Float(q4.Z))
		}
	}
	gl.End()
	return nil
}
