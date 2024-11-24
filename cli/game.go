package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hen6003/go-fighting-game/v2/pkg"
)

func main() {
	game := pkg.NewGame()

	// Specify the window size as you like. Here, a doubled size is specified.
	//ebiten.SetWindowSize(640, 480)
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Game")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
