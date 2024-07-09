package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type TextureParser struct {
	instance *TextureParser

	textureMap map[string]*sdl.Texture
}

func (tp *TextureParser) GetInstance() *TextureParser {
	if tp.instance == nil {
		tp.instance = &TextureParser{
			textureMap: make(map[string]*sdl.Texture),
		}
	}

	return tp.instance
}
func (tp *TextureParser) GetTextureMap() map[string]*sdl.Texture {
	return tp.textureMap
}

// func (tp *TextureParser) (){}

func (tp *TextureParser) parseTextures(path string) (*XMLTextures, error) {
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

func (tp *TextureParser) LoadTexture(id string, path string) error {
	surface, err := img.Load(path)
	if err != nil {
		return fmt.Errorf("failed to load image: %v", err)
	}

	texture, err := EngineInstance.GetInstance().GetRenderer().CreateTextureFromSurface(surface)
	if err != nil {
		return fmt.Errorf("failed to create texture: %v", err)
	}

	tp.textureMap[id] = texture
	return nil
}

func (tp *TextureParser) LoadAllTextures(path string) error {
	textures, err := tp.parseTextures(path)
	if err != nil {
		return fmt.Errorf("failed to get textures, %v", err)
	}

	for _, tex := range textures.Textures {
		err = tp.LoadTexture(tex.Id, tex.Src)
		if err != nil {
			return fmt.Errorf("failed to load the texture, id: %v, %v", tex.Id, err)
		}
	}

	return nil
}

func (tp *TextureParser) Destroy() {
	for _, texture := range tp.textureMap {
		texture.Destroy()
	}
	tp.instance = nil
}

type XMLTextures struct {
	Textures []XMLTexture `xml:"texture"`
}

type XMLTexture struct {
	Id  string `xml:"id,attr"`
	Src string `xml:"src,attr"`
}
