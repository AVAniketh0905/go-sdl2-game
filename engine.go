package main

import (
	"fmt"
	"go-game/phy"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

var PlayerGhost *Ghost

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
	err := TextureManagerInstance.LoadTexture("ghost", "assets/ghost_2.png")
	if err != nil {
		return fmt.Errorf("failed to load texture: %v", err)
	}

	PlayerGhost = NewGhost(&Properties{
		transform: &phy.Transform{
			Position: &phy.Vector{
				X: 10,
				Y: 20,
			},
		},
		width:  IMG_SIZE,
		height: IMG_SIZE,
		texId:  "ghost",
		flip:   sdl.FLIP_NONE,
	})

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
	PlayerGhost.Update(0.1)
}

func (e *Engine) Events() {
	InputInstance.GetInstance().Listen()
}

func (e *Engine) Render() {
	e.renderer.Clear()
	e.renderer.SetDrawColor(0, 0, 0, 255)
	PlayerGhost.Draw()
	e.renderer.Present()
}

func (e *Engine) Destroy() {
	e.renderer.Destroy()
	e.window.Destroy()
	img.Quit()
	sdl.Quit()
	e.IsRunning = false
}
