package main

import (
	"fmt"
	"go-game/phy"

	"github.com/veandco/go-sdl2/sdl"
)

type GameState interface {
	Init()
	Draw()
	Update()
	Exit()
}

type PlayState struct {
	GameState
	levelMap    *GameMap[TileLayer]
	gameObjects []Object
	renderer    *sdl.Renderer
}

func (p *PlayState) Init() error {
	p.renderer = EngineInstance.GetInstance().GetRenderer()

	err := TextureManagerInstance.GetInstance().LoadAllTextures("assets/textures.xml")
	if err != nil {
		return fmt.Errorf("failed to load textures, %v", err)
	}

	err = MapParserInstance.GetInstance().Load()
	if err != nil {
		return fmt.Errorf("failed to load map parser, %v", err)
	}

	p.levelMap = MapParserInstance.GetInstance().GetGameMap("level1")
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

	p.gameObjects = append(p.gameObjects, player, enemy)

	CameraInstance.GetInstance().SetTarget(player.GetOrigin())
	return nil
}

func (p *PlayState) OpenMenu() {}

func (p *PlayState) PauseGame() {}

func (p *PlayState) Events() {
	// hanlde for menu settings
	// Ctrl + M to toggle menu state on/off
	// Esc to toggle play state on/off
}

func (p *PlayState) Draw() {
	p.renderer.SetDrawColor(0, 0, 0, 255)
	p.renderer.Clear()
	TextureManagerInstance.GetInstance().Draw("bg", 0, 0, WIDTH, HEIGHT, 1, 1, 0.5, sdl.FLIP_NONE)
	p.levelMap.Draw()
	for _, gObj := range p.gameObjects {
		gObj.Draw()
	}
	p.renderer.Present()
}

func (p *PlayState) Update() {
	p.Events()

	dt := TimeInstance.GetInstance().GetDeltaTime()
	p.levelMap.Update(dt)
	for _, gObj := range p.gameObjects {
		gObj.Update(dt)
	}
	CameraInstance.GetInstance().Update(dt)
}

func (p *PlayState) Exit() {
	for _, gObj := range p.gameObjects {
		gObj.Destroy()
	}
	TextureManagerInstance.GetInstance().Destroy()
}
