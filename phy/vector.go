package phy

import (
	"fmt"
	"math"
)

type Vector struct {
	X float64
	Y float64
}

func (v *Vector) Add(v2 *Vector) {
	v.X += v2.X
	v.Y += v2.Y
}

func (v *Vector) Sub(v2 *Vector) {
	v.X -= v2.X
	v.Y -= v2.Y
}

func (v *Vector) Mult(scalar float64) {
	v.X *= scalar
	v.Y *= scalar
}

func (v *Vector) Div(scalar float64) {
	v.X /= scalar
	v.Y /= scalar
}

func (v *Vector) Mag() float64 {
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func (v *Vector) Normalize() {
	mag := v.Mag()
	if mag != 0 {
		v.Div(mag)
	}
}

func (v *Vector) Limit(max float64) {
	if v.Mag() > max {
		v.Normalize()
		v.Mult(max)
	}
}

func (v *Vector) SetMag(mag float64) {
	v.Normalize()
	v.Mult(mag)
}

func (v *Vector) Copy() *Vector {
	return &Vector{X: v.X, Y: v.Y}
}

func (v *Vector) String() string {
	return fmt.Sprintf("X: %d, Y: %d", int(v.X), int(v.Y))
}
