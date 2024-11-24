package pkg

import (
	//"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const SCREEN_WIDTH = 320
const SCREEN_HEIGHT = 240

var (
	whiteImage = ebiten.NewImage(1, 1)
)

func init() {
	whiteImage.Fill(color.White)
}

// Game implements ebiten.Game interface.
type Game struct {
	input_buffer InputBuffer
	players      [2]Player
}

func NewGame() *Game {
	return &Game{
		players: [2]Player{NewPlayer(0), NewPlayer(1)},
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	g.input_buffer.Update()
	g.players[0].Input(g.input_buffer)
	g.players[0].LocalUpdate()
	g.players[1].LocalUpdate()

	if g.input_buffer.A.JustPressed() {
		if g.input_buffer.Motion.Find([]MotionDirection{2, 3, 6}) {
			println("FIREBALL")
		}
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	for i, p := range g.players {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Scale(PLAYER_WIDTH, PLAYER_HEIGHT)
		op.GeoM.Translate(p.pos.X, float64(screen.Bounds().Dy())-p.pos.Y)
		op.ColorScale.ScaleWithColor(color.RGBA{100 * uint8(1-i), 100 * uint8(i), 100, 255})

		screen.DrawImage(whiteImage, &op)
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
