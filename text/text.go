package text

import (
	//"fmt"
	gl "github.com/chsc/gogl/gl21"
)

func String(x, y, z float64, str string) {
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.BLEND)

	gl.RasterPos3d(gl.Double(x), gl.Double(y), gl.Double(z))
	for _, ch := range str {
		if ch == 'a' {
			//fmt.Print(ch)
		}
		//glut.BitmapCharacter(glut.BITMAP_9_BY_15, string(ch))
	}

	gl.Disable(gl.BLEND)
}

func Float(x, y, z, n float64) {
	s := ""
	//fmt.Sprintf(s, "%f", n)
	String(x, y, z, s)
}

func Int(x, y, z float64, n int) {
	s := ""
	//fmt.Sprintf(s, "%d", n)
	String(x, y, z, s)
}

func StrokeText(x, y, z float64, str string) {
}

func StrokeNumber(x, y, z, n float64) {
}
