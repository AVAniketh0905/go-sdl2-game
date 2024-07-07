package main

const WIDTH = 800
const HEIGHT = 600

const IMG_SIZE = 64
const TILE_SIZE = 32

const JUMP_TIME = 15
const ATTACK_TIME = 5

const RUN_FORCE = 5
const JUMP_FORCE = 15

const FPS = 60
const DELTA_TIME = 1.5 // target delta time (ms)

const (
	DEFAULT_BTN = iota
	HOVER_BTN
	ACTIVE_BTN
)

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
