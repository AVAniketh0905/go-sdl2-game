package main

import (
	"fmt"
	"go-game/phy"

	"github.com/veandco/go-sdl2/sdl"
)

type Camera struct {
	instance    *Camera
	target      *phy.Point
	position    *phy.Vector
	viewBox     *sdl.Rect
	levelWidth  int32
	levelHeight int32
}

func (c *Camera) GetInstance() *Camera {
	if c.instance == nil {
		c.instance = NewCamera()
	}

	return c.instance
}

func NewCamera() *Camera {
	return &Camera{
		target:      &phy.Point{X: 0, Y: 0},
		position:    &phy.Vector{X: WIDTH, Y: HEIGHT},
		viewBox:     &sdl.Rect{},
		levelWidth:  WIDTH,
		levelHeight: HEIGHT,
	}
}

func (c *Camera) GetViewBox() *sdl.Rect {
	return c.viewBox
}

func (c *Camera) GetPosition() *phy.Vector {
	return c.position
}

func (c *Camera) GetLevelWidth() int32 {
	return c.levelWidth
}

func (c *Camera) GetLevelHeight() int32 {
	return c.levelHeight
}

func (c *Camera) SetTarget(t *phy.Point) {
	c.target = t
}

func (c *Camera) SetLevelLimit(w, h int32) {
	c.levelWidth = w
	c.levelHeight = h
}

func (c *Camera) SyncObject(pos *phy.Point, scrollRatio int) phy.Point {
	return phy.Point{X: pos.X - c.position.X*float64(scrollRatio), Y: pos.Y - c.position.Y*float64(scrollRatio)}
}

func (c *Camera) Update(dt float64) error {
	if c.target == nil {
		return fmt.Errorf("target does not exist")
	}

	c.viewBox.X = Limit(0, int32(c.target.X)-c.levelWidth/2, 2*c.levelWidth-c.viewBox.W)
	c.viewBox.Y = Limit(0, int32(c.target.Y)-2*c.levelHeight/3, c.levelHeight-c.viewBox.H)

	c.position = &phy.Vector{X: float64(c.viewBox.X), Y: float64(c.viewBox.Y)}

	return nil
}
