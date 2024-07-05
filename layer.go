package main

type Layer interface {
	Draw()
	Update(dt float64)
	Destroy()
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

func (gm *GameMap[l]) Update(dt float64) {
	for _, layer := range gm.layers {
		layer.Update(dt)
	}
}
