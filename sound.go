package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/mix"
)

var (
	EffectMap = make(map[string]*mix.Chunk)
	MusicMap  = make(map[string]*mix.Music)
)

type SoundManager struct {
	instance *SoundManager

	musicMap  map[string]*mix.Music
	effectMap map[string]*mix.Chunk
}

func (sm *SoundManager) GetInstance() (*SoundManager, error) {
	if err := mix.OpenAudio(FREQ, uint16(mix.DEFAULT_FORMAT), 2, CHUNK_SIZE); err != nil {
		return nil, fmt.Errorf("failed to run open audio, %v", err)
	}

	if sm.instance == nil {
		sm.instance = &SoundManager{
			musicMap:  MusicMap,
			effectMap: EffectMap,
		}
	}

	return sm.instance, nil
}

func (sm *SoundManager) LoadMusic(id, src string) error {
	music, err := mix.LoadMUS(src)
	if err != nil {
		return fmt.Errorf("failed to load music from, %v, %v", src, err)
	}

	sm.musicMap[id] = music
	return nil
}

func (sm *SoundManager) LoadEffect(id, src string) error {
	effect, err := mix.LoadWAV(src)
	if err != nil {
		return fmt.Errorf("failed to load effect from, %v, %v", src, err)
	}

	sm.effectMap[id] = effect
	return nil
}

func (sm *SoundManager) PlayMusic(id string) error {
	music := sm.musicMap[id]

	if err := music.Play(-1); err != nil {
		return fmt.Errorf("failed to play music, %v", err)
	}

	return nil
}

func (sm *SoundManager) PlayEffect(id string) error {
	effect := sm.effectMap[id]

	if _, err := effect.Play(-1, 0); err != nil {
		return fmt.Errorf("failed to play music, %v", err)
	}

	return nil
}

func (sm *SoundManager) Destroy() {
	sm.musicMap = make(map[string]*mix.Music)
	sm.effectMap = make(map[string]*mix.Chunk)
}
