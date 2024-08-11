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
	bg       *sdl.Color
	menuObjs []Object
}

func PlayStateInit() (*PlayState, error) {
	p := &PlayState{}
	err := LevelManagerInsatance.GetInstance().Init()
	if err != nil {
		return nil, fmt.Errorf("failed to load lvls, %v", err)
	}
	p.bg = LevelManagerInsatance.GetInstance().GetBgColor()

	props := Properties{
		transform: &phy.Transform{X: 70, Y: 82},
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
	switch LevelManagerInsatance.GetInstance().GetState() {
	case SUCCESS:
		EngineInstance.GetInstance().SetCurrStateName(SUCCESS)
	case FAIL:
		EngineInstance.GetInstance().SetCurrStateName(FAIL)
	}
}

func (p PlayState) Draw() {
	EngineInstance.GetInstance().GetRenderer().SetDrawColor(p.bg.R, p.bg.G, p.bg.B, p.bg.A)
	EngineInstance.GetInstance().GetRenderer().Clear()
	LevelManagerInsatance.GetInstance().Draw()
	for _, mobj := range p.menuObjs {
		mobj.Draw()
	}
	EngineInstance.GetInstance().GetRenderer().Present()
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
	staticObjs []Object
	texts      []*Text
}

func MenuStateInit() (*MenuState, error) {
	var m MenuState

	m.staticObjs = make([]Object, 0)

	playBtn, err := NewButton(&Properties{
		transform: &phy.Transform{X: WIDTH/2 + 50, Y: HEIGHT/2 + 100},
		width:     256,
		height:    256,
		flip:      sdl.FLIP_NONE,
	}, []string{"play_def_btn", "play_hov_btn", "play_act_btn"}, m.StartGame)
	if err != nil {
		return nil, err
	}

	welcomeText, err := NewText(32, "assets/fonts/Array-Regular.ttf", sdl.Color{R: 255, G: 255, B: 255, A: 255}, "Welcome to the game!", WIDTH/2-150, HEIGHT/2-150, 600, 200)
	if err != nil {
		return nil, err
	}

	m.staticObjs = append(m.staticObjs, playBtn)
	m.texts = append(m.texts, welcomeText)

	// settingsBtn, err := NewButton(&Properties{
	// 	transform: &phy.Transform{X: 10, Y: 20},
	// 	width:     128,
	// 	height:    128,
	// 	flip:      sdl.FLIP_NONE,
	// }, []string{"default_btn", "hover_btn", "active_btn"}, m.Settings)
	// if err != nil {
	// 	return nil, err
	// }
	// m.staticObjs = append(m.staticObjs, settingsBtn)

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
	EngineInstance.GetInstance().GetRenderer().SetDrawColor(0, 0, 0, 255)
	EngineInstance.GetInstance().GetRenderer().Clear()
	TextureManagerInstance.GetInstance().Draw("menu_bg", 0, 0, WIDTH, HEIGHT, 1, 1, 0.5, sdl.FLIP_NONE)
	for _, text := range m.texts {
		text.Draw()
	}
	for _, mobj := range m.staticObjs {
		mobj.Draw()
	}
	EngineInstance.GetInstance().GetRenderer().Present()
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
	MenuState
	objs  []Object
	texts []*Text
}

func SuccessStateInit() (*SuccessState, error) {
	s := &SuccessState{}

	homeBtn, err := NewButton(&Properties{
		transform: &phy.Transform{X: WIDTH/2 + 70, Y: HEIGHT/2 + 70},
		width:     256,
		height:    256,
		flip:      sdl.FLIP_NONE,
	}, []string{"default_btn", "hover_btn", "active_btn"}, s.GoHome)
	if err != nil {
		return nil, err
	}

	completedText, err := NewText(32, "assets/fonts/Array-Regular.ttf", sdl.Color{R: 0, B: 0, G: 0, A: 255}, "Level Completed!", WIDTH/2-70, HEIGHT/2-100, 400, 100)
	if err != nil {
		return nil, err
	}

	s.objs = append(s.objs, homeBtn)
	s.texts = append(s.texts, completedText)
	return s, nil
}

func (s *SuccessState) GoHome() {
	LevelManagerInsatance.GetInstance().Destroy()
	EngineInstance.GetInstance().SetCurrStateName(MENU)
	LevelManagerInsatance.GetInstance().Init()
}

func (s SuccessState) Draw() {
	EngineInstance.GetInstance().GetRenderer().SetDrawColor(0, 255, 0, 255)
	EngineInstance.GetInstance().GetRenderer().Clear()
	for _, text := range s.texts {
		text.Draw()
	}
	for _, obj := range s.objs {
		obj.Draw()
	}
	EngineInstance.GetInstance().GetRenderer().Present()

}

func (s SuccessState) Update(dt uint64) {
	if InputInstance.GetInstance().IsKeyDown(sdl.SCANCODE_M) {
		LevelManagerInsatance.GetInstance().Destroy()
		EngineInstance.GetInstance().SetCurrStateName(MENU)
		LevelManagerInsatance.GetInstance().Init()
	}

	for _, obj := range s.objs {
		obj.Update(dt)
	}
}

func (s SuccessState) Exit() {}

func (s SuccessState) Quit() {}

type FailureState struct {
	MenuState

	objs  []Object
	texts []*Text
}

func FailureStateInit() (*FailureState, error) {
	f := &FailureState{}

	homeBtn, err := NewButton(&Properties{
		transform: &phy.Transform{X: WIDTH/2 + 50, Y: HEIGHT/2 + 50},
		width:     200,
		height:    200,
		flip:      sdl.FLIP_NONE,
	}, []string{"default_btn", "hover_btn", "active_btn"}, f.GoHome)
	if err != nil {
		return nil, err
	}

	failureText, err := NewText(32, "assets/fonts/Array-Regular.ttf", sdl.Color{R: 255, G: 255, B: 255, A: 255}, "Level Failed!", WIDTH/2-100, HEIGHT/2-100, 400, 100)
	if err != nil {
		return nil, err
	}

	f.objs = append(f.objs, homeBtn)
	f.texts = append(f.texts, failureText)

	return f, nil
}

func (f *FailureState) GoHome() {
	LevelManagerInsatance.GetInstance().Destroy()
	EngineInstance.GetInstance().SetCurrStateName(MENU)
	LevelManagerInsatance.GetInstance().Init()
}

func (f FailureState) Draw() {
	EngineInstance.GetInstance().GetRenderer().SetDrawColor(255, 0, 0, 255)
	EngineInstance.GetInstance().GetRenderer().Clear()
	for _, text := range f.texts {
		text.Draw()
	}
	for _, obj := range f.objs {
		obj.Draw()
	}
	EngineInstance.GetInstance().GetRenderer().Present()
}

func (f FailureState) Update(dt uint64) {
	if InputInstance.GetInstance().IsKeyDown(sdl.SCANCODE_M) {
		LevelManagerInsatance.GetInstance().Destroy()
		EngineInstance.GetInstance().SetCurrStateName(MENU)
		LevelManagerInsatance.GetInstance().Init()
	}

	for _, obj := range f.objs {
		obj.Update(dt)
	}
}
