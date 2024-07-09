package main

import (
	"go-game/phy"

	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
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

func NewPlayer(props *Properties) *Player {
	collider := &phy.Collider{}
	collider.SetBuffer(-30, -5, 0, 0)
	return &Player{
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

func (p *Player) updateOrigin() {
	p.origin.X = p.transform.X + float64(p.width)/2
	p.origin.Y = p.transform.Y + float64(p.height)/2
}

func (p *Player) animationState() {
	p.anim.SetProps("player_idle", 0, 2, 80) // DEFAULT PROPS

	if p.isRunning {
		p.anim.SetProps("player_walk", 0, 4, 80) // RUNNING PROPS
	}

	if p.isJumping {
		// p.anim.SetProps(JUMPING_PROPS)
		p.anim.SetProps("player_walk", 0, 4, 80)
	}

	if p.isFalling {
		// p.anim.SetProps(FALLING_PROPS)
		p.anim.SetProps("player_damage", 0, 2, 80)
	}

	if p.isAttacking {
		// p.anim.SetProps(ATTACK_PROPS)
		p.anim.SetProps("player_attack", 0, 4, 50)
	}

	if p.isCrouching {
		// p.anim.SetProps(CROUCH_PROPS)
		p.anim.SetProps("player_death", 0, 4, 150)
	}
}

func (p Player) Draw() {
	transform := p.GetTransform()
	p.anim.Draw(int(transform.X), int(transform.Y), IMG_SIZE, IMG_SIZE, 1, 1)

	cam := CameraInstance.GetInstance().GetPosition()
	// to account for initialization problems
	if p.collider.Get() != nil {
		box := p.collider.Get()
		box.X -= int32(cam.X)
		box.Y -= int32(cam.Y)
		EngineInstance.GetInstance().GetRenderer().DrawRect(box)
	}
}

func (p *Player) Controls(dt float64) {
	p.RunningControls(dt)
	p.JumpControls(dt)
	p.CrouchControls(dt)
	p.AttackControls(dt)
}

func (p *Player) RunningControls(dt float64) {
	if InputInstance.GetInstance().GetAxisKey(HORIZONTAL) == -1 && !p.isAttacking {
		p.rb.AddForce(phy.Vector{X: RUN_FORCE, Y: 0})
		p.anim.SetFlip(sdl.FLIP_HORIZONTAL)
		p.isRunning = true
	}

	if InputInstance.GetInstance().GetAxisKey(HORIZONTAL) == 1 && !p.isAttacking {
		p.rb.AddForce(phy.Vector{X: -RUN_FORCE, Y: 0})
		p.anim.SetFlip(sdl.FLIP_NONE)
		p.isRunning = true
	}
}

func (p *Player) JumpControls(dt float64) {
	if InputInstance.GetInstance().GetAxisKey(VERTICAL) == 1 && p.isGrounded {
		p.isJumping = true
		p.isGrounded = false
		p.rb.AddForce(phy.Vector{X: 0, Y: p.jumpForce})
	}

	if InputInstance.GetInstance().GetAxisKey(VERTICAL) == 1 && p.isJumping && p.jumpTime > 0 {
		p.jumpTime -= dt
		p.rb.AddForce(phy.Vector{X: 0, Y: p.jumpForce})
	} else {
		p.isJumping = false
		p.jumpTime = JUMP_TIME
	}

	if p.rb.GetVelocity().Y > 0 && !p.isGrounded {
		p.isFalling = true
	} else {
		p.isFalling = false
	}

	if p.isAttacking && p.attackTime > 0 {
		p.attackTime -= dt
	} else {
		p.isAttacking = false
		p.attackTime = ATTACK_TIME
	}
}

func (p *Player) CrouchControls(dt float64) {
	if InputInstance.GetInstance().GetAxisKey(VERTICAL) == -1 {
		p.rb.UnsetForces()
		p.isCrouching = true
	}
}

func (p *Player) AttackControls(dt float64) {
	if InputInstance.GetInstance().IsKeyDown(sdl.SCANCODE_SPACE) {
		p.rb.UnsetForces()
		p.isAttacking = true
	}
}

func (p *Player) Update(dt float64) {
	p.isRunning = false
	p.isCrouching = false
	p.rb.UnsetForces()

	p.Controls(dt)
	p.animationState()

	p.rb.Update(dt)

	disp := p.rb.GetDisplacement()

	p.LastSafePosition.Set(phy.Vector{X: p.GetTransform().X, Y: p.LastSafePosition.Y})
	p.transform.TranslateX(disp.X)
	p.collider.Set(int32(p.transform.X), int32(p.transform.Y), 2*TILE_SIZE, 2*TILE_SIZE)

	if CollisionHandlerInstance.GetInstance().MapCollision(p.collider.Get()) {
		p.transform.Set(phy.Vector{X: p.LastSafePosition.X, Y: p.transform.Y})
	}

	p.LastSafePosition.Set(phy.Vector{X: p.LastSafePosition.X, Y: p.GetTransform().Y})
	p.transform.TranslateY(disp.Y)
	p.collider.Set(int32(p.transform.X), int32(p.transform.Y), TILE_SIZE, 2*TILE_SIZE)

	if CollisionHandlerInstance.GetInstance().MapCollision(p.collider.Get()) {
		p.isGrounded = true
		p.transform.Set(phy.Vector{X: p.transform.X, Y: p.LastSafePosition.Y})
	} else {
		p.isGrounded = false
	}

	p.updateOrigin()

	p.anim.Update(dt)
}

func (p Player) Destroy() {
	p.anim.Destroy()
}
