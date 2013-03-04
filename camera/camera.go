package camera

import (
	"linAlg"
	//"math"
)

type Camera struct {
	eye, at                     linAlg.Point
	up, right                   linAlg.Vector
	dx, dy                      int
	near, far, zoom  float64
}

func New(/*eye, at linAlg.Point*/) *Camera {
	return &Camera{
		eye:   linAlg.Point{X: 0, Y: -5, Z: 0},
		at:    linAlg.Point{X: 0, Y: -0, Z: 0},
		up:    linAlg.Vector{X: 0, Y: 0, Z: 1},
		right: linAlg.Vector{X: 1, Y: 0, Z: 0},
		zoom:  40,
	}
}

// Position of the EYE or camera
func (c *Camera) Eye() linAlg.Point {
	return c.eye
}

// Point AT which the camera is looking
func (c *Camera) At() linAlg.Point {
	return c.at
}

// The UP vector of the camera
func (c *Camera) Up() linAlg.Vector {
	return c.up
}

// The Right vector of the camera
func (c *Camera) Right() linAlg.Vector {
	return c.right
}

// The Forward vector of the camera
func (c *Camera) Forward() linAlg.Vector {
	return c.at.Sub(c.eye)
}

func (c Camera) Zoom() float64 {
	return c.zoom
}

// The near cutoff plane of visibility as a unit of distance from the camera
func (c Camera) NearPlane() float64 {
	return c.near
}

// The far cutoff plane of visibility as a unit of distance from the camera
func (c Camera) FarPlane() float64 {
	return c.far
}

func (c *Camera) MoveCloser(amount float64) {
	l := c.Forward().Length()
	if l + amount >= 0 {
		c.eye = c.at.Add(c.Forward().Mul((l - amount) / l))
	}
}

func (c Camera) SpherePitch(amount float64) {
	c.Sphere(amount, c.right)
}

func (c Camera) SphereYaw(amount float64) {
	c.Sphere(amount, c.up)
}

func (c Camera) SphereRoll(amount float64) {
	c.Sphere(amount, c.Forward())
}

func (c *Camera) Sphere(amount float64, axis linAlg.Vector) {
	c.up = c.up.Rotate(amount, axis)
	c.right = c.right.Rotate(amount, axis)
	c.eye = c.at.Add(c.Forward().Neg().Rotate(amount, axis))
}

// Rotate up and down about the left to right axis
func (c Camera) PanPitch(amount float64) {
	c.Pan(amount, c.right)
}

func (c Camera) PanYaw(amount float64) {
	c.Pan(amount, c.up)
}

func (c Camera) PanRoll(amount float64) {
	c.Pan(amount, c.Forward())
}

func (c *Camera) Pan(radians float64, axis linAlg.Vector) {
	c.up = c.up.Rotate(radians, axis)
	c.right = c.right.Rotate(radians, axis)
	forward := c.Forward().Rotate(radians, axis)
	c.at = c.eye.Add(forward)
}

func (c Camera) SlideUp(amount float64) {
	c.Slide(c.up.Mul(amount))
}

func (c Camera) SlideDown(amount float64) {
	c.Slide(c.up.Mul(-amount))
}

func (c Camera) SlideRight(amount float64) {
	c.Slide(c.right.Mul(amount))
}

func (c Camera) SlideForward(amount float64) {
	c.Slide(c.Forward().Unit().Mul(amount))
}

func (c *Camera) Slide(v linAlg.Vector) {
	c.eye = c.eye.Add(v)
	c.at = c.at.Add(v)
}

func (c *Camera) Click(x, y int) {
	c.dx, c.dy = x, y
}

func (c *Camera) Drag(x, y int) {
	c.SpherePitch(float64(c.dy - y))
	c.SphereYaw(float64(c.dx - x))
	c.dx, c.dy = x, y
}
