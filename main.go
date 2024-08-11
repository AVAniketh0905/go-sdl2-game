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
var TextureManagerInstance *TextureManager = &TextureManager{}
var LevelManagerInsatance *LevelManager = &LevelManager{}

var CollisionHandlerInstance *CollisionHandler = &CollisionHandler{}
var DamageHandlerInstance *DamageHandler = &DamageHandler{}
var CoinHandlerInstance *CoinHandler = &CoinHandler{}

var TextureParserInstance *TextureParser = &TextureParser{}
var MapParserInstance *MapParser = &MapParser{}
var ObjectParserInstance *ObjectParser = &ObjectParser{}
var SoundParserInstance *SoundParser = &SoundParser{}

func Core() {
	err := EngineInstance.GetInstance().Load()
	if err != nil {
		log.Fatalf("Failed to load EngineInstance: %v", err)
	}
	defer Destroy()

	for EngineInstance.GetInstance().IsRunning {
		TimeInstance.GetInstance().Start()
		dt := TimeInstance.GetInstance().GetDeltaTime()

		EngineInstance.GetInstance().Events()
		EngineInstance.GetInstance().Update(dt)
		EngineInstance.GetInstance().Draw()

		TimeInstance.GetInstance().Tick()
	}
}

func Destroy() {
	EngineInstance.GetInstance().Destroy()
	TextureManagerInstance.GetInstance().Destroy()
	TextureParserInstance.GetInstance().Destroy()
	MapParserInstance.GetInstance().Destroy()
	ObjectParserInstance.GetInstance().Destroy()
	SoundParserInstance.GetInstance().Destroy()
	SoundManagerInstance.GetInstance().Destroy()
	LevelManagerInsatance.GetInstance().Destroy()
}

func main() {
	Core()
}
