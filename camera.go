package main

import (
	"fmt"
	"go-game/phy"

	"github.com/veandco/go-sdl2/sdl"
)

const CAM_OFFSET = 200

type Camera struct {
	instance    *Camera
	target      *phy.Point
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
		target:  &phy.Point{X: 0, Y: 0},
		viewBox: &sdl.Rect{},
	}
}

func (c *Camera) GetViewBox() *sdl.Rect {
	return c.viewBox
}

func (c *Camera) GetLevelWidth() int32 {
	return c.levelWidth
}

func (c *Camera) GetLevelHeight() int32 {
	return c.levelHeight
}

func (c *Camera) GetTarget() *phy.Point {
	return c.target
}

func (c *Camera) SetTarget(t *phy.Point) {
	c.target = t
}

func (c *Camera) SetLevelLimit(w, h int32) {
	c.levelWidth = w
	c.levelHeight = h
}

func (c *Camera) SetViewBox(x, y, w, h int32) {
	c.viewBox.X = x
	c.viewBox.Y = y
	c.viewBox.W = w
	c.viewBox.H = h
}

func (c *Camera) SyncObject(pos *phy.Point, scrollRatio int) phy.Point {
	return phy.Point{X: pos.X - float64(c.viewBox.X)*float64(scrollRatio), Y: pos.Y - float64(c.viewBox.Y)*float64(scrollRatio)}
}

func (c *Camera) IsInside() bool {
	return c.target.X > float64(c.viewBox.X) && c.target.X < float64(c.viewBox.X+c.viewBox.W) &&
		c.target.Y > float64(c.viewBox.Y) && c.target.Y < float64(c.viewBox.Y+c.viewBox.H)
}

func (c *Camera) Draw() {
	r := EngineInstance.GetInstance().GetRenderer()
	r.SetDrawColor(255, 255, 255, 255)
	r.DrawRect(&sdl.Rect{X: 0, Y: 0, W: c.viewBox.W, H: c.viewBox.H})
	r.SetDrawColor(0, 0, 0, 255)
}

func (c *Camera) Update(dt uint64) error {
	if c.target == nil {
		return fmt.Errorf("target does not exist")
	}

	offSetX := c.viewBox.W / 2
	offSetY := c.viewBox.H/2 + CAM_OFFSET

	c.viewBox.X = Limit(0, int32(c.target.X)-offSetX, c.viewBox.W)
	c.viewBox.Y = Limit(0, int32(c.target.Y)-offSetY, c.viewBox.H)
	return nil
}
