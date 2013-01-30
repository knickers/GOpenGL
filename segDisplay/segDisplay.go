package segDisplay

import (
	gl "github.com/chsc/gogl/gl21"
)

func New(width, height float32) {
	return new(SegDisplay{
		mW: width / 2,
		mH: height / 2,
	})
}

type SegDisplay struct {
	mW, mH float32
}

func (s *segDisplay) Draw(ch rune) {
	if (ch == '0' || ch == '2' || ch == '3' || ch == '5' || ch == '6' ||
		ch == '7' || ch == '8' || ch == '9' || ch == 'A' || ch == 'B' ||
		ch == 'C' || ch == 'D' || ch == 'E' || ch == 'F' || ch == 'G' ||
		ch == 'I' || ch == 'O' || ch == 'P' || ch == 'Q' || ch == 'R' ||
		ch == 'S' || ch == 'T' || ch == 'Z') {
		s.top()
	}
	if (ch == '0' || ch == '4' || ch == '5' || ch == '6' || ch == '8' ||
		ch == '9' || ch == 'A' || ch == 'C' || ch == 'E' || ch == 'F' ||
		ch == 'G' || ch == 'H' || ch == 'K' || ch == 'L' || ch == 'M' ||
		ch == 'N' || ch == 'O' || ch == 'P' || ch == 'Q' || ch == 'R' ||
		ch == 'U' || ch == 'V' || ch == 'W') {
		s.upperLeft()
	}
	if ch == 'M' || ch == 'N' || ch == 'S' || ch == 'X' || ch == 'Y' {
		s.upperMidLeft()
	}
	if ch == 'B' || ch == 'D' || ch == 'I' || ch == 'T' {
		s.upperMid()
	}
	if (ch == '1' || ch == '7' || ch == 'K' || ch == 'M' || ch == 'V' ||
		ch == 'X' || ch == 'Y' || ch == 'Z') {
		s.upperMidRight()
	}
	if (ch == '0' || ch == '1' || ch == '2' || ch == '3' || ch == '4' ||
		ch == '8' || ch == '9' || ch == 'A' || ch == 'B' || ch == 'D' ||
		ch == 'H' || ch == 'J' || ch == 'M' || ch == 'N' || ch == 'O' ||
		ch == 'P' || ch == 'Q' || ch == 'R' || ch == 'U' || ch == 'W') {
		s.upperRight()
	}
	if (ch == '2' || ch == '4' || ch == '5' || ch == '6' || ch == '8' ||
		ch == '9' || ch == 'A' || ch == 'E' || ch == 'F' || ch == 'H' ||
		ch == 'K' || ch == 'P' || ch == 'R') {
		s.MidLeft()
	}
	if (ch == '2' || ch == '3' || ch == '4' || ch == '5' || ch == '6' ||
		ch == '8' || ch == '9' || ch == 'A' || ch == 'B' || ch == 'E' ||
		ch == 'F' || ch == 'G' || ch == 'H' || ch == 'P' || ch == 'R' ||
		ch == 'S') {
		s.MidRight()
	}
	if (ch == '0' || ch == '2' || ch == '4' || ch == '6' || ch == '8' ||
		ch == 'A' || ch == 'C' || ch == 'E' || ch == 'F' || ch == 'G' ||
		ch == 'H' || ch == 'J' || ch == 'K' || ch == 'L' || ch == 'M' ||
		ch == 'N' || ch == 'O' || ch == 'P' || ch == 'Q' || ch == 'R' ||
		ch == 'U' || ch == 'V' || ch == 'W') {
		s.lowerLeft()
	}
	if ch == 'V' || ch == 'W' || ch == 'X' || ch == 'Z' {
		s.lowerMidLeft()
	}
	if (ch == '7' || ch == 'B' || ch == 'D' || ch == 'I' || ch == 'T' ||
		ch == 'Y') {
		s.lowerMid()
	}
	if (ch == 'K' || ch == 'N' || ch == 'Q' || ch == 'R' || ch == 'W' ||
		ch == 'X') {
		s.lowerMidRight()
	}
	if (ch == '1' || ch == '3' || ch == '4' || ch == '5' || ch == '6' ||
		ch == '8' || ch == '9' || ch == 'A' || ch == 'B' || ch == 'D' ||
		ch == 'G' || ch == 'H' || ch == 'J' || ch == 'M' || ch == 'N' ||
		ch == 'O' || ch == 'Q' || ch == 'S' || ch == 'U' || ch == 'W') {
		s.lowerRight()
	}
	if (ch == '2' || ch == '3' || ch == '5' || ch == '6' || ch == '8' ||
		ch == 'B' || ch == 'C' || ch == 'D' || ch == 'E' || ch == 'G' ||
		ch == 'I' || ch == 'J' || ch == 'L' || ch == 'O' || ch == 'Q' ||
		ch == 'S' || ch == 'U' || ch == 'Z') {
		s.bottom()
	}
}

func (s *SegDisplay) top() {
	gl.Begin(gl.POLYGON)
	gl.Vertex2d(-s.mW*0.8, +s.mH*0.8)
	gl.Vertex2d(-s.mW*0.9, +s.mH*0.9)
	gl.Vertex2d(-s.mW*0.8, +s.mH*1.0)
	gl.Vertex2d(+s.mW*0.8, +s.mH*1.0)
	gl.Vertex2d(+s.mW*0.9, +s.mH*0.9)
	gl.Vertex2d(+s.mW*0.8, +s.mH*0.8)
	gl.End()
}

func (s *SegDisplay) upperLeft() {
	gl.Begin(gl.POLYGON)
	gl.Vertex2d(-s.mW*1.0, +s.mH*0.8)
	gl.Vertex2d(-s.mW*0.9, +s.mH*0.9)
	gl.Vertex2d(-s.mW*0.8, +s.mH*0.8)
	gl.Vertex2d(-s.mW*0.8, +s.mH*0.1)
	gl.Vertex2d(-s.mW*0.9, +s.mH*0.0)
	gl.Vertex2d(-s.mW*1.0, +s.mH*0.1)
	gl.End()
}

func (s *SegDisplay) upperMidLeft() {
	gl.Begin(gl.POLYGON)
	gl.Vertex2d(-s.mW*0.8, +s.mH*0.1)
	gl.Vertex2d(-s.mW*0.9, +s.mH*0.1)
	gl.Vertex2d(-s.mW*0.9, +s.mH*0.2)
	gl.Vertex2d(-s.mW*0.1, +s.mH*0.9)
	gl.Vertex2d(-s.mW*0.0, +s.mH*0.9)
	gl.Vertex2d(-s.mW*0.0, +s.mH*0.8)
	gl.End()
}

func (s *SegDisplay) upperMid() {
	gl.Begin(gl.POLYGON)
	gl.Vertex2d(-s.mW*0.1, +s.mH*0.8)
	gl.Vertex2d(+s.mW*0.0, +s.mH*0.9)
	gl.Vertex2d(+s.mW*0.1, +s.mH*0.8)
	gl.Vertex2d(+s.mW*0.1, +s.mH*0.1)
	gl.Vertex2d(+s.mW*0.0, +s.mH*0.0)
	gl.Vertex2d(-s.mW*0.1, +s.mH*0.1)
	gl.End()
}

func (s *SegDisplay) upperMidRight() {
	gl.Begin(gl.POLYGON)
	gl.Vertex2d(+s.mW*0.8, +s.mH*0.1)
	gl.Vertex2d(+s.mW*0.9, +s.mH*0.1)
	gl.Vertex2d(+s.mW*0.9, +s.mH*0.2)
	gl.Vertex2d(+s.mW*0.1, +s.mH*0.9)
	gl.Vertex2d(+s.mW*0.0, +s.mH*0.9)
	gl.Vertex2d(+s.mW*0.0, +s.mH*0.8)
	gl.End()
}

func (s *SegDisplay) upperRight() {
	gl.Begin(gl.POLYGON)
	gl.Vertex2d(+s.mW*1.0, +s.mH*0.8)
	gl.Vertex2d(+s.mW*0.9, +s.mH*0.9)
	gl.Vertex2d(+s.mW*0.8, +s.mH*0.8)
	gl.Vertex2d(+s.mW*0.8, +s.mH*0.1)
	gl.Vertex2d(+s.mW*0.9, +s.mH*0.0)
	gl.Vertex2d(+s.mW*1.0, +s.mH*0.1)
	gl.End()
}

func (s *SegDisplay) midLeft() {
	gl.Begin(gl.POLYGON)
	gl.Vertex2d(-s.mW*0.8, -s.mH*0.1)
	gl.Vertex2d(-s.mW*0.9, +s.mH*0.0)
	gl.Vertex2d(-s.mW*0.8, +s.mH*0.1)
	gl.Vertex2d(-s.mW*0.1, +s.mH*0.1)
	gl.Vertex2d(-s.mW*0.0, +s.mH*0.0)
	gl.Vertex2d(-s.mW*0.1, -s.mH*0.1)
	gl.End()
}

func (s *SegDisplay) midRight() {
	gl.Begin(gl.POLYGON)
	gl.Vertex2d(+s.mW*0.1, -s.mH*0.1)
	gl.Vertex2d(+s.mW*0.0, +s.mH*0.0)
	gl.Vertex2d(+s.mW*0.1, +s.mH*0.1)
	gl.Vertex2d(+s.mW*0.8, +s.mH*0.1)
	gl.Vertex2d(+s.mW*0.9, +s.mH*0.0)
	gl.Vertex2d(+s.mW*0.8, -s.mH*0.1)
	gl.End()
}

func (s *SegDisplay) lowerLeft() {
	gl.Begin(gl.POLYGON)
	gl.Vertex2d(-s.mW*1.0, -s.mH*0.8)
	gl.Vertex2d(-s.mW*0.9, -s.mH*0.9)
	gl.Vertex2d(-s.mW*0.8, -s.mH*0.8)
	gl.Vertex2d(-s.mW*0.8, -s.mH*0.1)
	gl.Vertex2d(-s.mW*0.9, -s.mH*0.0)
	gl.Vertex2d(-s.mW*1.0, -s.mH*0.1)
	gl.End()
}

func (s *SegDisplay) lowerMidLeft() {
	gl.Begin(gl.POLYGON)
	gl.Vertex2d(-s.mW*0.8, -s.mH*0.1)
	gl.Vertex2d(-s.mW*0.9, -s.mH*0.1)
	gl.Vertex2d(-s.mW*0.9, -s.mH*0.2)
	gl.Vertex2d(-s.mW*0.1, -s.mH*0.9)
	gl.Vertex2d(-s.mW*0.0, -s.mH*0.9)
	gl.Vertex2d(-s.mW*0.0, -s.mH*0.8)
	gl.End()
}

func (s *SegDisplay) lowerMid() {
	gl.Begin(gl.POLYGON)
	gl.Vertex2d(-s.mW*0.1, -s.mH*0.8)
	gl.Vertex2d(+s.mW*0.0, -s.mH*0.9)
	gl.Vertex2d(+s.mW*0.1, -s.mH*0.8)
	gl.Vertex2d(+s.mW*0.1, -s.mH*0.1)
	gl.Vertex2d(+s.mW*0.0, -s.mH*0.0)
	gl.Vertex2d(-s.mW*0.1, -s.mH*0.1)
	gl.End()
}

func (s *SegDisplay) lowerMidRight() {
	gl.Begin(gl.POLYGON)
	gl.Vertex2d(+s.mW*0.9, -s.mH*0.8)
	gl.Vertex2d(+s.mW*0.9, -s.mH*0.9)
	gl.Vertex2d(+s.mW*0.8, -s.mH*0.9)
	gl.Vertex2d(+s.mW*0.0, -s.mH*0.2)
	gl.Vertex2d(+s.mW*0.0, -s.mH*0.1)
	gl.Vertex2d(+s.mW*0.1, -s.mH*0.1)
	gl.End()
}

func (s *SegDisplay) lowerRight() {
	gl.Begin(gl.POLYGON)
	gl.Vertex2d(+s.mW*1.0, -s.mH*0.8)
	gl.Vertex2d(+s.mW*0.9, -s.mH*0.9)
	gl.Vertex2d(+s.mW*0.8, -s.mH*0.8)
	gl.Vertex2d(+s.mW*0.8, -s.mH*0.1)
	gl.Vertex2d(+s.mW*0.9, -s.mH*0.0)
	gl.Vertex2d(+s.mW*1.0, -s.mH*0.1)
	gl.End()
}

func (s *SegDisplay) bottom() {
	gl.Begin(gl.POLYGON)
	gl.Vertex2d(-s.mW*0.8, -s.mH*0.8)
	gl.Vertex2d(-s.mW*0.9, -s.mH*0.9)
	gl.Vertex2d(-s.mW*0.8, -s.mH*1.0)
	gl.Vertex2d(+s.mW*0.8, -s.mH*1.0)
	gl.Vertex2d(+s.mW*0.9, -s.mH*0.9)
	gl.Vertex2d(+s.mW*0.8, -s.mH*0.8)
	gl.End()
}
