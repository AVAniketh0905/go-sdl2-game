package main

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

func (tl *TileLayer) Draw() {

}

func (tl *TileLayer) Update() {

}

func (tl *TileLayer) Destroy() {
}
