package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type TextureManager struct {
	instance *TextureManager

	textureMap map[string]*sdl.Texture
}

func (tm *TextureManager) GetInstance() *TextureManager {
	if tm.instance == nil {
		tm.instance = &TextureManager{
			textureMap: make(map[string]*sdl.Texture),
		}
	}
	return tm.instance
}

func (tm *TextureManager) SetTextureMap(id string, texture *sdl.Texture) error {
	tm.textureMap[id] = texture
	return nil
}

func (tm *TextureManager) LoadTexture(id string, path string) error {
	surface, err := img.Load(path)
	if err != nil {
		return fmt.Errorf("failed to load image: %v", err)
	}

	texture, err := EngineInstance.GetRenderer().CreateTextureFromSurface(surface)
	if err != nil {
		return fmt.Errorf("failed to create texture: %v", err)
	}

	tm.SetTextureMap(id, texture)
	return nil
}

func (tm *TextureManager) Draw(id string, x int, y int, width int, height int, flip sdl.RendererFlip) error {
	cam := CameraInstance.GetInstance().GetPosition()

	src := sdl.Rect{X: 0, Y: 0, W: int32(width), H: int32(height)}
	dst := sdl.Rect{
		X: int32(x) - int32(cam.X*0.5),
		Y: int32(y) - int32(cam.Y*0.5),
		W: int32(width),
		H: int32(height),
	}

	err := EngineInstance.GetRenderer().CopyEx(tm.textureMap[id], &src, &dst, 0, nil, flip)
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

	err := EngineInstance.GetRenderer().CopyEx(tm.textureMap[id], &src, &dst, 0, nil, flip)
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

	err := EngineInstance.GetRenderer().CopyEx(tm.textureMap[tileSetId], &src, &dst, 0, nil, flip)
	if err != nil {
		return fmt.Errorf("failed to copy texture: %v", err)
	}
	return nil
}

func (tm *TextureManager) DropTexture(id string) {

}

func (tm *TextureManager) Destroy() {
	for _, texture := range tm.textureMap {
		texture.Destroy()
	}
	tm.textureMap = nil
	tm.instance = nil
}
