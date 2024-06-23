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

	anim Animation
}

func NewGhost(props *Properties) *Ghost {
	return &Ghost{
		Character: *NewCharacter(props),
		anim:      *NewAnimation(100, 6, sdl.FLIP_NONE, "ghost"),
	}
}

func (g Ghost) Draw() {
	transform := g.GetTransform()
	g.anim.Draw(transform.Position.X, transform.Position.Y, 32, 32)
}

func (g *Ghost) Update(dt float64) {
	g.anim.Update(dt)
}

func (g Ghost) Destroy() {
	g.anim.Destroy()
}
