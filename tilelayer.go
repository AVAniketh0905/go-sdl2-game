package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type TileSet struct {
	FirstId int
	LastId  int

	Name string
	Src  string

	RowCount int
	ColCount int

	TileCount int
	TileSize  int
}

type TileSetList []TileSet

type TileSetMap [][]int

type TileLayer struct {
	tileSize int

	rowCount int
	colCount int

	tileMap  TileSetMap
	tileSets TileSetList

	Layer
}

func NewTileLayer(tileSize int, rowCount int, colCount int, tileMap TileSetMap, tileSets TileSetList) *TileLayer {
	for _, ts := range tileSets {
		TextureParserInstance.GetInstance().LoadTexture(ts.Name, "assets/maps/"+ts.Src)
	}

	return &TileLayer{
		tileSize: tileSize,
		rowCount: rowCount,
		colCount: colCount,
		tileMap:  tileMap,
		tileSets: tileSets,
	}
}

func (tl *TileLayer) GetHeight() int {
	return tl.colCount
}

func (tl *TileLayer) GetWidth() int {
	return tl.rowCount
}

func (tl *TileLayer) GetTileMap() TileSetMap {
	return tl.tileMap
}

func (tl TileLayer) Draw() {
	for i := range tl.colCount {
		for j := range tl.rowCount {
			tileId := tl.tileMap[i][j]

			if tileId == 0 {
				continue
			} else {
				var index int
				if len(tl.tileSets) > 1 {
					for k, tSet := range tl.tileSets {
						if tileId > tSet.FirstId && tileId < tSet.LastId {
							tileId = tileId + tSet.TileCount - tSet.LastId
							index = k
							break
						}
					}
				}

				ts := tl.tileSets[index]
				tileRow := tileId / ts.ColCount
				tileCol := tileId - tileRow*ts.ColCount - 1

				if tileId%ts.ColCount == 0 {
					tileRow--
					tileCol = ts.ColCount - 1
				}

				TextureManagerInstance.GetInstance().DrawTile(
					ts.Name,
					ts.TileSize,
					j*ts.TileSize, // cam.x
					i*ts.TileSize, // cam.y
					tileRow,
					tileCol,
					sdl.FLIP_NONE,
				)

			}
		}
	}
}

func (tl TileLayer) Update(dt uint64) {
}

func (tl TileLayer) Destroy() {
}

func (tl TileLayer) String() string {
	size := fmt.Sprintf("Size: {%d, %d}", len(tl.tileMap), len(tl.tileMap[0]))
	return fmt.Sprintf("TileLayer{tileSize: %d, rowCount: %d, colCount: %d, tileMap: %v, tileSets: %v}", tl.tileSize, tl.rowCount, tl.colCount, size, tl.tileSets)
}
