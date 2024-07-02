package phy

type Transform struct {
	X float64
	Y float64
}

func (t *Transform) Set(v Vector) {
	t.X = v.X
	t.Y = v.Y
}

func (t *Transform) TranslateX(x float64) {
	t.X += x
}

func (t *Transform) TranslateY(y float64) {
	t.Y += y
}

func (t *Transform) Translate(v *Vector) {
	t.X += v.X
	t.Y += v.Y
}

func (t *Transform) String() string {
	vec := Vector{X: t.X, Y: t.Y}
	return vec.String()
}
