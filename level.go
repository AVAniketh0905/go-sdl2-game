package main

import (
	"fmt"
	"go-game/phy"

	"github.com/veandco/go-sdl2/sdl"
)

type LevelManager struct {
	instance *LevelManager

	bg          *sdl.Color
	levelMap    *GameMap[TileLayer]
	gameObjects []Object
}

func (lm *LevelManager) GetInstance() *LevelManager {
	if lm.instance == nil {
		lm.instance = &LevelManager{}
	}

	return lm.instance
}

func (lm *LevelManager) SetBackground(bg *sdl.Color) {
	lm.bg = bg
}

func (lm *LevelManager) Init() error {
	err := SoundParserInstance.GetInstance().Load("assets/sounds.xml")
	if err != nil {
		return fmt.Errorf("failed to load sounds, %v", err)
	}
	err = TextureParserInstance.GetInstance().LoadAllTextures("assets/textures.xml")
	if err != nil {
		return fmt.Errorf("failed to load textures, %v", err)
	}
	err = MapParserInstance.GetInstance().Load("level1", "assets/maps/go-sdl2-level_1.tmx")
	if err != nil {
		return fmt.Errorf("failed to load map, %v", err)
	}
	lm.levelMap = MapParserInstance.GetInstance().GetGameMap("level1")
	lm.gameObjects, err = ObjectParserInstance.GetInstance().Load("assets/objects.xml")
	if err != nil {
		return fmt.Errorf("failed to load game objects, %x", err)
	}
	lvlLayers := lm.levelMap.GetLayers()
	tileSize := lvlLayers[0].tileSize
	width, height := lvlLayers[0].GetWidth()*tileSize, lvlLayers[0].GetHeight()*tileSize

	CollisionHandlerInstance.GetInstance().SetCollisionMap(lvlLayers[0].tileMap, lvlLayers[0].tileSize)

	lm.bg = &sdl.Color{R: 0, G: 0, B: 0, A: 255}

	player, err := CreateObjectFactory("Player", &Properties{
		transform: &phy.Transform{X: 10, Y: 0},
		width:     IMG_SIZE,
		height:    IMG_SIZE,
		texId:     "player_idle",
		flip:      sdl.FLIP_NONE,
	})
	if err != nil {
		return err
	}

	enemyObjs, err := ObjectParserInstance.GetInstance().Load("assets/objects.xml")
	if err != nil {
		return err
	}

	lm.gameObjects = append(lm.gameObjects, player)
	lm.gameObjects = append(lm.gameObjects, enemyObjs...)

	CameraInstance.GetInstance().SetTarget(player.GetOrigin())
	CameraInstance.GetInstance().SetLevelLimit(WIDTH, HEIGHT)
	CameraInstance.GetInstance().SetViewBox(0, 0, int32(width), int32(height))
	return nil
}

func (lm *LevelManager) GetBgColor() *sdl.Color {
	return lm.bg
}

func (lm *LevelManager) Draw() {
	TextureManagerInstance.GetInstance().Draw("bg", 0, 0, WIDTH, HEIGHT, 1, 1, 0.5, sdl.FLIP_NONE)
	CameraInstance.GetInstance().Draw()
	lm.levelMap.Draw()
	for _, obj := range lm.gameObjects {
		obj.Draw()
	}
}

func (lm *LevelManager) Update(dt uint64) {
	CameraInstance.GetInstance().Update(dt)
	lm.levelMap.Update(dt)
	for _, obj := range lm.gameObjects {
		obj.Update(dt)
	}
}

func (lm *LevelManager) Destroy() {
	for _, obj := range lm.gameObjects {
		obj.Destroy()
	}
	lm.levelMap = nil
	lm.instance = nil
}
