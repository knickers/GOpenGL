package colors

import (
	gl "github.com/chsc/gogl/gl21"
)

var (
	Black       = []gl.Float{0.0, 0.0, 0.0, 1}
	White       = []gl.Float{1.0, 1.0, 1.0, 1}
	Red         = []gl.Float{1.0, 0.0, 0.0, 1}
	RedTrans    = []gl.Float{1.0, 0.0, 0.0, 0.75}
	Green       = []gl.Float{0.0, 1.0, 0.0, 1}
	Blue        = []gl.Float{0.0, 0.0, 1.0, 1}
	BlueDark    = []gl.Float{0.1, 0.5, 0.6, 1}
	BlueLight   = []gl.Float{0.1, 0.5, 0.8, 1}
	BlueTrans   = []gl.Float{0.0, 0.0, 1.0, 0.35}
	Yellow      = []gl.Float{1.0, 1.0, 0.0, 1}
	YellowTrans = []gl.Float{1.0, 1.0, 0.0, 0.5}
	Navy        = []gl.Float{0.2, 0.2, 0.7, 1}
	NavyTrans   = []gl.Float{0.0, 0.0, 0.5, 0.5}
	Silver      = []gl.Float{0.7, 0.8, 0.8, 1}
	Ivery       = []gl.Float{0.8, 0.9, 0.5, 1}
	Brown       = []gl.Float{0.3, 0.2, 0.1, 1}
	Gray        = []gl.Float{0.4, 0.4, 0.3, 1}
	Gray10      = []gl.Float{0.1, 0.1, 0.1, 1}
	Gray15      = []gl.Float{.15, .15, .15, 1}
	Gray25      = []gl.Float{.25, .25, .25, 1}
	Gray50      = []gl.Float{0.5, 0.5, 0.5, 1}
	Gray75      = []gl.Float{.75, .75, .75, 1}
	GrayLight   = []gl.Float{0.6, 0.6, 0.5, 1}
	GrayDark    = []gl.Float{0.3, 0.3, 0.2, 1}
	Clear       = []gl.Float{0.0, 0.0, 0.0, 0}
	OrangeTrans = []gl.Float{1.0, 0.5, .25, 0.5}
)

func Material(c []gl.Float) {
	gl.Materialfv(gl.FRONT, gl.AMBIENT_AND_DIFFUSE, &c[0])
}

func Custom(r, g, b, a gl.Float) {
	color := []gl.Float{r, g, b, a}
	gl.Materialfv(gl.FRONT, gl.AMBIENT_AND_DIFFUSE, &color[0])
}
