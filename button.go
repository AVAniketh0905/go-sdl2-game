package main

import (
	"fmt"
	"go-game/phy"

	"github.com/veandco/go-sdl2/sdl"
)

type Button struct {
	Object
	*Properties

	texIds []string
	mode   int

	isReleased bool
	callback   func()
}

func NewButton(props *Properties, texIds []string, callback func()) (*Button, error) {
	if len(texIds) < 3 {
		return nil, fmt.Errorf("please provide atleast 3 textures for btn, provided only %v", len(texIds))
	}

	b := &Button{
		mode:       DEFAULT_BTN,
		Properties: props,
		texIds:     texIds,
		callback:   callback,
	}
	b.SetPosition(props.transform.X-float64(props.width)/2, props.transform.Y-float64(props.height)/2)

	return b, nil
}

func (b *Button) GetTransform() *phy.Transform {
	return b.transform
}

func (b *Button) SetPosition(x, y float64) {
	b.transform.X = x
	b.transform.Y = y
}

func (b *Button) IsPointInMouse(mousePos *phy.Point) bool {
	if (mousePos.X >= b.transform.X && mousePos.X <= b.transform.X+float64(b.width/2)) && (mousePos.Y >= b.transform.Y && mousePos.Y <= b.transform.Y+float64(b.height/2)) {
		return true
	}

	return false
}

func (b *Button) GetOrigin() *phy.Point {
	return (*phy.Point)(b.transform)
}

func (b *Button) Draw() {
	TextureManagerInstance.GetInstance().Draw(b.texId, int(b.transform.X), int(b.transform.Y), b.width, b.height, 0.5, 0.5, 0, b.flip)
}

func (b *Button) Update(dt uint64) {
	mousePos, state := InputInstance.GetInstance().GetMousePosition()

	if b.IsPointInMouse(mousePos) {
		if state == sdl.BUTTON_LEFT && b.isReleased {
			b.callback()
			b.isReleased = false
			b.mode = ACTIVE_BTN
		} else if state != sdl.BUTTON_LEFT {
			b.isReleased = true
			b.mode = HOVER_BTN
		}
	} else {
		b.mode = DEFAULT_BTN
	}

	b.texId = b.texIds[b.mode]
}

func (b *Button) Destroy() {
	b.callback = func() {}
	b.texIds = []string{}
}
