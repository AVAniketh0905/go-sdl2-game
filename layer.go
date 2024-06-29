package main

type Layer struct {
	Object
}

type GameMap struct {
	layers []Layer
}

func (gm *GameMap) GetLayers() []Layer {
	return gm.layers
}

func (gm *GameMap) Draw() {
	for _, layer := range gm.layers {
		layer.Draw()
	}
}

func (gm *GameMap) Update() {
	for _, layer := range gm.layers {
		layer.Update()
	}
}
