package main

import (
	"fmt"
	"go-game/phy"
)

type Enemy struct {
	Character

	anim *SeqAnimation
	rb   *phy.RigidBody

	collider *phy.Collider

	LastSafePosition *phy.Vector
}

func NewEnemy(props *Properties, repeat bool, path, seqId string) (*Enemy, error) {
	collider := &phy.Collider{}
	collider.SetBuffer(-5, -10, 0, 0)

	anim, err := NewSeqAnimation(repeat, path, seqId)
	if err != nil {
		return nil, fmt.Errorf("failed to load seq animation, %v", err)
	}

	return &Enemy{
		Character:        *NewCharacter(props),
		anim:             anim,
		rb:               phy.NewRigidBody(props.transform),
		collider:         collider,
		LastSafePosition: &phy.Vector{},
	}, nil
}

func (e *Enemy) Draw() {
	pos := e.transform
	e.anim.Draw(int(pos.X), int(pos.Y), 1, 1, e.flip)
}

func (e *Enemy) Update(dt float64) {
	e.rb.UnsetForces()
	e.rb.Update(dt)

	disp := e.rb.GetDisplacement()

	e.LastSafePosition.Set(phy.Vector{X: e.GetTransform().X, Y: e.LastSafePosition.Y})
	e.transform.TranslateX(disp.X)
	e.collider.Set(int32(e.transform.X), int32(e.transform.Y), 2*TILE_SIZE, 2*TILE_SIZE)

	if CollisionHandlerInstance.GetInstance().MapCollision(e.collider.Get()) {
		e.transform.Set(phy.Vector{X: e.LastSafePosition.X, Y: e.transform.Y})
	}

	e.LastSafePosition.Set(phy.Vector{X: e.LastSafePosition.Y, Y: e.GetTransform().Y})
	e.transform.TranslateY(disp.Y)
	e.collider.Set(int32(e.transform.X), int32(e.transform.Y), TILE_SIZE, TILE_SIZE)

	if CollisionHandlerInstance.GetInstance().MapCollision(e.collider.Get()) {
		e.transform.Set(phy.Vector{X: e.transform.X, Y: e.LastSafePosition.Y})
	}

	e.anim.Update(dt)
	if e.anim.IsEnded() {
		e.anim.SetCurrentSeq("player_idle")
		e.anim.SetRepeat(true)
	}
}
