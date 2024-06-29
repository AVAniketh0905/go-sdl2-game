package main

import "fmt"

type MapParser struct {
	instance *MapParser
	gameMap  map[string]*GameMap
}

func (mp *MapParser) GetInstance() *MapParser {
	if mp.instance == nil {
		mp.instance = &MapParser{}
	}

	return mp.instance
}

func (mp *MapParser) GetGameMap(id string) *GameMap {
	return mp.gameMap[id]
}

func (mp *MapParser) Load() error {
	if err := mp.parse("level1", "assets/maps/map.tmx"); err != nil {
		return fmt.Errorf("failed to load level1", err)
	}

	return nil
}

func (mp *MapParser) Destroy() {

}

func (mp *MapParser) parse(id string, src string) error {
	return nil
}

func (mp *MapParser) parseTileSet() (tileSet TileSet) {

	return TileSet{}
}

func (mp *MapParser) parseTileLayers() *TileLayer {
	return &TileLayer{}
}
