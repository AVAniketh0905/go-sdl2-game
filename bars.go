package main

import (
	"go-game/phy"
)

const BAR_OFFSET = 0

type Bars struct {
	Object
	*Properties

	barWidth int
	texIds   []string
}

func NewBars(props *Properties, texIds []string) (*Bars, error) {
	return &Bars{
		Properties: props,
		texIds:     texIds,
		barWidth:   BAR_OFFSET + props.width,
	}, nil
}

func (b *Bars) GetTransform() *phy.Transform {
	return b.transform
}

func (b *Bars) SetBarWidth(w int) {
	b.barWidth = Limit(0, w+BAR_OFFSET, b.width)
}

func (b *Bars) Draw() {
	// Draw base
	base := b.texIds[0]
	TextureManagerInstance.GetInstance().Draw(base, int(b.transform.X), int(b.transform.Y), b.width, b.height, 4, 4, 0, b.flip)

	// // Draw bar
	bar := b.texIds[1]
	TextureManagerInstance.GetInstance().Draw(bar, int(b.transform.X), int(b.transform.Y), b.barWidth, b.height, 4, 4, 0, b.flip)
}

func (b *Bars) Update(dt uint64) {
}

func (b *Bars) Destroy() {
	b.Properties = nil
	b.texIds = nil
}
