package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

// EngineInstance is a global variable that holds the instance of the Engine (Singlerton)
var EngineInstance *Engine = &Engine{}
var TextureManagerInstance *TextureManager = &TextureManager{}
var InputInstance *Input = &Input{}

func main() {
	EngineInstance.GetInstance()
	err := EngineInstance.Init()
	if err != nil {
		log.Fatalf("Failed to initialize EngineInstance: %v", err)
	}
	err = EngineInstance.Load()
	if err != nil {
		log.Fatalf("Failed to load EngineInstance: %v", err)
	}
	defer EngineInstance.Destroy()
	defer TextureManagerInstance.Destroy()

	for EngineInstance.IsRunning {
		// TODO: handle FPS
		EngineInstance.Events()
		EngineInstance.Update()
		EngineInstance.Render()

		// Temporary fix for high CPU usage
		sdl.Delay(10000 / 60)
	}

}
