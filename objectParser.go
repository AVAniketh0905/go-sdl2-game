package main

import (
	"encoding/xml"
	"fmt"
	"go-game/phy"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type ObjectParser struct {
	instance *ObjectParser
}

func (op *ObjectParser) GetInstance() *ObjectParser {
	if op.instance == nil {
		op.instance = &ObjectParser{}
	}

	return op.instance
}

func (op *ObjectParser) Load(src string) ([]Object, error) {
	objects, err := op.parse(src)
	if err != nil {
		return nil, fmt.Errorf("failed to objects from xml file, %v", err)
	}

	return objects, nil
}

func (op *ObjectParser) convertStrToFlip(xmlFlip string) sdl.RendererFlip {
	switch xmlFlip {
	case "None":
		return sdl.FLIP_NONE
	case "Horizontal":
		return sdl.FLIP_HORIZONTAL
	case "Vertical":
		return sdl.FLIP_VERTICAL
	}

	return sdl.FLIP_NONE
}

func (op *ObjectParser) convertToObject(xmlObj XMLObject) (Object, error) {
	flip := op.convertStrToFlip(xmlObj.Flip)

	props := Properties{
		transform: &phy.Transform{X: xmlObj.X, Y: xmlObj.Y},
		width:     xmlObj.Width,
		height:    xmlObj.Height,
		texId:     xmlObj.TexId,
		flip:      flip,
	}

	obj, err := CreateObjectFactory(xmlObj.Type, &props)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to object, %v", err)
	}

	return obj, nil
}

func (op *ObjectParser) parse(src string) ([]Object, error) {
	data, err := os.ReadFile(src)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from file, %v, %v", src, err)
	}

	var xmlObjects XMLObjects

	err = xml.Unmarshal(data, &xmlObjects)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal xml, %v", err)
	}

	var objects []Object
	for _, xmlObj := range xmlObjects.Objects {
		obj, err := op.convertToObject(xmlObj)
		if err != nil {
			return nil, fmt.Errorf("failed to convert the following xmlObj to object, %v, %v", xmlObj, err)
		}
		objects = append(objects, obj)
	}

	return objects, nil
}

func (op *ObjectParser) Destroy() {
	op.instance = nil
}

type XMLObjects struct {
	Objects []XMLObject `xml:"object,"`
}

type XMLObject struct {
	Type   string  `xml:"type,attr"`
	X      float64 `xml:"x,attr"`
	Y      float64 `xml:"y,attr"`
	Width  int     `xml:"width,attr"`
	Height int     `xml:"height,attr"`
	TexId  string  `xml:"texId,attr"`
	Flip   string  `xml:"flip,attr"`
}
