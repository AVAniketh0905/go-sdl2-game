package main

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

// EngineInstance is a global variable that holds the instance of the Engine (Singlerton)
var InputInstance *Input = &Input{}
var CameraInstance *Camera = &Camera{}
var EngineInstance *Engine = &Engine{}

var TimeInstance *Time = &Time{}
var SoundManagerInstance *SoundManager = &SoundManager{}
var CollisionHandlerInstance *CollisionHandler = &CollisionHandler{}
var TextureManagerInstance *TextureManager = &TextureManager{}
var LevelManagerInsatance *LevelManager = &LevelManager{}

var TextureParserInstance *TextureParser = &TextureParser{}
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
	defer TextureParserInstance.GetInstance().Destroy()
	defer MapParserInstance.GetInstance().Destroy()
	defer ObjectParserInstance.GetInstance().Destroy()
	defer SoundParserInstance.GetInstance().Destroy()

	var lastTime uint64
	for EngineInstance.GetInstance().IsRunning {
		startTime := sdl.GetTicks64()
		dt := startTime - lastTime

		if dt > uint64(1000/FPS) {
			fmt.Println("FPS: ", 1000.0/float64(dt))
			lastTime = startTime

			EngineInstance.GetInstance().Events()
			EngineInstance.GetInstance().Update(dt / 10)
			EngineInstance.GetInstance().Draw()
		}
		TimeInstance.GetInstance().Tick()
	}
}

func main() {
	Core()
}
