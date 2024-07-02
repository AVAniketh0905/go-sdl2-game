package main

import (
	"encoding/xml"
	"fmt"
	"os"

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

func (tm *TextureManager) parseTextures(path string) (*XMLTextures, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to real file from %v, %v", path, err)
	}

	var textures XMLTextures
	err = xml.Unmarshal(data, &textures)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to xml %v", err)
	}

	return &textures, nil
}

func (tm *TextureManager) LoadTexture(id string, path string) error {
	surface, err := img.Load(path)
	if err != nil {
		return fmt.Errorf("failed to load image: %v", err)
	}

	texture, err := EngineInstance.GetInstance().GetRenderer().CreateTextureFromSurface(surface)
	if err != nil {
		return fmt.Errorf("failed to create texture: %v", err)
	}

	tm.SetTextureMap(id, texture)
	return nil
}

func (tm *TextureManager) LoadAllTextures(path string) error {
	textures, err := tm.parseTextures(path)
	if err != nil {
		return fmt.Errorf("failed to get textures, %v", err)
	}

	for _, tex := range textures.Textures {
		err = tm.LoadTexture(tex.Id, tex.Src)
		if err != nil {
			return fmt.Errorf("failed to load the texture, id: %v, %v", tex.Id, err)
		}
	}

	return nil
}

func (tm *TextureManager) Draw(id string, x, y, width, height int, scrollRatio float64, flip sdl.RendererFlip) error {
	cam := CameraInstance.GetInstance().GetPosition()
	src := sdl.Rect{X: 0, Y: 0, W: int32(width), H: int32(height)}
	dst := sdl.Rect{
		X: int32(x) - int32(cam.X*scrollRatio),
		Y: int32(y) - int32(cam.Y*scrollRatio),
		W: int32(width),
		H: int32(height),
	}

	err := EngineInstance.GetInstance().GetRenderer().CopyEx(tm.textureMap[id], &src, &dst, 0, nil, flip)
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

	err := EngineInstance.GetInstance().GetRenderer().CopyEx(tm.textureMap[id], &src, &dst, 0, nil, flip)
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

	err := EngineInstance.GetInstance().GetRenderer().CopyEx(tm.textureMap[tileSetId], &src, &dst, 0, nil, flip)
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

type XMLTextures struct {
	Textures []XMLTexture `xml:"texture"`
}

type XMLTexture struct {
	Id  string `xml:"id,attr"`
	Src string `xml:"src,attr"`
}
