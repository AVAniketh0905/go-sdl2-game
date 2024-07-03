package main

import (
	"fmt"
	"go-game/phy"

	"github.com/veandco/go-sdl2/sdl"
)

type Object interface {
	Draw()
	Update(dt float64)
	Destroy()
}

type Properties struct {
	transform *phy.Transform
	width     int
	height    int
	texId     string
	flip      sdl.RendererFlip
}

func (p *Properties) String() string {
	return fmt.Sprint("Transform: ", p.transform, " Width: ", p.width, " Height: ", p.height, " Texture ID: ", p.texId, " Flip: ", p.flip)
}

type GameObject struct {
	Object     // Embedded
	Properties // From Properties
	origin     *phy.Point
}

func NewGameObject(props *Properties) *GameObject {
	px := props.transform.X + float64(props.width)/2
	py := props.transform.Y + float64(props.height)/2

	return &GameObject{
		Properties: *props,
		origin:     &phy.Point{X: px, Y: py},
	}
}

func (g *GameObject) GetTransform() *phy.Transform {
	return g.transform
}

func (g *GameObject) GetOrigin() *phy.Point {
	return g.origin
}

func (g GameObject) Draw() {
}

func (g GameObject) Update(dt float64) {
}

func (g GameObject) Destroy() {
}

// Character
type Character struct {
	GameObject

	Name string
}

func NewCharacter(props *Properties) *Character {
	return &Character{
		GameObject: *NewGameObject(props),
	}
}

func (c Character) Draw() {
}

func (c Character) Update(dt float64) {
}

func (c Character) Destroy() {
}
