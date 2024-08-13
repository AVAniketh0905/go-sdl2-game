package main

import (
	"fmt"
	"go-game/phy"

	"github.com/veandco/go-sdl2/sdl"
)

type LevelManager struct {
	instance *LevelManager

	bg          *sdl.Color
	state       GStateType
	levelMap    *GameMap[TileLayer]
	healthBar   *Bars
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
	// lm.gameObjects, err = ObjectParserInstance.GetInstance().Load("assets/objects.xml")
	// if err != nil {
	// 	return fmt.Errorf("failed to load game objects, %x", err)
	// }
	lvlLayers := lm.levelMap.GetLayers()

	lm.state = PLAY

	blockTiles := lvlLayers[0] // block layer
	tileSize := blockTiles.tileSize
	width, height := blockTiles.GetWidth()*tileSize, blockTiles.GetHeight()*tileSize

	CollisionHandlerInstance.GetInstance().SetCollisionMap(blockTiles.tileMap, blockTiles.tileSize)

	damageTiles := lvlLayers[1] // damage layer
	DamageHandlerInstance.GetInstance().SetCollisionMap(damageTiles.tileMap, damageTiles.tileSize)

	coinTiles := lvlLayers[2] // coin layer
	CoinHandlerInstance.GetInstance().SetCollisionMap(coinTiles.tileMap, coinTiles.tileSize)

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

	// enemyObjs, err := ObjectParserInstance.GetInstance().Load("assets/objects.xml")
	// if err != nil {
	// 	return err
	// }

	healthBar, err := NewBars(&Properties{
		transform: &phy.Transform{X: 80, Y: 20},
		width:     64,
		height:    16,
		texId:     "",
		flip:      sdl.FLIP_NONE,
	}, []string{"health_bar", "health_bar_fill"})
	if err != nil {
		return err
	}
	lm.healthBar = healthBar

	lm.gameObjects = append(lm.gameObjects, player)
	lm.gameObjects = append(lm.gameObjects, healthBar)
	// lm.gameObjects = append(lm.gameObjects, enemyObjs...)

	CameraInstance.GetInstance().SetTarget(player.GetOrigin())
	CameraInstance.GetInstance().SetLevelLimit(int32(width), int32(height))
	CameraInstance.GetInstance().SetViewBox(0, 0, WIDTH, HEIGHT)
	return nil
}

func (lm *LevelManager) GetBgColor() *sdl.Color {
	return lm.bg
}

func (lm *LevelManager) GetState() GStateType {
	return lm.state
}

func (lm *LevelManager) SetState(state GStateType) {
	switch state {
	case PLAY:
		lm.state = PLAY
	case SUCCESS:
		lm.state = SUCCESS
	case FAIL:
		lm.state = FAIL
	}
}

func (lm *LevelManager) UpdateHealthBar(newHealth int) {
	lm.healthBar.SetBarWidth((lm.healthBar.width * (newHealth) / (MAX_HEALTH)))
}

func (lm *LevelManager) Draw() {
	TextureManagerInstance.GetInstance().Draw("bg", 0, 0, WIDTH, HEIGHT, 1, 1, 0.9, sdl.FLIP_NONE, false)
	lm.levelMap.Draw()
	for _, obj := range lm.gameObjects {
		obj.Draw()
	}
	CameraInstance.GetInstance().Draw()
}

func (lm *LevelManager) ApplyState() {
	if CameraInstance.GetInstance().GetTarget().X > 1880 {
		SoundManagerInstance.GetInstance().PlayEffect("success")
		lm.SetState(SUCCESS)
	} else if CameraInstance.GetInstance().GetTarget().Y > 500 {
		SoundManagerInstance.GetInstance().PlayEffect("fail")
		lm.SetState(FAIL)
	} else {
		lm.SetState(PLAY)
	}
}

func (lm *LevelManager) Update(dt uint64) {
	lm.levelMap.Update(dt)
	for _, obj := range lm.gameObjects {
		obj.Update(dt)
	}
	CameraInstance.GetInstance().Update(dt)
	if lm.GetState() != FAIL {
		lm.ApplyState()
	}
}

func (lm *LevelManager) Destroy() {
	for _, obj := range lm.gameObjects {
		obj.Destroy()
	}
	lm.gameObjects = nil
	lm.levelMap = nil
	lm.instance = nil
}
