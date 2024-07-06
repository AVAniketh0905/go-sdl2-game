package main

import (
	"fmt"
	"go-game/phy"

	"github.com/veandco/go-sdl2/sdl"
)

type GameState interface {
	Draw()
	Update(float64)
	Exit()
}

type PlayState struct {
	GameState
	levelMap    *GameMap[TileLayer]
	gameObjects []Object
	renderer    *sdl.Renderer
}

func PlayStateInit() (*PlayState, error) {
	var p PlayState
	renderer := EngineInstance.GetInstance().GetRenderer()

	err := TextureManagerInstance.GetInstance().LoadAllTextures("assets/textures.xml")
	if err != nil {
		return nil, fmt.Errorf("failed to load textures, %v", err)
	}

	err = MapParserInstance.GetInstance().Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load map parser, %v", err)
	}

	levelMap := MapParserInstance.GetInstance().GetGameMap("level1")

	lvlLayers := levelMap.GetLayers()
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
		return nil, err
	}

	enemyObjs, err := ObjectParserInstance.GetInstance().Load("assets/objects.xml")
	if err != nil {
		return nil, err
	}

	p.renderer = renderer
	p.levelMap = levelMap
	p.gameObjects = append(p.gameObjects, player)
	p.gameObjects = append(p.gameObjects, enemyObjs...)

	CameraInstance.GetInstance().SetTarget(player.GetOrigin())

	return &p, nil
}

func (p *PlayState) OpenMenu() {
	if InputInstance.GetInstance().IsKeyDown(sdl.SCANCODE_M) {
		EngineInstance.GetInstance().SetCurrStateName(MENU)
	}
}

func (p *PlayState) PauseGame() {}

func (p *PlayState) Events() {
	p.OpenMenu()
}

func (p PlayState) Draw() {
	p.renderer.SetDrawColor(0, 0, 0, 255)
	p.renderer.Clear()
	TextureManagerInstance.GetInstance().Draw("bg", 0, 0, WIDTH, HEIGHT, 1, 1, 0.5, sdl.FLIP_NONE)
	p.levelMap.Draw()
	for _, gObj := range p.gameObjects {
		gObj.Draw()
	}
	p.renderer.Present()
}

func (p PlayState) Update(dt float64) {
	p.Events()

	p.levelMap.Update(dt)
	for _, gObj := range p.gameObjects {
		gObj.Update(dt)
	}
	CameraInstance.GetInstance().Update(dt)
}

func (p PlayState) Exit() {
	for _, gObj := range p.gameObjects {
		gObj.Destroy()
	}
	TextureManagerInstance.GetInstance().Destroy()
}

type MenuState struct {
	GameState
	renderer *sdl.Renderer
}

func MenuStateInit() (*MenuState, error) {
	var m MenuState
	m.renderer = EngineInstance.GetInstance().GetRenderer()

	return &m, nil
}

func (m *MenuState) Settings() {}

func (m *MenuState) StartGame() {
	if InputInstance.GetInstance().IsKeyDown(sdl.SCANCODE_RSHIFT) {
		EngineInstance.GetInstance().GetCurrState().Exit()
		EngineInstance.GetInstance().SetCurrStateName(PLAY)
	}
}

func (m *MenuState) SaveGame() {}

func (m *MenuState) Events() {
	m.StartGame()
}

func (m MenuState) Draw() {
	m.renderer.SetDrawColor(0, 0, 0, 250)
	m.renderer.Clear()
	m.renderer.Present()
}

func (m MenuState) Update(dt float64) {
	m.Events()
}

func (m MenuState) Exit() {}

func (m *MenuState) Quit() {}
