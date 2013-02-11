package camera

import (
	"linAlg"
	"math"
)

type Camera struct {
	mUp                                                  linAlg.Vector
	mEye, mAt                                            linAlg.Point
	mUpAngle, mPanAngle, mPanSpeed, mPanAmount, mDx, mDy float64
	mDragSpeed, mNearPlane, mFarPlane, mDistance, mZoom  float64
}

func New(/*eye, at linAlg.Point*/) *Camera {
	return &Camera{
		mEye: linAlg.Point{X: 0, Y: -5, Z: 0},
		mAt:  linAlg.Point{X: 0, Y: 0, Z: 0},
		mUp:  linAlg.Vector{X: 0, Y: 0, Z: 1},
	}
}

// Position of the EYE or camera
func (c *Camera) EYE() linAlg.Point {
	return c.mEye.Add(c.mAt)
}

// Point AT which the camera is looking
func (c *Camera) AT() linAlg.Point {
	return c.mAt
}

// The UP vector of the camera
func (c *Camera) UP() linAlg.Vector {
	return c.mUp
}

func (c *Camera) Zoom() float64 {
	return c.mZoom
}

// The near cutoff plane of visibility as a unit of distance from the camera
func (c *Camera) NearPlane() float64 {
	return c.mNearPlane
}

// The far cutoff plane of visibility as a unit of distance from the camera
func (c *Camera) FarPlane() float64 {
	return c.mFarPlane
}

func (c *Camera) Move() {
	c.mNearPlane = c.mDistance - 60
	c.mFarPlane = c.mDistance + 70
	if c.mNearPlane <= 0 {
		c.mNearPlane = 0.1
	}
	// Set camera position
	c.mEye = linAlg.SpherePoint(c.mDistance, c.mUpAngle, c.mPanAngle)
	// Right vector
	v := linAlg.Vector{
		X: math.Cos(c.mPanAngle + math.Pi/2),
		Y: math.Sin(c.mPanAngle + math.Pi/2),
		Z: 0,
	}
	// Forward vector
	w := linAlg.Vector(c.mAt.Sub(c.mEye))

	c.mUp = v.Cross(w)
	c.mUp = c.mUp.Unit()
	return
}

func (c *Camera) MoveIn() {
	c.mDistance -= 5
	c.Move()
}

func (c *Camera) MoveOut() {
	c.mDistance += 5
	c.Move()
}

func (c *Camera) MoveUp() {
	c.mUpAngle += c.mPanAmount
	c.Move()
}

func (c *Camera) MoveDown() {
	c.mUpAngle -= c.mPanAmount
	c.Move()
}

func (c *Camera) MoveLeft() {
	c.mPanAngle -= c.mPanAmount
	c.Move()
}

func (c *Camera) MoveRight() {
	c.mPanAngle += c.mPanAmount
	c.Move()
}

func (c *Camera) PanUp(amount float64) {
}

func (c *Camera) PanDown(amount float64) {
}

func (c *Camera) PanLeft(amount float64) {
}

func (c *Camera) PanRight(amount float64) {
}

func (c *Camera) SlideUp(amount float64) {
	c.Slide(linAlg.Vector{0, 0, amount})
}

func (c *Camera) SlideDown(amount float64) {
	c.Slide(linAlg.Vector{0, 0, -amount})
}

func (c *Camera) SlideLeft(amount float64) {
	c.Slide(linAlg.Vector{0, -amount, 0})
}

func (c *Camera) SlideRight(amount float64) {
	c.Slide(linAlg.Vector{0, amount, 0})
}

func (c *Camera) Slide(v linAlg.Vector) {
	c.mEye = c.mEye.Add(linAlg.Point(v))
	c.mAt = c.mAt.Add(linAlg.Point(v))
}

func (c *Camera) SetPan() {
	return
}

func (c *Camera) Pan() {
	c.mPanAngle += c.mPanAmount
	c.Move()
}

func (c *Camera) Click(x, y int) {
	c.mDx, c.mDy = float64(x), float64(y)
}

func (c *Camera) Drag(x, y int) {
	c.mPanAngle -= float64(x)-c.mDx / c.mDragSpeed
	c.mUpAngle -= float64(y)-c.mDy / c.mDragSpeed
	c.mDx, c.mDy = float64(x), float64(y)
	c.Move()
}
