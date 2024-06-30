package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
)

type MapParser struct {
	instance  *MapParser
	gameIdMap map[string]GameMap
}

func (mp *MapParser) GetInstance() *MapParser {
	if mp.instance == nil {
		mp.instance = &MapParser{}
	}

	return mp.instance
}

func (mp *MapParser) GetGameMap(id string) GameMap {
	return mp.gameIdMap[id]
}

func (mp *MapParser) Load() error {
	if err := mp.parse("level1", "assets/maps/map.tmx"); err != nil {
		return fmt.Errorf("failed to load level1, %v", err)
	}

	return nil
}

func (mp *MapParser) Destroy() {
	mp.instance = nil
}

func parseTileSets(xmlMap XMLMap) (tileSets TileSetList) {
	for _, ts := range xmlMap.Tilesets {
		tmpTs := TileSet{}
		tmpTs.FirstId = ts.FirstGId
		tmpTs.LastId = ts.FirstGId + ts.TileCount - 1

		tmpTs.Name = ts.Name
		tmpTs.Src = ts.Image.Source

		tmpTs.ColCount = ts.NumCols
		tmpTs.RowCount = ts.TileCount / tmpTs.ColCount

		tmpTs.TileCount = ts.TileCount
		tmpTs.TileSize = ts.TileWidth

		tileSets = append(tileSets, tmpTs)
	}

	return tileSets
}

func getData(stream []byte, width int) (TileSetMap, error) {
	data := TileSetMap{}
	j, tmpJ := 0, []int{}
	for _, b := range stream {
		i, err := strconv.Atoi(string(b))
		if err != nil {
			return nil, fmt.Errorf("failed to convert to string, %b, %v", b, err)
		}
		if j == width {
			data = append(data, tmpJ)
			j, tmpJ = 0, []int{}
		} else {
			tmpJ = append(tmpJ, i)
		}
		j++
	}

	return data, nil
}

func parseTileLayers(xmlMap XMLMap, tileSets TileSetList, tileSize, RowCount, ColCount int) (layers []Layer, err error) {
	for _, l := range xmlMap.Layers {
		stream := l.Data.Content

		data, err := getData(stream, l.Width)
		if err != nil {
			return nil, fmt.Errorf("failed to load data into TileSetMap, %v", err)
		}

		newL := NewTileLayer(
			tileSize,
			RowCount,
			ColCount,
			data,
			tileSets,
		)
		layers = append(layers, *newL)
	}
	return layers, nil
}

func (mp *MapParser) parse(id string, src string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("failed to read from source %v %v", src, err)
	}

	var xmlMap XMLMap

	err = xml.Unmarshal(data, &xmlMap)
	if err != nil {
		return fmt.Errorf("failed to unmarshal xml %v", err)
	}

	tileSets := parseTileSets(xmlMap)
	tileLayers, err := parseTileLayers(xmlMap, tileSets, xmlMap.TileWidth, xmlMap.Width, xmlMap.Height)
	if err != nil {
		return fmt.Errorf("failed to parseTileLayers, %v", err)
	}

	var gameMap GameMap
	gameMap.layers = append(gameMap.layers, tileLayers...)

	mp.gameIdMap[id] = gameMap
	return nil
}

// XML
type XMLMap struct {
	XMLName   xml.Name     `xml:"map"`
	Width     int          `xml:"width,attr"`
	Height    int          `xml:"height,attr"`
	TileWidth int          `xml:"tilewidth,attr"`
	Tilesets  []XMLTileset `xml:"tileset"`
	Layers    []XMLLayer   `xml:"layer"`
}

type XMLTileset struct {
	FirstGId   int      `xml:"firstgid,attr"`
	TileWidth  int      `xml:"tilewidth,attr"`
	TileHeight int      `xml:"tileheight,attr"`
	TileCount  int      `xml:"tilecount,attr"`
	NumCols    int      `xml:"columns,attr"`
	Name       string   `xml:"name,attr"`
	Image      XMLImage `xml:"tileset>image"`
}

type XMLImage struct {
	Source string `xml:"source,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
}

type XMLLayer struct {
	ID     int     `xml:"id,attr"`
	Name   string  `xml:"name,attr"`
	Width  int     `xml:"width,attr"`
	Height int     `xml:"height,attr"`
	Data   XMLData `xml:"data"`
}

type XMLData struct {
	Content []byte `xml:",chardata"`
}
