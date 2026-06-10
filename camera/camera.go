package camera

import "github.com/go-gl/mathgl/mgl32"

type CameraMovement int

const (
	Forward CameraMovement = iota
	Backward
	Left
	Right
)

type Camera struct {
	position    mgl32.Vec3
	front       mgl32.Vec3
	up          mgl32.Vec3
	yaw         float32
	pitch       float32
	fov         float32
	speed       float32
	sensitivity float32
	lastMouseX  float32
	lastMouseY  float32
	firstMouse  bool
}

func New() *Camera {
	return &Camera{
		position:    mgl32.Vec3{0, 0, 3},
		front:       mgl32.Vec3{0, 0, -1},
		up:          mgl32.Vec3{0, 1, 0},
		yaw:         -90,
		pitch:       0,
		fov:         45,
		speed:       2.5,
		sensitivity: 0.1,
		lastMouseX:  400,
		lastMouseY:  300,
		firstMouse:  true,
	}
}

func (c *Camera) GetFov() float32 {
	return c.fov
}

func (c *Camera) GetViewMatrix() mgl32.Mat4 {
	return mgl32.LookAtV(c.position, c.position.Add(c.front), c.up)
}

func (c *Camera) ProcessKeyboard(dir CameraMovement, deltatime float32) {
	speed := c.speed * deltatime

	switch dir {
	case Forward:
		c.position = c.position.Add(c.front.Mul(speed))
	case Backward:
		c.position = c.position.Sub(c.front.Mul(speed))
	case Left:
		c.position = c.position.Sub(c.front.Cross(c.up).Normalize().Mul(speed))
	case Right:
		c.position = c.position.Add(c.front.Cross(c.up).Normalize().Mul(speed))
	}
}

func (c *Camera) recomputeFront() {
	c.front = mgl32.Vec3{
		cos32(mgl32.DegToRad(c.yaw)) * cos32(mgl32.DegToRad(c.pitch)),
		sin32(mgl32.DegToRad(c.pitch)),
		sin32(mgl32.DegToRad(c.yaw)) * cos32(mgl32.DegToRad(c.pitch)),
	}.Normalize()
}

func (c *Camera) ProcessMouseMovement(x, y float32) {
	if c.firstMouse {
		c.lastMouseX, c.lastMouseY = x, y
		c.firstMouse = false
		return
	}

	dx := (x - c.lastMouseX) * c.sensitivity
	dy := (c.lastMouseY - y) * c.sensitivity
	c.lastMouseX, c.lastMouseY = x, y

	c.yaw += dx
	c.pitch += dy

	// Clamp pitch
	if c.pitch > 89.0 {
		c.pitch = 89.0
	}
	if c.pitch < -89.0 {
		c.pitch = -89.0
	}

	c.recomputeFront()
}

func (c *Camera) ProcessMouseScroll(yoff float32) {
	c.fov -= float32(yoff)
	if c.fov > 45 {
		c.fov = 45
	}
	if c.fov < 1 {
		c.fov = 1
	}
}
