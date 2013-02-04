// Slider package by Nick Cox

// Ported From:
// Ferbruary 2011 c++ slider class
// Reconstructed and converted to horizontal and vertical subclasses.
// HandleMouseClick converted into two different methods, click and drag
// This prevents grabbing an adjacent slider by accident
// Added different drawing styles

// Derived from:
// October 2008 c++ Slider class by Bart Stander for cs 1410
// This package is for maintaining a slider bar in OpenGL.
// It remembers its own position, color, and value.

package slider

import (
	gl "github.com/chsc/gogl/gl21"
	"GOpenGL/shapes"
	"linAlg"
	"math"
)

type Style int
type Direction int

const (
	// Types
	BASIC = Style(0)
	SHINY
	BUBBLE
	NARROW
	BLOCK
	// Directions
	DIR_UP = Direction(0)
	DIR_DOWN
	DIR_RIGHT
	DIR_LEFT
)

func New(style Style, dir Direction, width, height float64) {
	return Slider{
		Height: height,
		Width:  width,
		Type:   style,
		Dir:    dir,
	}
}

type Slider struct {
	Color   []float64
	grabbed bool
	Type    Style
	Dir     Direction
	Pos     linAlg.Point
	bottom, top, right, left, Width, Height, value float64
}

func (s *Slider) SetValue(v float64) {
	s.value = math.Max(math.Min(v, 1), 0)
}

func (s *Slider) GetValue() float64 {
	return s.value
}

func (s *Slider) Click(x, y int) bool {
	s.grabbed = false
	if x >= s.left && x <= s.right && y >= s.bottom && y <= s.top {
		s.grabbed = true
		s.Drag(x, y)
	}
	return s.grabbed
}

func isHorisontal(dir Direction) bool {
	return dir == DIR_RIGHT || dir == DIR_LEFT
}

func (s *Slider) Drag(x, y int) bool {
	if s.grabbed {
		if s.Dir == DIR_RIGHT {
			if x > s.right {
				s.value = 1
			} else if x < s.left {
				s.value = 0
			} else {
				s.value = (float64(x) - s.left) / s.Width
			}
		} else if s.Dir == DIR_LEFT {
			if x > s.right {
				s.value = 0
			} else if x < s.left {
				s.value = 1
			} else {
				s.value = (s.right - float64(x)) / s.Width
			}
		} else if s.Dir == DIR_UP {
			if y > s.top {
				s.value = 1
			} else if y < s.bottom {
				s.value = 0
			} else {
				s.value = (float64(y) - s.bottom) / s.Height
			}
		} else if s.Dir == DIR_LEFT {
			if y > s.top {
				s.value = 0
			} else if y < s.bottom {
				s.value = 1
			} else {
				s.value = (s.top - float64(y)) / s.Height
			}
		}
		return true
	}
	return false
}

func (s *Slider) Increment(d float64) bool {
	s.value += d
	if s.value > 1 {
		s.value = 1
		return false
	} else if s.value < 0 {
		s.value = 0
		return false
	}
	return true
}

func (s *Slider) Draw() {
	gl.PushMatrix()
	gl.Translated(gl.Double(s.Pos.X), gl.Double(s.Pos.Y), gl.Double(s.Pos.Z))
	//gl.Rotated()
	switch s.Type {
	case BASIC:
		s.drawBasic()
	case SHINY:
		s.drawShiny()
	case BLOCK:
		s.drawBlock()
	case BUBBLE:
		s.drawBubble()
	case NARROW:
		s.drawNarrow()
	}
	gl.PopMatrix()
}

func (s *Slider) drawBasic() {
	middle := s.left + s.value*s.Width
	leftBottom := linAlg.Point{s.left, s.bottom, 0}

	gl.Color3d(s.Color[0], s.Color[1], s.Color[2])
	shapes.Rectangle(leftBottom, linAlg.Point{middle, s.top, 0}, true)

	gl.Color3d(0, 0, 0)
	shapes.Rectangle(leftBottom, linAlg.Point{s.right, s.top, 0}, false)
}

func (s *Slider) drawShiny() {
}

func (s *Slider) drawBlock() {
}

func (s *Slider) drawBubble() {
}

func (s *Slider) drawNarrow() {
}
