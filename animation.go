package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type PublicAnim struct {
	animSpeed  int
	flip       sdl.RendererFlip
	frameCount int
	texId      string
}

type Animation struct {
	PublicAnim
	frame int
	row   int
}

func NewAnimation(props PublicAnim) *Animation {
	return &Animation{
		PublicAnim: props,
		frame:      0,
		row:        0,
	}
}

func (a *Animation) SetProps(props PublicAnim) {
	a.animSpeed = props.animSpeed
	a.frameCount = props.frameCount
	a.flip = props.flip
	a.texId = props.texId
}

func (a Animation) Draw(x, y, width, height int) {
	err := TextureManagerInstance.GetInstance().DrawFrame(a.texId, x, y, width, height, a.row, a.frame, a.flip)
	if err != nil {
		panic(err)
	}
}

func (a *Animation) Update(dt float64) {
	time := sdl.GetTicks64()
	a.frame = int(int(time)/(a.animSpeed)) % a.frameCount
}

func (a Animation) Destroy() {
	TextureManagerInstance.Destroy()
}
