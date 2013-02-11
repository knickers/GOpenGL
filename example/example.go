package main

import (
	"fmt"
	gl "github.com/chsc/gogl/gl21"
	"github.com/jteeuwen/glfw"
	"GOpenGL"
	//"GOpenGL/camera"
	"GOpenGL/colors"
	"GOpenGL/shapes"
	"os"
	"os/signal"
)

const (
	gX = 640
	gY = 480
)

var G *GOpenGL.OpenGl

func main() {
	G = GOpenGL.New("Empty 3d Environment", gX, gY)

	G.Set3dCallback(display3d)
	G.Set2dCallback(display2d)
	G.SetKeyboardCallback(keyboard)
	G.SetMouseButtonCallback(mouseClick)
	G.SetMousePosCallback(mouseMove)
	G.SetMouseWheelCallback(mouseWheel)
	G.Init()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown
}

var (
	rotx gl.Float
	roty gl.Float
)

func display3d() {
	colors.Custom(1, 0, 0, 1)
	shapes.Axis(5)

	rotx += 0.5
	roty += 0.5

	gl.PushMatrix()
	//gl.Translatef(0, 0, -3.0)
	gl.Rotatef(rotx, 1, 0, 0)
	gl.Rotatef(roty, 0, 1, 0)
	colors.Material(colors.White)
	//shapes.Box(1, 1, 0.5, true)
	shapes.Cylinder(0.5, 1, 32, false, true)
	gl.PopMatrix()
}

func display2d() {
	gl.Color4f(0, 1, 0, 1)
	shapes.Rectangle(shapes.Point{100, 100, 0}, shapes.Point{10, 10, 0}, true)

	gl.PushMatrix()
	gl.Translatef(150, 55, 0)
	shapes.Circle(50, 3, true)
	gl.Translatef(120, 0, 0)
	shapes.Circle(50, 4, true)
	gl.Translatef(120, 0, 0)
	shapes.Circle(50, 5, true)
	gl.Translatef(120, 0, 0)
	shapes.Circle(50, 6, true)
	gl.PopMatrix()
}

func keyboard(key, state int) {
	fmt.Println("keyboard", key, state)
	switch key {
	case glfw.KeyEsc:
		os.Exit(0)
	case glfw.KeyTab:
	case glfw.KeyUp:
	case glfw.KeyDown:
	case glfw.KeyLeft:
	case glfw.KeyRight:
	default:
	}
}

func mouseMove(x, y int) {
	y = G.Height() - y
	fmt.Println("mousePos", x, y)
}

func mouseClick(button, state int) {
	fmt.Println("mouseClick", button, state)
}

func mouseWheel(delta int) {
	fmt.Println("mouseWheel", delta)
}
