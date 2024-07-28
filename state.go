package main

import (
	"fmt"
	"go-game/phy"
	"log"

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

func (p *PlayState) EndPlay() {
	switch LevelManagerInsatance.GetInstance().EndLevel() {
	case SUCCESS:
		print("Success")
		// TODO: Add Success Screen
		EngineInstance.GetInstance().SetCurrStateName(SUCCESS)
	case FAIL:
		print("Failure")
		// TODO: Add Failure Screen
		EngineInstance.GetInstance().SetCurrStateName(FAIL)
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
	p.EndPlay()
}

func (p PlayState) Exit() {
	for _, mobj := range p.menuObjs {
		mobj.Destroy()
	}
	LevelManagerInsatance.GetInstance().Destroy()
}

type MenuState struct {
	GameState
	renderer   *sdl.Renderer
	staticObjs []Object
}

func MenuStateInit() (*MenuState, error) {
	var m MenuState

	m.renderer = EngineInstance.GetInstance().GetRenderer()
	m.staticObjs = make([]Object, 0)

	playBtn, err := NewButton(&Properties{
		transform: &phy.Transform{X: 10, Y: 20},
		width:     128,
		height:    128,
		flip:      sdl.FLIP_NONE,
	}, []string{"default_btn", "hover_btn", "active_btn"}, m.StartGame)
	if err != nil {
		return nil, err
	}
	m.staticObjs = append(m.staticObjs, playBtn)

	settingsBtn, err := NewButton(&Properties{
		transform: &phy.Transform{X: 10, Y: 20},
		width:     128,
		height:    128,
		flip:      sdl.FLIP_NONE,
	}, []string{"default_btn", "hover_btn", "active_btn"}, m.Settings)
	if err != nil {
		return nil, err
	}
	m.staticObjs = append(m.staticObjs, settingsBtn)

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

func (m MenuState) drawBg(id string, x, y, width, height int, scaleX, scaleY, scrollRatio float64, flip sdl.RendererFlip) {
	dst_ := CameraInstance.GetInstance().SyncObject(&phy.Point{X: float64(x), Y: float64(y)}, int(scrollRatio))
	src := sdl.Rect{X: 0, Y: 0, W: int32(width), H: int32(height)}
	dst := sdl.Rect{
		X: int32(dst_.X),
		Y: int32(dst_.Y),
		W: int32(float64(width) * scaleX),
		H: int32(float64(height) * scaleY),
	}

	textureMap := TextureParserInstance.GetInstance().GetTextureMap()
	err := m.renderer.CopyEx(textureMap[id], &src, &dst, 0, nil, flip)
	if err != nil {
		log.Fatal(err)
		m.renderer.SetDrawColor(0, 155, 0, 250)
	}
}

func (m MenuState) Draw() {
	m.renderer.SetDrawColor(0, 0, 0, 250)
	m.renderer.Clear()
	m.drawBg("menu_bg", 0, 0, WIDTH, HEIGHT, 1, 1, 0.5, sdl.FLIP_NONE)
	for _, mobj := range m.staticObjs {
		mobj.Draw()
	}
	m.renderer.Present()
}

func (m MenuState) Update(dt uint64) {
	m.Events()
	for _, mobj := range m.staticObjs {
		mobj.Update(dt)
	}
}

func (m MenuState) Exit() {}

func (m *MenuState) Quit() {}

type SuccessState struct {
	GameState
	renderer *sdl.Renderer
}

func SuccessStateInit() (*SuccessState, error) {
	s := &SuccessState{}
	s.renderer = EngineInstance.GetInstance().GetRenderer()
	return s, nil
}

func (s SuccessState) Draw() {
	s.renderer.SetDrawColor(0, 255, 0, 255)
	s.renderer.Clear()
	s.renderer.Present()
}

func (s SuccessState) Update(dt uint64) {
	if InputInstance.GetInstance().IsKeyDown(sdl.SCANCODE_M) {
		LevelManagerInsatance.GetInstance().Destroy()
		EngineInstance.GetInstance().SetCurrStateName(MENU)
		LevelManagerInsatance.GetInstance().Init()
	}
}

func (s SuccessState) Exit() {}

func (s SuccessState) Quit() {}

type FailureState struct {
	GameState
	renderer *sdl.Renderer
}

func FailureStateInit() (*FailureState, error) {
	f := &FailureState{}
	f.renderer = EngineInstance.GetInstance().GetRenderer()
	return f, nil
}

func (f FailureState) Draw() {
	f.renderer.SetDrawColor(255, 0, 0, 255)
	f.renderer.Clear()
	f.renderer.Present()
}

func (f FailureState) Update(dt uint64) {
	if InputInstance.GetInstance().IsKeyDown(sdl.SCANCODE_M) {
		LevelManagerInsatance.GetInstance().Destroy()
		EngineInstance.GetInstance().SetCurrStateName(MENU)
		LevelManagerInsatance.GetInstance().Init()
	}
}
