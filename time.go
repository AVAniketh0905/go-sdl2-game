package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Time struct {
	instance *Time

	deltaTime uint64
	lastTime  uint64
}

func (t *Time) GetInstance() *Time {
	if t.instance == nil {
		t.instance = &Time{}
	}
	return t.instance
}

func (t *Time) GetDeltaTime() uint64 {
	return t.deltaTime
}

func (t *Time) Start() {
	t.lastTime = sdl.GetTicks64()
}

func (t *Time) Tick() {
	currentTime := sdl.GetTicks64()
	t.deltaTime = (currentTime - t.lastTime)

	if t.deltaTime > TIME_DELAY {
		t.deltaTime = TIME_DELAY
	}
}
