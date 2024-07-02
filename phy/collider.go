package phy

import "github.com/veandco/go-sdl2/sdl"

type Collider struct {
	buffer *sdl.Rect
	box    *sdl.Rect
}

func (c *Collider) Get() *sdl.Rect {
	return c.box
}

func (c *Collider) Set(X, Y, W, H int32) {
	c.box = &sdl.Rect{
		X: X - c.buffer.X,
		Y: Y - c.buffer.Y,
		W: W - c.buffer.W,
		H: H - c.buffer.H,
	}
}

func (c *Collider) SetBuffer(X, Y, W, H int32) {
	c.buffer = &sdl.Rect{X: X, Y: Y, W: W, H: H}
}
