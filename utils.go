package main

// images
const WIDTH = 800
const HEIGHT = 600

const IMG_SIZE = 64
const TILE_SIZE = 32

// time
const FPS = 30
const TIME_DELAY = 2
const TIME_OFFSET = 1

// player
const ATTACK_TIME = 5
const MAX_JUMP_HEIGHT = 10
const RUN_FORCE = 2.5
const JUMP_FORCE = 12

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

func Integrate(t, dt float64) {

}
