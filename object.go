package main

import (
	"fmt"
	"go-game/phy"

	"github.com/veandco/go-sdl2/sdl"
)

type Object interface {
	Draw()
	Update()
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
}

func NewGameObject(props *Properties) *GameObject {
	return &GameObject{
		Properties: *props,
	}
}

func (g *GameObject) GetTransform() *phy.Transform {
	return g.transform
}

func (g GameObject) Draw() {
	println("Drawing")
}

func (g GameObject) Update() {
	println("Updating")
}

func (g GameObject) Destroy() {
	println("Destroying")
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
