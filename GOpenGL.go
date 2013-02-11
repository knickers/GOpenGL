package GOpenGL

import (
	"fmt"
	gl "github.com/chsc/gogl/gl21"
	"github.com/jteeuwen/glfw"
	"GOpenGL/camera"
	"os"
	"time"
)

// A simple to use framework for building 3d models in openGl
type OpenGl struct {
	title      string
	width      int
	height     int
	Camera     *camera.Camera
	redisplay  chan bool
	display3d  func()
	display2d  func()
	keyboard   func(int, int) //glfw.KeyHandler
	mouseMove  func(int, int) //glfw.MousePosHandler
	mouseClick func(int, int) //glfw.MouseButtonHandler
	mouseWheel func(int)      //glfw.MouseWheelHandler
}

func (o *OpenGl) Set3dCallback(fn func())                  { o.display3d = fn }
func (o *OpenGl) Set2dCallback(fn func())                  { o.display2d = fn }
func (o *OpenGl) SetKeyboardCallback(fn func(int, int))    { o.keyboard = fn }
func (o *OpenGl) SetMousePosCallback(fn func(int, int))    { o.mouseMove = fn }
func (o *OpenGl) SetMouseButtonCallback(fn func(int, int)) { o.mouseClick = fn }
func (o *OpenGl) SetMouseWheelCallback(fn func(int))       { o.mouseWheel = fn }
func (o *OpenGl) SetCallbacks(disp3 func(), disp2 func(), keys func(int, int), bttn func(int, int), move func(int, int), scrl func(int)) {
	o.Set3dCallback(disp3)
	o.Set2dCallback(disp2)
	o.SetKeyboardCallback(keys)
	o.SetMousePosCallback(move)
	o.SetMouseButtonCallback(bttn)
	o.SetMouseWheelCallback(scrl)
}

func (o *OpenGl) Height() int { return o.height }
func (o *OpenGl) Width() int  { return o.width }

func New(title string, width, height int) *OpenGl {
	o := &OpenGl{
		title:  title,
		width:  width,
		height: height,
	}
	o.Camera = camera.New()
	return o
}

func (o *OpenGl) Init() {
	if err := glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return
	}
	defer glfw.Terminate()

	if err := glfw.OpenWindow(o.width, o.height, 0, 0, 0, 0, 16, 0, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr, "glfw: %s\n", err)
		return
	}
	defer glfw.CloseWindow()

	glfw.SetSwapInterval(1)
	glfw.SetWindowTitle(o.title)

	if err := gl.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "gl: %s\n", err)
	}

	glfw.SetKeyCallback(glfw.KeyHandler(o.keyboard))
	glfw.SetMouseButtonCallback(glfw.MouseButtonHandler(o.mouseClick))
	glfw.SetMousePosCallback(glfw.MousePosHandler(o.mouseMove))
	glfw.SetMouseWheelCallback(glfw.MouseWheelHandler(o.mouseWheel))

	if err := o.initScene(); err != nil {
		fmt.Fprintf(os.Stderr, "init: %s\n", err)
		return
	}

	for glfw.WindowParam(glfw.Opened) == 1 {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		if o.display3d != nil {
			o.set3dView()
			o.display3d()
		}

		if o.display2d != nil {
			o.set2dView()
			o.display2d()
		}

		glfw.SwapBuffers()
		time.Sleep(20 * time.Millisecond)
		//<-o.redisplay
	}
}

func (o *OpenGl) PostRedisplay() {
	o.redisplay <- true
}

func (o *OpenGl) initScene() (err error) {
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.ClearDepth(1)
	gl.DepthFunc(gl.LEQUAL)

	ambient := []gl.Float{0.5, 0.5, 0.5, 1}
	diffuse := []gl.Float{1, 1, 1, 1}
	position := []gl.Float{-5, 5, 10, 0}
	gl.Lightfv(gl.LIGHT0, gl.AMBIENT, &ambient[0])
	gl.Lightfv(gl.LIGHT0, gl.DIFFUSE, &diffuse[0])
	gl.Lightfv(gl.LIGHT0, gl.POSITION, &position[0])

	gl.Viewport(0, 0, gl.Sizei(o.width), gl.Sizei(o.height))
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Frustum(-1, 1, -1, 1, 1.0, 10.0)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()

	return nil
}

func (o *OpenGl) set3dView() {
	gl.Enable(gl.TEXTURE_2D)
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.LIGHTING)
	gl.Enable(gl.LIGHT0)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	//gl.Perspective(cam.Zoom(), o.mW / o.mH, cam.Near(), cam.Far())
	gl.MatrixMode(gl.MODELVIEW)
	//e := o.Camera.EYE()
	//a := o.Camera.AT()
	//u := o.Camera.UP()
	//gl.LookAt(e.X, e.Y, e.Z, a.X, a.Y, a.Z, u.X, u.Y, u.Z)
}

func (o *OpenGl) set2dView() {
	gl.Disable(gl.TEXTURE_2D)
	gl.Disable(gl.DEPTH_TEST)
	gl.Disable(gl.LIGHTING)
	gl.Disable(gl.LIGHT0)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, gl.Double(o.width), 0, gl.Double(o.height), -1, 1)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
}
