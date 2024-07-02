package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Input struct {
	instance  *Input
	keyStates []uint8
}

type Axis int

const (
	HORIZONTAL Axis = iota
	VERTICAL
)

func (a Axis) EnumIndex() int {
	return int(a)
}

func (i *Input) GetInstance() *Input {
	if i.instance == nil {
		i.instance = &Input{
			keyStates: sdl.GetKeyboardState(),
		}
	}
	return i.instance
}

func (i *Input) Listen() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			EngineInstance.GetInstance().IsRunning = false
		case *sdl.KeyboardEvent:
			if t.Type == sdl.KEYDOWN {
				i.KeyDown()
				break
			} else {
				i.KeyUp()
			}
		}
	}
}

func (i *Input) GetAxisKey(axis Axis) int {
	switch axis {
	case HORIZONTAL:
		if i.IsKeyDown(sdl.SCANCODE_A) || i.IsKeyDown(sdl.SCANCODE_LEFT) {
			return -1
		}
		if i.IsKeyDown(sdl.SCANCODE_D) || i.IsKeyDown(sdl.SCANCODE_RIGHT) {
			return 1
		}
	case VERTICAL:
		if i.IsKeyDown(sdl.SCANCODE_S) || i.IsKeyDown(sdl.SCANCODE_DOWN) {
			return -1
		}
		if i.IsKeyDown(sdl.SCANCODE_W) || i.IsKeyDown(sdl.SCANCODE_UP) {
			return 1
		}
	}

	return 0
}

func (i *Input) IsKeyDown(keyCode sdl.Scancode) bool {
	return i.keyStates[keyCode] == 1
}

func (i *Input) KeyDown() {
	i.keyStates = sdl.GetKeyboardState()
}

func (i *Input) KeyUp() {
	i.keyStates = sdl.GetKeyboardState()
}
