package text

import (
	"fmt"
	gl "github.com/chsc/gogl/gl21"
)

func Text(x, y, z float32, str string) {
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.BLEND)

	gl.RasterPos3f(x, y, z)
	for _, ch := range str {
		//glut.BitmapCharacter(glut.BITMAP_9_BY_15, string(ch))
	}

	gl.Disable(gl.BLEND)
}

func Number(x, y, z, n float32) {
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.BLEND)

	s := ""
	fmt.Fprintf(s, "%f", n)
	gl.RasterPos3f(x, y, z)
	for _, ch := range s {
		//glut.BitmapCharacter(glut.BITMAP_9_BY_15, string(ch))
	}

	gl.Disable(gl.BLEND)
}

func StrokeText(x, y, z float32, str string) {
}

func StrokeNumber(x, y, z, n float32) {
}
