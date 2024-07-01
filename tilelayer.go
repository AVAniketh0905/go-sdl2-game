package main

import (
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
		TextureManagerInstance.LoadTexture(ts.Name, "assets/maps/"+ts.Src)
	}

	return &TileLayer{
		tileSize: tileSize,
		rowCount: rowCount,
		colCount: colCount,
		tileMap:  tileMap,
		tileSets: tileSets,
	}
}

func (tl *TileLayer) GetTileMap() TileSetMap {
	return tl.tileMap
}

func (tl TileLayer) Draw() {
	for i := range tl.rowCount {
		for j := range tl.colCount {
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

				//fmt.Println(ts)

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

func (tl TileLayer) Update() {

}

func (tl TileLayer) Destroy() {
}
