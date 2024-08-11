package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Text struct {
	font    *ttf.Font
	texture *sdl.Texture

	dst sdl.Rect

	Content  string
	Color    sdl.Color
	FontSize int
	FontPath string
}

func NewText(fontSize int, fontPath string, color sdl.Color, text string, x, y, width, height int32) (*Text, error) {
	tr := &Text{}
	tr.FontSize = fontSize
	tr.FontPath = fontPath
	tr.Content = text
	tr.Color = color

	font, err := ttf.OpenFont(tr.FontPath, tr.FontSize)
	if err != nil {
		return nil, err
	}
	tr.font = font

	surface, err := tr.font.RenderUTF8Blended(tr.Content, tr.Color)
	if err != nil {
		return nil, err
	}

	texture, err := EngineInstance.GetInstance().GetRenderer().CreateTextureFromSurface(surface)
	if err != nil {
		return nil, err
	}

	tr.texture = texture

	tr.dst = sdl.Rect{X: x - surface.W/2, Y: y - surface.H, W: width, H: height}
	return tr, nil
}

func (tr *Text) Draw() {
	EngineInstance.GetInstance().GetRenderer().Copy(tr.texture, nil, &tr.dst)
}

func (tr *Text) Destroy() {
	tr.font.Close()
	tr.texture.Destroy()
}
