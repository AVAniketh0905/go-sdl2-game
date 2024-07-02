package main

import (
	"go-game/phy"

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

	isJumping  bool
	isGrounded bool

	jumpTime  float64
	jumpForce float64

	collider *phy.Collider

	LastSafePosition *phy.Vector
}

func NewGhost(props *Properties) *Ghost {
	collider := &phy.Collider{}
	collider.SetBuffer(0, 0, 0, 0)
	return &Ghost{
		Character:        *NewCharacter(props),
		anim:             NewAnimation(DEFAULT_PROPS),
		rb:               phy.NewRigidBody(props.transform),
		collider:         collider,
		jumpTime:         JUMP_TIME,
		jumpForce:        JUMP_FORCE,
		LastSafePosition: &phy.Vector{},
	}
}

func (g *Ghost) updateOrigin() {
	g.origin.X = g.transform.X + float64(g.width)/2
	g.origin.Y = g.transform.Y + float64(g.height)/2
}

func (g *Ghost) Draw() {
	transform := g.GetTransform()
	g.anim.Draw(int(transform.X), int(transform.Y), IMG_SIZE, IMG_SIZE)

	cam := CameraInstance.GetInstance().GetPosition()
	box := g.collider.Get()
	box.X -= int32(cam.X)
	box.Y -= int32(cam.Y)
	EngineInstance.GetInstance().GetRenderer().DrawRect(box)
}

func (g *Ghost) Controls(dt float64) {
	if InputInstance.GetInstance().IsKeyDown(sdl.SCANCODE_A) {
		g.rb.AddForce(phy.Vector{X: 15, Y: 0})
		g.anim.SetProps(RUNNING_R_PROPS)
	}

	if InputInstance.GetInstance().IsKeyDown(sdl.SCANCODE_D) {
		g.rb.AddForce(phy.Vector{X: -5, Y: 0})
		g.anim.SetProps(RUNNING_L_PROPS)
	}

	g.JumpControls(dt)
}

func (g *Ghost) JumpControls(dt float64) {
	if InputInstance.GetInstance().IsKeyDown(sdl.SCANCODE_W) && g.isGrounded {
		g.isJumping = true
		g.isGrounded = false
		g.rb.AddForce(phy.Vector{X: 0, Y: g.jumpForce})
	}

	if InputInstance.GetInstance().IsKeyDown(sdl.SCANCODE_W) && g.isJumping && g.jumpTime > 0 {
		g.jumpTime -= dt
		g.rb.AddForce(phy.Vector{X: 0, Y: g.jumpForce})
	} else {
		g.isJumping = false
		g.jumpTime = JUMP_TIME
	}
}

func (g *Ghost) Update(dt float64) {
	g.anim.SetProps(DEFAULT_PROPS)
	g.rb.UnsetForces()

	g.Controls(dt)
	g.rb.Update(dt)

	disp := g.rb.GetDisplacement()

	g.LastSafePosition.Set(phy.Vector{X: g.GetTransform().X, Y: g.LastSafePosition.Y})
	g.transform.TranslateX(disp.X)
	g.collider.Set(int32(g.transform.X), int32(g.transform.Y), 96, 96)

	if CollisionHandlerInstance.GetInstance().MapCollision(g.collider.Get()) {
		g.transform.Set(phy.Vector{X: g.LastSafePosition.X, Y: g.transform.Y})
	}

	g.LastSafePosition.Set(phy.Vector{X: g.LastSafePosition.X, Y: g.GetTransform().Y})
	g.transform.TranslateY(disp.Y)
	g.collider.Set(int32(g.transform.X), int32(g.transform.Y), 96, 96)

	if CollisionHandlerInstance.GetInstance().MapCollision(g.collider.Get()) {
		g.isGrounded = true
		g.transform.Set(phy.Vector{X: g.transform.X, Y: g.LastSafePosition.Y})
	} else {
		g.isGrounded = false
	}

	g.updateOrigin()
	g.anim.Update(dt)
}

func (g Ghost) Destroy() {
	g.anim.Destroy()
}
