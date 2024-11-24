package pkg

import (
	"fmt"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const HISTORY_QUEUE_LENGTH = 60

type ButtonBuffer struct {
	time uint32
}

func (b *ButtonBuffer) Update(pressing bool) {
	if pressing {
		b.time++
	} else {
		b.time = 0
	}
}

func (b *ButtonBuffer) Pressed() bool {
	return b.time != 0
}

func (b *ButtonBuffer) JustPressed() bool {
	return b.time == 1
}

type MotionDirection uint8

// Numpad notation
const (
	DownLeft MotionDirection = iota + 1
	Down
	DownRight
	Left
	Neutral
	Right
	UpLeft
	Up
	UpRight
)

// TODO: Is this the best way?
func NewMotionDirection(up, down, left, right bool) MotionDirection {
	// We can assume up and down, left and right won't appear together

	if up {
		if left {
			return UpLeft
		} else if right {
			return UpRight
		} else {
			return Up
		}
	} else if down {
		if left {
			return DownLeft
		} else if right {
			return DownRight
		} else {
			return Down
		}
	} else {
		if left {
			return Left
		} else if right {
			return Right
		} else {
			return Neutral
		}
	}
}

type MotionBuffer struct {
	Up    ButtonBuffer
	Down  ButtonBuffer
	Left  ButtonBuffer
	Right ButtonBuffer

	history     [HISTORY_QUEUE_LENGTH]MotionDirection
	historyHead int
}

func (b *MotionBuffer) Update() {
	// NOTE: This is "Neutral" SOCD, should "Last win" be an option?
	up := ebiten.IsKeyPressed(ebiten.KeySpace)
	down := ebiten.IsKeyPressed(ebiten.KeyS)
	left := ebiten.IsKeyPressed(ebiten.KeyA)
	right := ebiten.IsKeyPressed(ebiten.KeyD)

	b.Up.Update(!down && up)
	b.Down.Update(!up && down)
	b.Left.Update(!right && left)
	b.Right.Update(!left && right)

	dir := NewMotionDirection(
		b.Up.Pressed(),
		b.Down.Pressed(),
		b.Left.Pressed(),
		b.Right.Pressed(),
	)

	// Simple queue
	b.historyHead++
	b.historyHead %= HISTORY_QUEUE_LENGTH
	b.history[b.historyHead] = dir

	value := slices.Repeat([]rune{' '}, (b.historyHead*2)+1)
	fmt.Printf("%v\n%v^\n", b.history, string(value))
}

const HELD_LIMIT = 6
const PRE_LIMIT = 5

// TODO: Limit on how long you can hold key
func (b *MotionBuffer) Find(input []MotionDirection) bool {
	history_index := b.historyHead
	input_index := len(input) - 1
	held_time := 0
	pre_time := 0

	for history_index != (b.historyHead+1)%HISTORY_QUEUE_LENGTH {
		dir := b.history[history_index]

		if dir == input[input_index] {
			if input_index == 0 {
				return true
			}

			input_index--

			println("found", dir)
		} else {
			if input_index == len(input)-1 {
				if dir != Neutral {
					return false
				} else {
					if pre_time > PRE_LIMIT {
						return false
					}

					pre_time++
				}
			} else {
				if dir == input[input_index+1] {
					if held_time > HELD_LIMIT {
						return false
					}

					held_time++
				} else {
					return false
				}
			}
		}

		if history_index == 0 {
			history_index = HISTORY_QUEUE_LENGTH
		}
		history_index--
	}

	return false
}

type InputBuffer struct {
	A      ButtonBuffer
	B      ButtonBuffer
	C      ButtonBuffer
	D      ButtonBuffer
	Motion MotionBuffer
}

func (b *InputBuffer) Update() {
	b.A.Update(ebiten.IsKeyPressed(ebiten.KeyJ))
	b.B.Update(ebiten.IsKeyPressed(ebiten.KeyI))
	b.C.Update(ebiten.IsKeyPressed(ebiten.KeyO))
	b.D.Update(ebiten.IsKeyPressed(ebiten.KeyP))

	b.Motion.Update()
}
