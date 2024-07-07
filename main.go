package main

import (
	"log"
)

// EngineInstance is a global variable that holds the instance of the Engine (Singlerton)
var EngineInstance *Engine = &Engine{}
var CameraInstance *Camera = &Camera{}
var InputInstance *Input = &Input{}

var TextureManagerInstance *TextureManager = &TextureManager{}
var CollisionHandlerInstance *CollisionHandler = &CollisionHandler{}
var TimeInstance *Time = &Time{}

var MapParserInstance *MapParser = &MapParser{}
var ObjectParserInstance *ObjectParser = &ObjectParser{}

func Core() {
	err := EngineInstance.GetInstance().Load()
	if err != nil {
		log.Fatalf("Failed to load EngineInstance: %v", err)
	}
	defer EngineInstance.GetInstance().Destroy()
	defer TextureManagerInstance.GetInstance().Destroy()
	defer MapParserInstance.GetInstance().Destroy()

	for EngineInstance.GetInstance().IsRunning {
		EngineInstance.GetInstance().Events()
		EngineInstance.GetInstance().Update()
		EngineInstance.GetInstance().Draw()
		TimeInstance.GetInstance().Tick()
	}
}

func main() {
	Core()
	// Build()
}
