package main

import (
	"go-game/phy"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

var DEFAULT_PROPS PublicAnim = PublicAnim{
	animSpeed:  80,
	frameCount: 6,
	texId:      "ghost",
	flip:       sdl.FLIP_NONE,
}

var RUNNING_R_PROPS PublicAnim = PublicAnim{
	animSpeed:  80,
	frameCount: 3,
	texId:      "ghost_run",
	flip:       sdl.FLIP_HORIZONTAL,
}

var RUNNING_L_PROPS PublicAnim = PublicAnim{
	animSpeed:  80,
	frameCount: 3,
	texId:      "ghost_run",
	flip:       sdl.FLIP_NONE,
}

type Ghost struct {
	Character

	anim *Animation
	rb   *phy.RigidBody
}

func NewGhost(props *Properties) *Ghost {
	return &Ghost{
		Character: *NewCharacter(props),
		anim:      NewAnimation(DEFAULT_PROPS),
		rb:        phy.NewRigidBody(props.transform),
	}
}

func (g *Ghost) Draw() {
	log.Println("Drawing Ghost", g.transform, &g.transform)
	transform := g.GetTransform()
	g.anim.Draw(int(transform.X), int(transform.Y), IMG_SIZE, IMG_SIZE)
}

func (g *Ghost) Controls() {
	if InputInstance.GetInstance().IsKeyDown(sdl.SCANCODE_A) {
		g.rb.AddForce(phy.Vector{X: -1, Y: 0})
		g.anim.SetProps(RUNNING_R_PROPS)
	}

	if InputInstance.GetInstance().IsKeyDown(sdl.SCANCODE_D) {
		g.rb.AddForce(phy.Vector{X: 1, Y: 0})
		g.anim.SetProps(RUNNING_L_PROPS)

	}
}

func (g *Ghost) Update(dt float64) {
	g.anim.SetProps(DEFAULT_PROPS)

	g.Controls()
	g.rb.Update(dt)

	disp := g.rb.GetDisplacement()
	log.Println(disp, g.transform)
	g.transform.Translate(disp)
	log.Println(g.transform)

	g.anim.Update(dt)
	log.Println("Updating Ghost", g.transform, &g.transform)
}

func (g Ghost) Destroy() {
	g.anim.Destroy()
}
