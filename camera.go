package main

import (
	"fmt"
	"go-game/phy"

	"github.com/veandco/go-sdl2/sdl"
)

type Camera struct {
	instance *Camera
	target   *phy.Point
	position *phy.Vector
	viewBox  *sdl.Rect
}

func (c *Camera) GetInstance() *Camera {
	if c.instance == nil {
		c.instance = NewCamera()
	}

	return c.instance
}

func NewCamera() *Camera {
	return &Camera{
		target:   &phy.Point{X: 0, Y: 0},
		position: &phy.Vector{X: WIDTH, Y: HEIGHT},
		viewBox:  &sdl.Rect{},
	}
}

func (c *Camera) GetViewBox() *sdl.Rect {
	return c.viewBox
}

func (c *Camera) GetPosition() *phy.Vector {
	return c.position
}

func (c *Camera) SetTarget(t *phy.Point) {
	c.target = t
}

func (c *Camera) Update(dt float64) error {
	if c.target == nil {
		return fmt.Errorf("target does not exist")
	}

	c.viewBox.X = Limit(0, int32(c.target.X)-WIDTH/2, 2*WIDTH-c.viewBox.W)
	c.viewBox.Y = Limit(0, int32(c.target.Y)-HEIGHT/2, HEIGHT-c.viewBox.H)

	c.position = &phy.Vector{X: float64(c.viewBox.X), Y: float64(c.viewBox.Y)}

	return nil
}