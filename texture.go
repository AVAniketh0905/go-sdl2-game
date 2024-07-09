package main

import (
	"fmt"
	"go-game/phy"

	"github.com/veandco/go-sdl2/sdl"
)

type TextureManager struct {
	instance *TextureManager
}

func (tm *TextureManager) GetInstance() *TextureManager {
	if tm.instance == nil {
		tm.instance = &TextureManager{}
	}
	return tm.instance
}

func (tm *TextureManager) Draw(id string, x, y, width, height int, scaleX, scaleY, scrollRatio float64, flip sdl.RendererFlip) error {
	dst_ := CameraInstance.GetInstance().SyncObject(&phy.Point{X: float64(x), Y: float64(y)}, int(scrollRatio))
	src := sdl.Rect{X: 0, Y: 0, W: int32(width), H: int32(height)}
	dst := sdl.Rect{
		X: int32(dst_.X),
		Y: int32(dst_.Y),
		W: int32(float64(width) * scaleX),
		H: int32(float64(height) * scaleY),
	}

	textureMap := TextureParserInstance.GetInstance().GetTextureMap()
	err := EngineInstance.GetInstance().GetRenderer().CopyEx(textureMap[id], &src, &dst, 0, nil, flip)
	if err != nil {
		return fmt.Errorf("failed to copy texture: %v", err)
	}
	return nil
}

func (tm *TextureManager) DrawFrame(id string, x int, y int, width int, height int, currentRow int, currentFrame int, flip sdl.RendererFlip) error {
	cam := CameraInstance.GetInstance().GetPosition()

	src := sdl.Rect{
		X: int32(width) * int32(currentFrame),
		Y: int32(height) * int32(currentRow),
		W: int32(width),
		H: int32(height),
	}
	dst := sdl.Rect{
		X: int32(x) - int32(cam.X),
		Y: int32(y) - int32(cam.Y),
		W: int32(3 * width / 2),
		H: int32(3 * height / 2),
	}

	textureMap := TextureParserInstance.GetInstance().GetTextureMap()
	err := EngineInstance.GetInstance().GetRenderer().CopyEx(textureMap[id], &src, &dst, 0, nil, flip)
	if err != nil {
		return fmt.Errorf("failed to copy texture: %v", err)
	}
	return nil
}

func (tm *TextureManager) DrawTile(tileSetId string, tileSize int, x int, y int, row int, frame int, flip sdl.RendererFlip) error {
	cam := CameraInstance.GetInstance().GetPosition()

	src := sdl.Rect{
		X: int32(tileSize * frame),
		Y: int32(tileSize * row),
		W: int32(tileSize),
		H: int32(tileSize),
	}
	dst := sdl.Rect{
		X: int32(x) - int32(cam.X),
		Y: int32(y) - int32(cam.Y),
		W: int32(tileSize),
		H: int32(tileSize),
	}

	textureMap := TextureParserInstance.GetInstance().GetTextureMap()
	err := EngineInstance.GetInstance().GetRenderer().CopyEx(textureMap[tileSetId], &src, &dst, 0, nil, flip)
	if err != nil {
		return fmt.Errorf("failed to copy texture: %v", err)
	}
	return nil
}

func (tm *TextureManager) DropTexture(id string) {
}

func (tm *TextureManager) Destroy() {
	tm.instance = nil
}
