package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Animation struct {
	animSpeed  int
	flip       sdl.RendererFlip
	frame      int
	frameCount int
	row        int
	texId      string
}

func NewAnimation(animSpeed, frameCount int, flip sdl.RendererFlip, texId string) *Animation {
	return &Animation{
		animSpeed:  animSpeed,
		flip:       sdl.FLIP_NONE,
		frame:      0,
		frameCount: frameCount,
		row:        0,
		texId:      texId,
	}
}

func (a Animation) Draw(x, y, width, height int) {
	err := TextureManagerInstance.DrawFrame(a.texId, x, y, width, height, a.row, a.frame, a.flip)
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
