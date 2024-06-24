package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Input struct {
	instance  *Input
	keyStates []uint8
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
			EngineInstance.IsRunning = false
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

func (i *Input) IsKeyDown(keyCode sdl.Scancode) bool {
	return i.keyStates[keyCode] == 1
}

func (i *Input) KeyDown() {
	i.keyStates = sdl.GetKeyboardState()
}

func (i *Input) KeyUp() {
	i.keyStates = sdl.GetKeyboardState()
}
