package main

import (
	"go-game/phy"

	"github.com/veandco/go-sdl2/sdl"
)

type Ghost struct {
	Character

	anim *SpriteAnimation
	rb   *phy.RigidBody

	isJumping   bool
	isFalling   bool
	isGrounded  bool
	isRunning   bool
	isAttacking bool
	isCrouching bool

	jumpTime   float64
	attackTime float64

	jumpForce float64

	collider *phy.Collider

	LastSafePosition *phy.Vector
}

func NewGhost(props *Properties) *Ghost {
	collider := &phy.Collider{}
	collider.SetBuffer(-30, 0, 0, 0)
	return &Ghost{
		Character:        *NewCharacter(props),
		anim:             NewSpriteAnimation(props.texId, 6, 80, sdl.FLIP_NONE),
		rb:               phy.NewRigidBody(props.transform),
		collider:         collider,
		jumpTime:         JUMP_TIME,
		jumpForce:        JUMP_FORCE,
		attackTime:       ATTACK_TIME,
		LastSafePosition: &phy.Vector{},
	}
}

func (g *Ghost) updateOrigin() {
	g.origin.X = g.transform.X + float64(g.width)/2
	g.origin.Y = g.transform.Y + float64(g.height)/2
}

func (g *Ghost) animationState() {
	g.anim.SetProps("player_idle", 0, 2, 80) // DEFAULT PROPS

	if g.isRunning {
		g.anim.SetProps("player_walk", 0, 4, 80) // RUNNING PROPS
	}

	if g.isJumping {
		// g.anim.SetProps(JUMPING_PROPS)
		g.anim.SetProps("player_walk", 0, 4, 80)
	}

	if g.isFalling {
		// g.anim.SetProps(FALLING_PROPS)
		g.anim.SetProps("player_damage", 0, 2, 80)
	}

	if g.isAttacking {
		// g.anim.SetProps(ATTACK_PROPS)
		g.anim.SetProps("player_attack", 0, 4, 50)
	}

	if g.isCrouching {
		// g.anim.SetProps(CROUCH_PROPS)
		g.anim.SetProps("player_death", 0, 4, 150)
	}
}

func (g Ghost) Draw() {
	transform := g.GetTransform()
	g.anim.Draw(int(transform.X), int(transform.Y), IMG_SIZE, IMG_SIZE, 1, 1)

	cam := CameraInstance.GetInstance().GetPosition()
	box := g.collider.Get()
	box.X -= int32(cam.X)
	box.Y -= int32(cam.Y)
	EngineInstance.GetInstance().GetRenderer().DrawRect(box)
}

func (g *Ghost) Controls(dt float64) {
	g.RunningControls(dt)
	g.JumpControls(dt)
	g.CrouchControls(dt)
	g.AttackControls(dt)
}

func (g *Ghost) RunningControls(dt float64) {
	if InputInstance.GetInstance().GetAxisKey(HORIZONTAL) == -1 && !g.isAttacking {
		g.rb.AddForce(phy.Vector{X: RUN_FORCE, Y: 0})
		g.anim.SetFlip(sdl.FLIP_HORIZONTAL)
		g.isRunning = true
	}

	if InputInstance.GetInstance().GetAxisKey(HORIZONTAL) == 1 && !g.isAttacking {
		g.rb.AddForce(phy.Vector{X: -RUN_FORCE, Y: 0})
		g.anim.SetFlip(sdl.FLIP_NONE)
		g.isRunning = true
	}
}

func (g *Ghost) JumpControls(dt float64) {
	if InputInstance.GetInstance().GetAxisKey(VERTICAL) == 1 && g.isGrounded {
		g.isJumping = true
		g.isGrounded = false
		g.rb.AddForce(phy.Vector{X: 0, Y: g.jumpForce})
	}

	if InputInstance.GetInstance().GetAxisKey(VERTICAL) == 1 && g.isJumping && g.jumpTime > 0 {
		g.jumpTime -= dt
		g.rb.AddForce(phy.Vector{X: 0, Y: g.jumpForce})
	} else {
		g.isJumping = false
		g.jumpTime = JUMP_TIME
	}

	if g.rb.GetVelocity().Y > 0 && !g.isGrounded {
		g.isFalling = true
	} else {
		g.isFalling = false
	}

	if g.isAttacking && g.attackTime > 0 {
		g.attackTime -= dt
	} else {
		g.isAttacking = false
		g.attackTime = ATTACK_TIME
	}
}

func (g *Ghost) CrouchControls(dt float64) {
	if InputInstance.GetInstance().GetAxisKey(VERTICAL) == -1 {
		g.rb.UnsetForces()
		g.isCrouching = true
	}
}

func (g *Ghost) AttackControls(dt float64) {
	if InputInstance.GetInstance().IsKeyDown(sdl.SCANCODE_SPACE) {
		g.rb.UnsetForces()
		g.isAttacking = true
	}
}

func (g *Ghost) Update(dt float64) {
	g.isRunning = false
	g.isCrouching = false
	g.rb.UnsetForces()

	g.Controls(dt)
	g.animationState()

	g.rb.Update(dt)

	disp := g.rb.GetDisplacement()

	g.LastSafePosition.Set(phy.Vector{X: g.GetTransform().X, Y: g.LastSafePosition.Y})
	g.transform.TranslateX(disp.X)
	g.collider.Set(int32(g.transform.X), int32(g.transform.Y), 2*TILE_SIZE, 2*TILE_SIZE)

	if CollisionHandlerInstance.GetInstance().MapCollision(g.collider.Get()) {
		g.transform.Set(phy.Vector{X: g.LastSafePosition.X, Y: g.transform.Y})
	}

	g.LastSafePosition.Set(phy.Vector{X: g.LastSafePosition.X, Y: g.GetTransform().Y})
	g.transform.TranslateY(disp.Y)
	g.collider.Set(int32(g.transform.X), int32(g.transform.Y), TILE_SIZE, 2*TILE_SIZE)

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
