package main

type Layer interface {
	Object
}

type GameMap struct {
	layers []Layer // BUG
}

// func (gm *GameMap) GetLayers() []Layer {
// 	return gm.layers
// }

// func (gm *GameMap) Draw() {
// 	for _, layer := range gm.layers {
// 		layer.Draw()
// 	}
// }

// func (gm *GameMap) Update() {
// 	for _, layer := range gm.layers {
// 		layer.Update()
// 	}
// }