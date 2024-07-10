package main

import (
	"fmt"
	"go-game/phy"

	"github.com/veandco/go-sdl2/sdl"
)

type GameState interface {
	Draw()
	Update(uint64)
	Exit()
}

type PlayState struct {
	GameState
	renderer *sdl.Renderer
	bg       *sdl.Color
	menuObjs []Object
}

func PlayStateInit() (*PlayState, error) {
	p := &PlayState{}
	p.renderer = EngineInstance.GetInstance().GetRenderer()
	err := LevelManagerInsatance.GetInstance().Init()
	if err != nil {
		return nil, fmt.Errorf("failed to load lvls, %v", err)
	}
	p.bg = LevelManagerInsatance.GetInstance().GetBgColor()

	props := Properties{
		transform: &phy.Transform{X: 10, Y: 20},
		width:     128,
		height:    128,
		texId:     "",
		flip:      sdl.FLIP_NONE,
	}
	texIds := []string{"default_btn", "hover_btn", "active_btn"}
	button, err := NewButton(&props, texIds, p.OpenMenu)
	if err != nil {
		return nil, err
	}

	p.menuObjs = append(p.menuObjs, button)
	return p, nil
}

func (p *PlayState) OpenMenu() {
	EngineInstance.GetInstance().SetCurrStateName(MENU)
}

func (p *PlayState) PauseGame() {}

func (p *PlayState) Events() {
	if InputInstance.GetInstance().IsKeyDown(sdl.SCANCODE_M) {
		p.OpenMenu()
	}
}

func (p PlayState) Draw() {
	p.renderer.SetDrawColor(p.bg.R, p.bg.G, p.bg.B, p.bg.A)
	p.renderer.Clear()
	LevelManagerInsatance.GetInstance().Draw()
	for _, mobj := range p.menuObjs {
		mobj.Draw()
	}
	p.renderer.Present()
}

func (p PlayState) Update(dt uint64) {
	p.Events()
	LevelManagerInsatance.GetInstance().Update(dt)
	for _, mobj := range p.menuObjs {
		mobj.Update(dt)
	}
}

func (p PlayState) Exit() {
	for _, mobj := range p.menuObjs {
		mobj.Destroy()
	}
	LevelManagerInsatance.GetInstance().Destroy()
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
	EngineInstance.GetInstance().GetCurrState().Exit()
	EngineInstance.GetInstance().SetCurrStateName(PLAY)
}

func (m *MenuState) SaveGame() {}

func (m *MenuState) Events() {
	if InputInstance.GetInstance().IsKeyDown(sdl.SCANCODE_RSHIFT) {
		m.StartGame()
	}
}

func (m MenuState) Draw() {
	m.renderer.SetDrawColor(0, 0, 0, 250)
	m.renderer.Clear()
	m.renderer.Present()
}

func (m MenuState) Update(dt uint64) {
	m.Events()
}

func (m MenuState) Exit() {}

func (m *MenuState) Quit() {}
