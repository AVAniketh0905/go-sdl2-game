package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type GStateType string

const (
	PLAY GStateType = "play"
	MENU GStateType = "menu"
)

type Engine struct {
	instance *Engine
	window   *sdl.Window
	renderer *sdl.Renderer

	screenSize *sdl.Rect
	states     map[GStateType]GameState

	currStateName GStateType

	IsRunning bool
}

func EngineInit() (*Engine, error) {
	var e Engine

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize SDL: %v", err)
	}

	err = img.Init(img.INIT_PNG)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize SDL: %v", err)
	}

	window, err := sdl.CreateWindow("Game", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WIDTH, HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, fmt.Errorf("failed to create window: %v", err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		return nil, fmt.Errorf("failed to create renderer: %v", err)
	}

	var DM sdl.DisplayMode

	if err := mix.Init(int(mix.INIT_MP3)); err != nil {
		return nil, fmt.Errorf("failed to load sdl mixer, %v", err)
	}

	e.window = window
	e.renderer = renderer
	DM, err = sdl.GetCurrentDisplayMode(0)
	if err != nil {
		return nil, err
	}
	e.screenSize = &sdl.Rect{X: 0, Y: 0, W: DM.W, H: DM.H}
	e.states = map[GStateType]GameState{}
	e.IsRunning = true
	return &e, nil
}

func (e *Engine) Load() error {
	playState, err := PlayStateInit()
	if err != nil {
		return err
	}

	menuState, err := MenuStateInit()
	if err != nil {
		return err
	}

	e.states["play"] = playState
	e.states["menu"] = menuState

	e.currStateName = MENU
	SoundManagerInstance.GetInstance().PlayMusic("eerie")
	return nil
}

// Getter Methods
func (e *Engine) GetInstance() *Engine {
	if e.instance == nil {
		engine, err := EngineInit()
		if err != nil {
			log.Fatal("failed to load the engine")
		}
		e.instance = engine
	}
	return e.instance
}

func (e *Engine) GetWindow() *sdl.Window {
	return e.window
}

func (e *Engine) GetRenderer() *sdl.Renderer {
	return e.renderer
}

func (e *Engine) GetScreenSize() *sdl.Rect {
	return e.screenSize
}

// Game State
func (e *Engine) GetCurrState() GameState {
	return e.states[e.currStateName]
}

func (e *Engine) GetCurrStateName() GStateType {
	return e.currStateName
}

func (e *Engine) SetCurrStateName(stateName GStateType) {
	e.currStateName = stateName
}

// Game Engine
func (e *Engine) Update(dt uint64) {
	state := e.GetCurrState()
	state.Update(dt)
}

func (e *Engine) Events() {
	InputInstance.GetInstance().Listen()
}

func (e *Engine) Draw() {
	state := e.GetCurrState()
	state.Draw()
}

func (e *Engine) Destroy() {
	e.renderer.Destroy()
	e.window.Destroy()
	img.Quit()
	mix.Quit()
	sdl.Quit()
	e.IsRunning = false
}
