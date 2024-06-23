package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Engine struct {
	instance *Engine
	window   *sdl.Window
	renderer *sdl.Renderer

	IsRunning bool
}

func (e *Engine) Init() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("failed to initialize SDL: %v", err)
	}

	err = img.Init(img.INIT_PNG)
	if err != nil {
		return fmt.Errorf("failed to initialize SDL: %v", err)
	}

	window, err := sdl.CreateWindow("Game", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, WIDTH, HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("failed to create window: %v", err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		return fmt.Errorf("failed to create renderer: %v", err)
	}

	e.window = window
	e.renderer = renderer
	e.IsRunning = true
	return nil
}

func (e *Engine) Load() error {
	TextureManagerInstance.GetInstance()
	err := TextureManagerInstance.LoadTexture("ghost", "assets/ghost.png")
	if err != nil {
		return fmt.Errorf("failed to load texture: %v", err)
	}
	return nil
}

// Getter Methods
func (e *Engine) GetInstance() *Engine {
	if e.instance == nil {
		e.instance = &Engine{}
	}
	return e.instance
}

func (e *Engine) GetWindow() *sdl.Window {
	return e.window
}

func (e *Engine) GetRenderer() *sdl.Renderer {
	return e.renderer
}

func (e *Engine) Update() {
	//log.Print("Update")
}

func (e *Engine) Events() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			e.IsRunning = false
		}
	}
}

func (e *Engine) Render() {
	e.renderer.Clear()
	e.renderer.SetDrawColor(0, 0, 0, 255)
	TextureManagerInstance.Draw("ghost", 0, 0, 32, 32, sdl.FLIP_NONE)
	e.renderer.Present()
}

func (e *Engine) Destroy() {
	e.renderer.Destroy()
	e.window.Destroy()
	img.Quit()
	sdl.Quit()
	e.IsRunning = false
}
