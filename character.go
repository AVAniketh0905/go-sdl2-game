package main

import (
	"go-game/phy"
	"log"

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
	rb   *phy.RigidBody
}

func NewGhost(props *Properties) *Ghost {
	return &Ghost{
		Character: *NewCharacter(props),
		anim:      *NewAnimation(100, 6, sdl.FLIP_NONE, "ghost"),
		rb:        phy.NewRigidBody(props.transform.Position),
	}
}

func (g *Ghost) Draw() {
	log.Println("Drawing Ghost", g.transform.Position)
	transform := g.GetTransform()
	g.anim.Draw(int(transform.Position.X), int(transform.Position.Y), IMG_SIZE, IMG_SIZE)
}

func (g *Ghost) Update(dt float64) {
	g.rb.Update(0.1)
	pos := g.rb.GetPosition()
	g.transform.Translate(pos)
	g.anim.Update(dt)
}

func (g Ghost) Destroy() {
	g.anim.Destroy()
}
