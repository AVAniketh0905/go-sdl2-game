package main

import (
	"log"
)

// EngineInstance is a global variable that holds the instance of the Engine (Singlerton)
var InputInstance *Input = &Input{}
var CameraInstance *Camera = &Camera{}
var EngineInstance *Engine = &Engine{}

var TimeInstance *Time = &Time{}
var SoundManagerInstance *SoundManager = &SoundManager{}
var CollisionHandlerInstance *CollisionHandler = &CollisionHandler{}
var TextureManagerInstance *TextureManager = &TextureManager{}

var MapParserInstance *MapParser = &MapParser{}
var ObjectParserInstance *ObjectParser = &ObjectParser{}
var SoundParserInstance *SoundParser = &SoundParser{}

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
}
