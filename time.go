package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Time struct {
	instance *Time

	deltaTime float64
	lastTime  float64
}

func (t *Time) GetInstance() *Time {
	if t.instance == nil {
		t.instance = &Time{}
	}
	return t.instance
}

func (t *Time) GetDeltaTime() float64 {
	return t.deltaTime
}

func (t *Time) Tick() {
	currentTime := float64(sdl.GetTicks64())

	t.deltaTime += (currentTime - t.lastTime)

	if t.deltaTime >= DELTA_TIME {
		t.deltaTime = DELTA_TIME
	}

	t.lastTime = float64(sdl.GetTicks64())
}
