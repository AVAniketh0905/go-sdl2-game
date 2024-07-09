package main

// images
const WIDTH = 800
const HEIGHT = 600

const IMG_SIZE = 64
const TILE_SIZE = 32

// time
const FPS = 60
const DELTA_TIME = 1.5 // target delta time (ms)

// player
const JUMP_TIME = 15
const ATTACK_TIME = 5
const MAX_JUMP_HEIGHT = 10
const RUN_FORCE = 5
const JUMP_FORCE = 15 // 15

// ui
const (
	DEFAULT_BTN = iota
	HOVER_BTN
	ACTIVE_BTN
)

// sound
const FREQ = 44100
const CHUNK_SIZE = 2048

// mimics an comparable interface but only for int/float
type comp interface {
	~int32 | ~int | ~uint | ~float64 | ~float32
}

// type GameObjectTypes interface {
// 	Enemy | Player
// }

func Limit[T comp](nval, ll, ul T) T {
	if nval < ll {
		return ll
	}

	if nval > ul {
		return ul
	}

	return nval
}
