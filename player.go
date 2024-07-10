package main

import (
	"go-game/phy"

	"github.com/veandco/go-sdl2/sdl"
)

type PState int

const (
	IDLE PState = iota
	RUN
	JUMP
	FALL
	CROUCH
	ATTACK
	DAMAGE
	DEATH
)

type Player struct {
	Character

	anim *SpriteAnimation
	rb   *phy.RigidBody

	// isJumping   bool
	// isFalling   bool
	canJump    bool
	isGrounded bool
	// isRunning   bool
	// isAttacking bool
	// isCrouching bool
	state PState

	jumpHeight int
	jumpForce  float64

	attackTime float64

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
		state:            IDLE,
		canJump:          true,
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
	switch p.state {
	case IDLE:
		p.anim.SetProps("player_idle", 0, 2, 80) // DEFAULT PROPS
	case RUN:
		p.anim.SetProps("player_walk", 0, 4, 80) // RUNNING PROPS
	case JUMP:
		p.anim.SetProps("player_walk", 0, 4, 80)
	case FALL:
		p.anim.SetProps("player_damage", 0, 2, 80)
	case CROUCH:
		p.anim.SetProps("player_death", 0, 4, 150)
	case ATTACK:
		p.anim.SetProps("player_attack", 0, 4, 50)
	case DAMAGE:
	case DEATH:
	}
}

func (p Player) Draw() {
	transform := p.GetTransform()
	p.anim.Draw(int(transform.X), int(transform.Y), IMG_SIZE, IMG_SIZE, 1, 1)

	cam := CameraInstance.GetInstance().GetViewBox()
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
	if InputInstance.GetInstance().GetAxisKey(HORIZONTAL) == -1 && p.state != ATTACK {
		p.rb.AddForce(phy.Vector{X: RUN_FORCE, Y: 0})
		p.anim.SetFlip(sdl.FLIP_HORIZONTAL)
		p.state = RUN
	}

	if InputInstance.GetInstance().GetAxisKey(HORIZONTAL) == 1 && p.state != ATTACK {
		p.rb.AddForce(phy.Vector{X: -RUN_FORCE, Y: 0})
		p.anim.SetFlip(sdl.FLIP_NONE)
		p.state = RUN
	}
}

func (p *Player) JumpControls(dt float64) {
	if InputInstance.GetInstance().GetAxisKey(VERTICAL) == 1 && p.isGrounded {
		p.canJump = false
		p.isGrounded = false
		p.jumpHeight = 0
		p.state = JUMP
		p.rb.AddForce(phy.Vector{X: 0, Y: p.jumpForce})
	}

	if InputInstance.GetInstance().GetAxisKey(VERTICAL) == 1 && !p.isGrounded && p.jumpHeight < MAX_JUMP_HEIGHT {
		p.jumpHeight++
		p.state = JUMP
		p.rb.AddForce(phy.Vector{X: 0, Y: p.jumpForce})
	}
}

func (p *Player) CrouchControls(dt float64) {
	if InputInstance.GetInstance().GetAxisKey(VERTICAL) == -1 {
		p.rb.UnsetForces()
		p.state = CROUCH
	}

	if p.rb.GetVelocity().Y > 0 && !p.isGrounded {
		p.state = FALL
	}
}

func (p *Player) AttackControls(dt float64) {
	if InputInstance.GetInstance().IsKeyDown(sdl.SCANCODE_SPACE) {
		p.rb.UnsetForces()
		p.state = ATTACK
	}

	if p.state == ATTACK && p.attackTime > 0 {
		p.attackTime -= dt
	} else {
		p.attackTime = ATTACK_TIME
	}
}

func (p *Player) Update(dt float64) {
	p.state = IDLE
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

	p.LastSafePosition.Set(phy.Vector{X: p.GetTransform().X, Y: p.LastSafePosition.Y})
	if p.state == CROUCH {
		p.collider.Set(int32(p.transform.X), int32(p.transform.Y)+TILE_SIZE/2, TILE_SIZE, 2*TILE_SIZE)
		if CollisionHandlerInstance.GetInstance().MapCollision(p.collider.Get()) {
			p.transform.Set(phy.Vector{X: p.transform.X, Y: p.LastSafePosition.Y})
		}
	}

	p.updateOrigin()
	p.anim.Update(dt)
}

func (p Player) Destroy() {
	p.anim.Destroy()
}
