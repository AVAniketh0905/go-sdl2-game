package main

type Layer interface {
	Object
}

type GameMap[l Layer] struct {
	layers []l
}

func (gm *GameMap[l]) GetLayers() []l {
	return gm.layers
}

func (gm *GameMap[l]) Draw() {
	for _, layer := range gm.layers {
		layer.Draw()
	}
}

func (gm *GameMap[l]) Update() {
	for _, layer := range gm.layers {
		layer.Update()
	}
}
