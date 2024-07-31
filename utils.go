package main

// images
const WIDTH = 800
const HEIGHT = 600

const IMG_SIZE = 64
const TILE_SIZE = 32

// time
const FPS = 60
const TIME_DELAY = 2
const TIME_OFFSET = 1

// player
const MAX_HEALTH = 100
const FIXED_HEALTH_DMG = 1
const ATTACK_TIME = 5
const JUMP_FORCE = 20
const RUN_FORCE = 2

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
