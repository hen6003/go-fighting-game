package pkg

import (
	"math"

	"github.com/joonazan/vec2"
)

const PLAYER_WIDTH = 32
const PLAYER_HEIGHT = 48

const DISTANCE_FROM_WALLS = 50

type Player struct {
	pos vec2.Vector
	vel vec2.Vector
}

func NewPlayer(id int) Player {
	var pos float64

	if id == 0 {
		pos = DISTANCE_FROM_WALLS
	} else {
		pos = SCREEN_WIDTH - PLAYER_WIDTH - DISTANCE_FROM_WALLS
	}

	return Player{pos: vec2.Vector{X: pos, Y: PLAYER_HEIGHT}, vel: vec2.Vector{}}
}

func (p *Player) LocalUpdate() {
	on_ground := p.pos.Y == PLAYER_HEIGHT

	if p.pos.Y > PLAYER_HEIGHT {
		p.vel.Y -= .98
	}

	if p.vel.X != 0 {
		var decel float64

		if on_ground {
			decel = math.Copysign(0.1, p.vel.X)
		} else {
			decel = math.Copysign(0.05, p.vel.X)
		}

		if math.Abs(decel) > math.Abs(p.vel.X) {
			p.vel.X = 0
		} else {
			p.vel.X -= decel
		}
	}

	p.pos.Add(p.vel)

	if p.pos.Y < PLAYER_HEIGHT {
		p.pos.Y = PLAYER_HEIGHT
		p.vel.Y = 0
	}
}

func (p *Player) Input(input_buffer InputBuffer) {
	on_ground := p.pos.Y == PLAYER_HEIGHT

	if input_buffer.Motion.Up.JustPressed() && on_ground {
		p.vel.Y += 13
	}

	if input_buffer.Motion.Down.Pressed() {
		//p.pos.Y += 1
	}

	if input_buffer.Motion.Left.Pressed() && p.vel.X > -2 {
		if on_ground {
			p.vel.X -= 0.3
		} else {
			p.vel.X -= 0.1
		}
	}

	if input_buffer.Motion.Right.Pressed() && p.vel.X < 2 {
		if on_ground {
			p.vel.X += 0.3
		} else {
			p.vel.X += 0.1
		}
	}
}
