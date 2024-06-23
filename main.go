package main

import "log"

// EngineInstance is a global variable that holds the instance of the Engine (Singlerton)
var EngineInstance *Engine = &Engine{}
var TextureManagerInstance *TextureManager = &TextureManager{}

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
		EngineInstance.Events()
		EngineInstance.Update()
		EngineInstance.Render()
	}

}
