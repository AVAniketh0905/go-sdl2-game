package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Character struct {
	GameObject

	Name string
}

func NewCharacter(props *Properties) *Character {
	return &Character{
		GameObject: *NewGameObject(props),
	}
}

type Ghost struct {
	Character

	AnimSpeed  int
	Row        int
	Frame      int
	FrameCount int
}

func NewGhost(props *Properties) *Ghost {
	return &Ghost{
		Character:  *NewCharacter(props),
		Row:        0,
		Frame:      0,
		FrameCount: 6,
		AnimSpeed:  100,
	}
}

func (g Ghost) Draw() {
	transform := g.GetTransform()
	TextureManagerInstance.DrawFrame(g.texId, transform.Position.X, transform.Position.Y, g.width, g.height, g.Row, g.Frame, sdl.FLIP_NONE)
}

func (g *Ghost) Update(dt float64) {
	time := sdl.GetTicks64()
	// always have a pointer reference
	g.Frame = int(int(time)/(g.AnimSpeed)) % g.FrameCount
}

func (g Ghost) Destroy() {
	TextureManagerInstance.Destroy()
}
