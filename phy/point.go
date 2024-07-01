package phy

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func (p *Point) Add(p2 *Point) {
	p.X += p2.X
	p.Y += p2.Y
}

func (p *Point) Sub(p2 *Point) {
	p.X -= p2.X
	p.Y -= p2.Y
}

func (p *Point) Mult(scalar float64) {
	p.X *= scalar
	p.Y *= scalar
}

func (p *Point) Div(scalar float64) {
	p.X /= scalar
	p.Y /= scalar
}

func (p *Point) Mag() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func (p *Point) Copy() *Point {
	return &Point{X: p.X, Y: p.Y}
}

func (p *Point) String() string {
	return fmt.Sprintf("X: %d, Y: %d", int(p.X), int(p.Y))
}
