package phy

type Transform struct {
	Position *Vector
}

func (t *Transform) TranslateX(v *Vector) {
	t.Position.X += v.X
}

func (t *Transform) TranslateY(v *Vector) {
	t.Position.Y += v.Y
}

func (t *Transform) Translate(v *Vector) {
	t.Position.Add(v)
}

func (t *Transform) String() string {
	return t.Position.String()
}
