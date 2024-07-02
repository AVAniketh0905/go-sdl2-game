package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type CollisionHandler struct {
	instance *CollisionHandler

	collisionTileSetMap TileSetMap
	collisionLayer      TileLayer
}

func (ch *CollisionHandler) GetInstance() *CollisionHandler {
	lvlLayers := EngineInstance.GetInstance().GetLevelMap().GetLayers()
	// TODO: TileLayer manually imputed as an genric
	layer := lvlLayers[0]

	if ch.instance == nil {
		ch.instance = &CollisionHandler{
			collisionLayer:      layer,
			collisionTileSetMap: layer.GetTileMap(),
		}
	}

	return ch.instance
}

func (ch *CollisionHandler) CheckCollision(a, b *sdl.Rect) bool {
	xOverlap := (a.X < b.X+b.W) && (a.X+a.W > b.X)
	yOverlap := (a.Y < b.Y+b.H) && (a.Y+a.H > b.Y)
	return xOverlap && yOverlap
}

func (ch *CollisionHandler) MapCollision(a *sdl.Rect) bool {
	// fixed numbers based on the map size
	tileSize, rowCount, colCount := TILE_SIZE, 30, 30
	l, r := a.X/int32(tileSize), (a.X+a.W)/int32(tileSize)
	t, b := a.Y/int32(tileSize), (a.Y+a.H)/int32(tileSize)
	l, r = max(l, 0), min(r, int32(colCount))
	t, b = max(t, 0), min(b, int32(rowCount))

	for i := l; i < r; i++ {
		for j := t; j < b; j++ {
			if ch.collisionTileSetMap[j][i] != 0 {
				return true
			}
		}
	}

	return false
}
