package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type CollisionHandler struct {
	instance *CollisionHandler

	collisionTileSetMap TileSetMap
	collisionLayer      TileLayer

	tileSize   int
	tileOffset int
	mapWidth   int
	mapHeight  int
}

func (ch *CollisionHandler) GetInstance() *CollisionHandler {
	if ch.instance == nil {
		ch.instance = &CollisionHandler{
			collisionLayer:      TileLayer{},
			collisionTileSetMap: [][]int{},
		}
	}

	return ch.instance
}

func (ch *CollisionHandler) SetCollisionMap(tileMap TileSetMap, tileSize int) {
	ch.collisionTileSetMap = tileMap
	ch.tileSize = tileSize + ch.tileOffset
	ch.mapHeight = len(tileMap)
	ch.mapWidth = len(tileMap[0])
}

func (ch *CollisionHandler) CheckCollision(a, b *sdl.Rect) bool {
	xOverlap := (a.X < b.X+b.W) && (a.X+a.W > b.X)
	yOverlap := (a.Y < b.Y+b.H) && (a.Y+a.H > b.Y)
	return xOverlap && yOverlap
}

func (ch *CollisionHandler) MapCollision(a *sdl.Rect) bool {
	l, r := a.X/int32(ch.tileSize), (a.X+a.W)/int32(ch.tileSize)
	t, b := a.Y/int32(ch.tileSize), (a.Y+a.H)/int32(ch.tileSize)
	l, r = max(l, 0), min(r, int32(ch.mapWidth))
	t, b = max(t, 0), min(b, int32(ch.mapHeight))

	for i := l; i <= r; i++ {
		for j := t; j <= b; j++ {
			if j < 0 || i < 0 {
				continue
			}
			if j >= int32(ch.mapHeight) || i >= int32(ch.mapWidth) {
				continue
			}

			if ch.collisionTileSetMap[j][i] != 0 {
				return true
			}
		}
	}

	return false
}

type DamageHandler struct {
	instance *DamageHandler

	CollisionHandler
}

func (dh *DamageHandler) GetInstance() *DamageHandler {
	if dh.instance == nil {
		dh.instance = &DamageHandler{}
	}

	dh.tileOffset = 2

	return dh.instance
}
