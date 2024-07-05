package main

import (
	"go-game/phy"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Particle struct {
	transform *phy.Transform
	size      int32
	lifeTime  time.Duration
	vel       *phy.Vector
}

func NewParticle(transform *phy.Transform, size int32, vx, vy float64, liefTime time.Duration) *Particle {
	return &Particle{
		transform: transform,
		size:      size,
		lifeTime:  time.Duration(sdl.GetTicks64()) + liefTime,
		vel:       &phy.Vector{X: vx, Y: vy},
	}
}

func (p *Particle) Move(dt float64) {
	disp := &phy.Vector{X: p.vel.X, Y: p.vel.Y}
	disp.Mult(dt)
	p.transform.Translate(disp)
}

func (p *Particle) Draw() {
	tex := sdl.Rect{
		X: int32(p.transform.X),
		Y: int32(p.transform.Y),
		W: p.size,
		H: p.size,
	}
	EngineInstance.GetInstance().GetRenderer().DrawRect(&tex)
}

func (p *Particle) IsDead() bool {
	window := EngineInstance.GetInstance().GetWindow()
	w, h := window.GetSize()
	return sdl.GetTicks64() >= uint64(p.lifeTime.Seconds()) || p.transform.X <= 0 || p.transform.Y <= 0 || p.transform.X >= float64(w) || p.transform.Y >= float64(h)
}
