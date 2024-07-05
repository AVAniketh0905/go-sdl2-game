package main

import (
	"fmt"
	"go-game/phy"
	"log"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Engine struct {
	instance    *Engine
	window      *sdl.Window
	renderer    *sdl.Renderer
	levelMap    *GameMap[TileLayer]
	gameObjects []Object
	states      []GameState

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

	e.window = window
	e.renderer = renderer
	e.IsRunning = true
	return &e, nil
}

func (e *Engine) Load() error {
	err := TextureManagerInstance.GetInstance().LoadAllTextures("assets/textures.xml")
	if err != nil {
		return fmt.Errorf("failed to load textures, %v", err)
	}

	err = MapParserInstance.GetInstance().Load()
	if err != nil {
		return fmt.Errorf("failed to load map parser, %v", err)
	}

	e.levelMap = MapParserInstance.GetInstance().GetGameMap("level1")
	lvlLayers := EngineInstance.GetInstance().GetLevelMap().GetLayers()
	tileSize := lvlLayers[0].tileSize
	width, height := lvlLayers[0].GetWidth()*tileSize, lvlLayers[0].GetHeight()*tileSize

	CameraInstance.GetInstance().SetScreenLimit(int32(width), int32(height))
	CollisionHandlerInstance.GetInstance().SetCollisionMap(lvlLayers[0].tileMap, lvlLayers[0].tileSize)

	player, err := CreateObjectFactory("Player", &Properties{
		transform: &phy.Transform{X: 10, Y: 20},
		width:     IMG_SIZE,
		height:    IMG_SIZE,
		texId:     "player_idle",
		flip:      sdl.FLIP_NONE,
	})
	if err != nil {
		return err
	}

	enemy, err := CreateObjectFactory("Enemy", &Properties{
		transform: &phy.Transform{X: 120, Y: 00},
		width:     IMG_SIZE,
		height:    IMG_SIZE,
		texId:     "boss_load",
		flip:      sdl.FLIP_NONE,
	})
	if err != nil {
		return fmt.Errorf("failed to load enemy, %v", err)
	}

	e.gameObjects = append(e.gameObjects, player, enemy)
	CameraInstance.GetInstance().SetTarget(player.GetOrigin())

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

func (e *Engine) GetLevelMap() *GameMap[TileLayer] {
	return e.levelMap
}

func (e *Engine) PopState()                     {}
func (e *Engine) PushState(curr *GameState)     {}
func (e *Engine) ChangeState(target *GameState) {}

// Game Engine
func (e *Engine) Update() {
	dt := TimeInstance.GetInstance().GetDeltaTime()
	e.levelMap.Update(dt)
	CameraInstance.GetInstance().Update(dt)
	for _, gObj := range e.gameObjects {
		gObj.Update(dt)
	}
}

func (e *Engine) Events() {
	InputInstance.GetInstance().Listen()
}

func (e *Engine) Draw() {
	e.renderer.SetDrawColor(0, 0, 0, 255)
	e.renderer.Clear()
	TextureManagerInstance.GetInstance().Draw("bg", 0, 0, WIDTH, HEIGHT, 1, 1, 0.5, sdl.FLIP_NONE)
	e.levelMap.Draw()
	for _, gObj := range e.gameObjects {
		gObj.Draw()
	}
	e.renderer.Present()
}

func (e *Engine) Destroy() {
	e.renderer.Destroy()
	e.window.Destroy()
	img.Quit()
	sdl.Quit()
	e.IsRunning = false
}
